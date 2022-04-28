package middleware

import (
	"github.com/gofiber/fiber/v2"
)

type Error struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	errResult := new(Error)
	errResult.Message = err.Error()
	errResult.Error = err

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		if code == fiber.ErrNotFound.Code {
			return NotFound(c)
		}
	}

	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	c.Set("Access-Control-Allow-Origin", "*")

	c.Status(code)

	return c.JSON(errResult)
}
