package api

import (
	"encoding/json"

	"github.com/Thund3rD3v/SuperGuardian/structs"
	"github.com/Thund3rD3v/SuperGuardian/utils"
	"github.com/gofiber/fiber/v2"
)

type EditGreetingsBody struct {
	Enabled   bool   `json:"enabled"`
	Title     string `json:"title"`
	Message   string `json:"message"`
	ChannelId string `json:"channelId"`
}

func GreetingsRoute() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		config := utils.GetConfig()

		return ctx.JSON(structs.Response{
			Success: true,
			Data: structs.AnyData{
				"enabled":   config.Greetings.Enabled,
				"title":     config.Greetings.Title,
				"message":   config.Greetings.Message,
				"channelId": config.Greetings.ChannelId,
			},
		})
	}
}

func EditGreetingsRoute() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Parse Json Body
		var body EditGreetingsBody
		raw := ctx.Request().Body()
		err := json.Unmarshal(raw, &body)
		if err != nil {
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Invalid JSON Body",
			})
		}

		config := utils.GetConfig()

		config.Greetings.Enabled = body.Enabled
		config.Greetings.Title = body.Title
		config.Greetings.Message = body.Message
		config.Greetings.ChannelId = body.ChannelId

		err = utils.WriteConfig(config)
		if err != nil {
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Could Not Save New Greetings",
			})
		}

		return ctx.JSON(structs.Response{
			Success: true,
			Message: "Saved Greetings!",
		})
	}
}
