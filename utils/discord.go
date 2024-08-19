package utils

import (
	"encoding/json"
	"time"

	"github.com/bwmarrin/discordgo"
)

func GetGuildMemberCount(s *discordgo.Session, gId string) int {
	// Request the guild
	guildJson, err := s.Request("GET", discordgo.EndpointGuild(gId)+"?with_counts=true", nil)
	if err != nil {
		return 0
	}

	guildDiscord := &discordgo.Guild{}

	// Unmarshal the bytes json into a interface
	err = json.Unmarshal(guildJson, guildDiscord)

	if err != nil {
		return 0
	}

	return guildDiscord.ApproximateMemberCount
}

func RespondWithError(s *discordgo.Session, i *discordgo.InteractionCreate) {
	config := GetConfig()

	embed := discordgo.MessageEmbed{
		Title:       "Unexpected error",
		Color:       config.Colors.Error,
		Description: "There was an unexpected error while processing your request",
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{&embed},
			Flags:  64,
		},
	})
}
