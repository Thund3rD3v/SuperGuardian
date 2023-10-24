package api

import (
	"github.com/Thund3rD3v/SuperGuardian/structs"
	"github.com/Thund3rD3v/SuperGuardian/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
)

func InfoRoute(session *discordgo.Session) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.JSON(structs.Response{
			Success: true,
			Data: structs.AnyData{
				"username":  session.State.User.Username,
				"avatarUrl": session.State.User.AvatarURL("128x128"),
				"color":     session.State.User.AccentColor,
			},
		})
	}
}

func ChannelsRoute(session *discordgo.Session) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		config := utils.GetConfig()

		channels, err := session.GuildChannels(config.GuildId)
		if err != nil {
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Could Not Retrieve Channels",
			})
		}

		return ctx.JSON(structs.Response{
			Success: true,
			Data: structs.AnyData{
				"channels": channels,
			},
		})
	}
}

func RolesRoute(session *discordgo.Session) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		config := utils.GetConfig()

		roles, err := session.GuildRoles(config.GuildId)
		if err != nil {
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Could Not Retrieve Roles",
			})
		}

		return ctx.JSON(structs.Response{
			Success: true,
			Data: structs.AnyData{
				"roles": roles,
			},
		})
	}
}
