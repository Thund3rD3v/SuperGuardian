package api

import (
	"encoding/json"

	"github.com/Thund3rD3v/SuperGuardian/structs"
	"github.com/Thund3rD3v/SuperGuardian/utils"
	"github.com/gofiber/fiber/v2"
)

type EditJoinRolesBody struct {
	Enabled     bool     `json:"enabled"`
	IncludeBots bool     `json:"includeBots"`
	Roles       []string `json:"roles"`
}

func JoinRolesRoute() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		config := utils.GetConfig()

		return ctx.JSON(structs.Response{
			Success: true,
			Data: structs.AnyData{
				"enabled":     config.JoinRoles.Enabled,
				"includeBots": config.JoinRoles.IncludeBots,
				"roles":       config.JoinRoles.Roles,
			},
		})
	}
}

func EditJoinRolesRoute() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Parse Json Body
		var body EditJoinRolesBody
		raw := ctx.Request().Body()
		err := json.Unmarshal(raw, &body)
		if err != nil {
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Invalid JSON Body",
			})
		}

		config := utils.GetConfig()

		config.JoinRoles.Enabled = body.Enabled
		config.JoinRoles.IncludeBots = body.IncludeBots
		config.JoinRoles.Roles = body.Roles

		err = utils.WriteConfig(config)
		if err != nil {
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Could Not Save Join Roles",
			})
		}

		return ctx.JSON(structs.Response{
			Success: true,
			Message: "Saved Join Roles!",
		})
	}
}
