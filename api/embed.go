package api

import (
	"encoding/json"
	"fmt"

	"github.com/Thund3rD3v/SuperGuardian/structs"
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
)

type SendEmbedBody struct {
	ChannelId string                 `json:"channelId"`
	Content   string                 `json:"content,omitempty"`
	Embed     discordgo.MessageEmbed `json:"embed"`
}

func SendEmbedRoute(session *discordgo.Session) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Parse Json Body
		var body SendEmbedBody
		raw := ctx.Request().Body()
		err := json.Unmarshal(raw, &body)
		if err != nil {
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Invalid JSON Body",
			})
		}

		message := discordgo.MessageSend{
			Content: body.Content,
			Embeds:  []*discordgo.MessageEmbed{&body.Embed},
		}

		// Send Embed
		_, err = session.ChannelMessageSendComplex(body.ChannelId, &message)
		if err != nil {
			fmt.Println("Error Sending Embed: " + err.Error())
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Error Sending Embed!",
			})
		}

		return ctx.JSON(structs.Response{
			Success: true,
			Message: "Embed Sent!",
		})
	}
}
