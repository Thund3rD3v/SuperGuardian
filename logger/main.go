package logger

import (
	"fmt"
	"log"

	"github.com/Thund3rD3v/SuperGuardian/utils"
	"github.com/bwmarrin/discordgo"
)

func LogMebmerAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	config := utils.GetConfig()

	_, err := s.ChannelMessageSend(config.Logging.ChannelId, fmt.Sprintf("%s, has entered SuperGuardian's Hideout", m.User.Username))
	if err != nil {
		log.Println("Error while sending logging," + err.Error())
		return
	}
}
