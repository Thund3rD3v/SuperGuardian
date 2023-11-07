package api

import (
	"fmt"

	"github.com/Thund3rD3v/SuperGuardian/database"
	"github.com/Thund3rD3v/SuperGuardian/structs"
	"github.com/Thund3rD3v/SuperGuardian/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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

func LeaderboardRoute(db *gorm.DB, session *discordgo.Session) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		config := utils.GetConfig()

		cursor, err := ctx.ParamsInt("cursor")
		if err != nil {
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Could Not Read Cursor",
			})
		}

		members, err := database.GetMemberByCursor(db, cursor, 10)
		if err != nil {
			fmt.Println("Error While Getting Members: ", err)
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Could Not Get Members",
			})
		}

		var newMembers []structs.LeaderboardMember

		for i := 0; i < len(members); i++ {
			member := members[i]
			memberAccount, err := session.User(member.Id)

			if err != nil {
				fmt.Println("Error While Getting Member: ", err)
				continue
			}

			newMembers = append(newMembers, structs.LeaderboardMember{
				Id:       member.Id,
				Username: memberAccount.Username,
				Avatar:   memberAccount.AvatarURL(""),
				Level:    member.Level,
				Xp:       member.Xp,
				MaxXp:    member.Level*config.Levels.XpMultiplier + config.Levels.BaseXp,
			})
		}

		return ctx.JSON(structs.Response{
			Success: true,
			Data: structs.AnyData{
				"members": newMembers,
			},
		})
	}
}
