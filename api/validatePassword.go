package api

import (
	"strings"

	"github.com/Thund3rD3v/SuperGuardian/structs"
	"github.com/gofiber/fiber/v2"
)

func ValidatePasswordMiddleware(password string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		headers := ctx.GetReqHeaders()

		if strings.Join(headers["Authorization"], "") != password {
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Invalid Password",
			})
		}

		return ctx.Next()
	}
}
