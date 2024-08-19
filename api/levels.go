package api

import (
	"encoding/json"

	"github.com/Thund3rD3v/SuperGuardian/structs"
	"github.com/Thund3rD3v/SuperGuardian/utils"
	"github.com/gofiber/fiber/v2"
)

type EditLevelsBody struct {
	Enabled      bool   `json:"enabled"`
	ChannelId    string `json:"channelId"`
	Title        string `json:"title"`
	Message      string `json:"message"`
	CoolDown     int    `json:"coolDown"`
	MinXp        int    `json:"minXp"`
	MaxXp        int    `json:"maxXp"`
	BaseXp       int    `json:"baseXp"`
	XpMultiplier int    `json:"xpMultiplier"`
}

func LevelsRoute() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		config := utils.GetConfig()

		return ctx.JSON(structs.Response{
			Success: true,
			Data: structs.AnyData{
				"enabled":      config.Levels.Enabled,
				"channelId":    config.Levels.ChannelId,
				"title":        config.Levels.Title,
				"message":      config.Levels.Message,
				"coolDown":     config.Levels.CoolDown,
				"minXp":        config.Levels.MinXp,
				"maxXp":        config.Levels.MaxXp,
				"baseXp":       config.Levels.BaseXp,
				"xpMultiplier": config.Levels.XpMultiplier,
			},
		})
	}
}

func EditLevelsRoute() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Parse Json Body
		var body EditLevelsBody
		raw := ctx.Request().Body()
		err := json.Unmarshal(raw, &body)
		if err != nil {
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Invalid JSON Body",
			})
		}

		config := utils.GetConfig()

		config.Levels.Enabled = body.Enabled
		config.Levels.ChannelId = body.ChannelId
		config.Levels.Title = body.Title
		config.Levels.Message = body.Message
		config.Levels.CoolDown = body.CoolDown
		config.Levels.MinXp = body.MinXp
		config.Levels.MaxXp = body.MaxXp
		config.Levels.BaseXp = body.BaseXp
		config.Levels.XpMultiplier = body.XpMultiplier

		err = utils.WriteConfig(config)
		if err != nil {
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Could Not Save Levels",
			})
		}

		return ctx.JSON(structs.Response{
			Success: true,
			Message: "Saved Levels!",
		})
	}
}
