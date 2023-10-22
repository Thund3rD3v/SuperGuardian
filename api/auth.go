package api

import (
	"encoding/json"

	"github.com/Thund3rD3v/SuperGuardian/structs"
	"github.com/gofiber/fiber/v2"
)

type LoginBody struct {
	Password string `json:"password"`
}

func LoginRoute(password string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var body LoginBody
		raw := ctx.Request().Body()

		err := json.Unmarshal(raw, &body)
		if err != nil {
			ctx.JSON(structs.Response{
				Success: false,
				Message: "Invalid JSON Body",
			})
		}

		if body.Password == password {
			return ctx.JSON(structs.Response{
				Success: true,
				Message: "Logged In",
			})
		}

		return ctx.JSON(structs.Response{
			Success: false,
			Message: "Invalid Password",
		})
	}
}
