package web

import (
	"errors"
	"os"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Webserver struct {
	fiber.Config
}

func NewWebserver() *fiber.App {
	return fiber.New(fiber.Config{
		AppName: os.Getenv("PROJECT_NAME"),
		Prefork: false,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		EnableTrustedProxyCheck: true,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			var code int = fiber.StatusNotFound
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
			return c.Status(code).JSON(fiber.Map{
				`errors`: e.Error(),
			})
		},
		Immutable: true,
		BodyLimit: 40 * 1024 * 1024,
	})
}