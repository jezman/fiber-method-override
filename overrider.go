package overrider

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// processing only post requests
		if c.Method() != fiber.MethodPost {
			return c.Next()
		}

		// checking the hidden field "_method"
		_method := strings.ToUpper(c.FormValue("_method"))

		if _method == "" {
			return c.Next()
		}

		// override method
		switch _method {
		case fiber.MethodPut:
			c.Method(fiber.MethodPut)
		case fiber.MethodPatch:
			c.Method(fiber.MethodPatch)
		case fiber.MethodDelete:
			c.Method(fiber.MethodDelete)
		default:
			return c.Next()
		}

		return c.Next()
	}
}
