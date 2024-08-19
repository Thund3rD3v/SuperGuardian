package api

import (
	"encoding/json"
	"fmt"

	"github.com/Thund3rD3v/SuperGuardian/structs"
	"github.com/Thund3rD3v/SuperGuardian/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
)

type VerifyBody struct {
	Id                string `json:"id"`
	VerificationToken string `json:"verificationToken"`
}

func VerifyRoute(s *discordgo.Session) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		config := utils.GetConfig()

		// Check If Verification Is Disabled
		if !config.Verification.Enabled {
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Verification Is Disabled",
			})
		}

		// Parse Json Body
		var body VerifyBody
		raw := ctx.Request().Body()
		err := json.Unmarshal(raw, &body)
		if err != nil {
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Invalid JSON Body",
			})
		}

		verified, err := utils.VerifyCaptcha(body.VerificationToken)
		if err != nil {
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "An Unexpected Error Has Occurred",
			})
		}

		if verified {
			memberId, err := utils.Decrypt(body.Id)
			if err != nil {
				fmt.Println("Error Decrypting id: ", err.Error())
				return ctx.JSON(structs.Response{
					Success: false,
					Message: "Invalid id",
				})
			}

			for i := 0; i < len(config.Verification.Roles); i++ {
				err := s.GuildMemberRoleAdd(config.GuildId, memberId, config.Verification.Roles[i])
				if err != nil {
					fmt.Println(fmt.Sprintf("Error Adding Role '%v' To Member: ", config.JoinRoles.Roles[i]), err.Error())
					return ctx.JSON(structs.Response{
						Success: false,
						Message: "An Unexpected Error Has Occurred",
					})
				}
			}

			guild, err := s.Guild(config.GuildId)
			if err != nil {
				fmt.Println("Error While Getting Guild On Verification: ", err.Error())
				return ctx.JSON(structs.Response{
					Success: true,
					Message: config.Verification.VerifiedText,
				})
			}

			user, err := s.User(memberId)
			if err != nil {
				fmt.Println("Error While Getting User On Verification: ", err.Error())
				return ctx.JSON(structs.Response{
					Success: true,
					Message: config.Verification.VerifiedText,
				})
			}

			text, err := utils.FormatText(config.Verification.VerifiedText, s, guild, user)
			if err != nil {
				fmt.Println("Error While Getting Formatting Text On Verification: ", err.Error())
				return ctx.JSON(structs.Response{
					Success: true,
					Message: config.Verification.VerifiedText,
				})
			}

			return ctx.JSON(structs.Response{
				Success: true,
				Message: text,
			})
		}

		return ctx.JSON(structs.Response{
			Success: false,
			Message: "Invalid Hcaptcha",
		})
	}
}
