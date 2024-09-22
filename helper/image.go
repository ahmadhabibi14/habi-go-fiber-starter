package helper

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"myapi/internal/bootstrap/logger"

	"github.com/google/uuid"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"github.com/nfnt/resize"
)

func IMG_ConvertToWEBP(writer io.Writer, reader io.Reader, contentType string, width uint, height uint) error {
	var img			image.Image
	var err 		error
	var options *encoder.Options
	
	switch contentType {
	case `image/png`:
		img, err = png.Decode(reader)
	case `image/jpeg`:
		img, err = jpeg.Decode(reader)
	default:
		img, _, err = image.Decode(reader)
	}

	if err != nil {
		errMsg := `invalid file, must be image`
		logger.Log.Err(err).Msg(errMsg)
		return errors.New(errMsg)
	}

	options, err = encoder.NewLossyEncoderOptions(encoder.PresetDefault, 80)
	if err != nil {
		return err
	}

	img = resize.Resize(width, height, img, resize.Lanczos3)

	err = webp.Encode(writer, img, options)
	if err != nil {
		errMsg := `failed to encode image to webp`
		logger.Log.Err(err).Msg(errMsg)
		return errors.New(errMsg)
	}

	return nil
}

func ResizeImageWEBP(fileHeader multipart.FileHeader, minWidth, minHeight, stdWidth uint) (out []byte, fileName string, err error) {
	buff := make([]byte, 512)
	reader, errOpen := fileHeader.Open()
	if errOpen != nil {
		err = errors.New(`failed to open file`)
		logger.Log.Err(errOpen).Msg(err.Error())

		return
	}

	_, errRead := reader.Read(buff)
	if errRead != nil {
		reader.Close()
		err = errors.New(`failed to read file`)

		logger.Log.Err(errRead).Msg(err.Error())

		return
	}

	_, errSeek := reader.Seek(0, io.SeekStart)
	if errSeek != nil {
		err = errors.New(`error while processing file`)
		logger.Log.Err(errSeek).Msg(err.Error())

		return
	}

	fileName = strings.Replace(uuid.New().String(), "-", "", -1) + ".webp"

	var img			image.Image
	var options *encoder.Options

	contentType := http.DetectContentType(buff)
	
	switch contentType {
	case `image/png`:
		img, err = png.Decode(reader)
	case `image/jpeg`:
		img, err = jpeg.Decode(reader)
	default:
		img, _, err = image.Decode(reader)
	}

	if err != nil {
		errMsg := `invalid file, must be image`
		logger.Log.Err(err).Msg(errMsg)
		err = errors.New(errMsg)

		return
	}

	options, err = encoder.NewLossyEncoderOptions(encoder.PresetDefault, 80)
	if err != nil {
		return
	}

	rect := img.Bounds()
	point := rect.Size()

	width, height := uint(point.X), uint(point.Y)

	if width < minWidth {
		err = errors.New(`image width must be at least ` + strconv.Itoa(int(minWidth)) + `px`)
		return
	} else if width > stdWidth {
		width = uint(stdWidth)
	}

	if height < minHeight {
		err = errors.New(`image height must be at least ` + strconv.Itoa(int(minHeight)) + `px`)
		return
	}

	img = resize.Resize(width, 0, img, resize.Lanczos3)

	output := new(bytes.Buffer)
	err = webp.Encode(output, img, options)
	if err != nil {
		errMsg := `failed to encode image to webp`
		logger.Log.Err(err).Msg(errMsg)
		err = errors.New(errMsg)

		return
	}

	out = output.Bytes()

	return
}