package bot

import (
	"fmt"
	"time"

	"github.com/Thund3rD3v/SuperGuardian/utils"
	"github.com/bwmarrin/discordgo"
)

func SendGreetings(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	// Get Config
	config := utils.GetConfig()

	// Check If Greetings Is Enabled
	if config.Greetings.Enabled {
		guild, err := s.Guild(m.GuildID)
		if err != nil {
			fmt.Println("Error Getting Guild: " + err.Error())
			return
		}

		title, err := utils.FormatText(config.Greetings.Title, s, guild, m.User)
		if err != nil {
			fmt.Println("Error Formatting Greetings Title: " + err.Error())
			return
		}

		message, err := utils.FormatText(config.Greetings.Message, s, guild, m.User)
		if err != nil {
			fmt.Println("Error Formatting Greetings Message: " + err.Error())
			return
		}

		// Create Embed
		embed := discordgo.MessageEmbed{
			Title:       title,
			Color:       config.Colors.Main,
			Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: m.User.AvatarURL("")},
			Description: message,
			Timestamp:   time.Now().Format(time.RFC3339),
		}

		// Send Embed
		_, err = s.ChannelMessageSendEmbed(config.Greetings.ChannelId, &embed)
		if err != nil {
			fmt.Println("Error Greeting Member: " + err.Error())
			return
		}
	}
}
