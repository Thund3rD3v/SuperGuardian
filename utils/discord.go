package utils

import (
	"encoding/json"

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
