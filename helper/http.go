package helper

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

// Parse body request to a valid type
func ReadBody[T any | struct{}](c *fiber.Ctx) (out T, err error) {
	err = c.BodyParser(&out)
	if err != nil {
		err = errors.New(`invalid payload, please check your request and try again`)
		return
	}

	err = ValidateStruct(out)
	if err != nil {
		return
	}

	return
}

// Parse query parameter(s) to a valid type
func ReadQuery[T any | struct{}](c *fiber.Ctx) (out T, err error) {
	err = c.QueryParser(&out)
	if err != nil {
		err = errors.New(`invalid query, please check your query parameter and try again`)
		return
	}

	err = ValidateStruct(out)
	if err != nil {
		return
	}
	
	return
}