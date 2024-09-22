package logger

import (
	"io"
	"os"
	"strconv"

	"myapi/configs"

	"github.com/rs/zerolog"
	rotateLogs "github.com/sutantodadang/go-rotate-logs"
)

var Log zerolog.Logger

func InitLogger() {
	logLevel, err := strconv.Atoi(os.Getenv(`LOG_LEVEL`))
	if err != nil {
		logLevel = int(zerolog.InfoLevel)
	}

	var zlog zerolog.Logger

	if os.Getenv(`PROJECT_ENV`) == `prod` {
		logFile := &rotateLogs.RotateLogsWriter{
			Config: rotateLogs.Config{
				Directory:     configs.OS_PATH_APPLICATION_LOG,
				Filename:      "app.log",
				MaxSize:       10,
				UsingTime:     true,
				FormatTime:    "02-01-2006",
				CleanOldFiles: true,
				MaxAge:        60,
			},
		}
		zlog = zerolog.New(logFile).
			Level(zerolog.Level(logLevel)).
			With().Timestamp().Caller().Logger()
	} else {
		var output io.Writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: `2006/01/02 03:04 PM`,
			PartsOrder: []string{
				zerolog.TimestampFieldName,
				zerolog.LevelFieldName,
				zerolog.CallerFieldName,
				zerolog.MessageFieldName,
			},
		}
		zlog = zerolog.New(output).
			Level(zerolog.Level(logLevel)).
			With().Timestamp().Caller().Logger()
	}

	Log = zlog
}
