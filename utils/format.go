package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func FormatText(text string, s *discordgo.Session, guild *discordgo.Guild, user *discordgo.User) (string, error) {
	var formattedText = text

	if strings.Contains(text, "{user.id}") {
		formattedText = strings.ReplaceAll(text, "{user.id}", user.ID)
	}
	if strings.Contains(text, "{user.name}") {
		formattedText = strings.ReplaceAll(formattedText, "{user.name}", user.Username)
	}
	if strings.Contains(text, "{user.mention}") {
		formattedText = strings.ReplaceAll(formattedText, "{user.mention}", user.Username)
	}
	if strings.Contains(text, "{user.avatar}") {
		formattedText = strings.ReplaceAll(formattedText, "{user.avatar}", user.AvatarURL(""))
	}

	if strings.Contains(text, "{guild.id}") {
		formattedText = strings.ReplaceAll(text, "{guild.id}", guild.ID)
	}
	if strings.Contains(text, "{guild.name}") {
		formattedText = strings.ReplaceAll(formattedText, "{guild.name}", guild.Name)
	}
	if strings.Contains(text, "{guild.icon}") {
		formattedText = strings.ReplaceAll(formattedText, "{guild.icon}", guild.IconURL(("")))
	}
	if strings.Contains(text, "{guild.banner}") {
		formattedText = strings.ReplaceAll(formattedText, "{guild.banner}", guild.BannerURL(""))
	}
	if strings.Contains(text, "{guild.members}") {
		formattedText = strings.ReplaceAll(formattedText, "{guild.members}", strconv.Itoa(GetGuildMemberCount(s, guild.ID)))
	}

	if strings.Contains(text, "{verification.link}") {
		config := GetConfig()
		encryptedId, err := Encrypt(user.ID)
		if err != nil {
			return "", err
		}
		formattedText = strings.ReplaceAll(formattedText, "{verification.link}", fmt.Sprintf("%s/verify?id=%s", config.DashboardUrl, encryptedId))
	}

	return formattedText, nil
}
