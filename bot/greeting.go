package bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/Thund3rD3v/SuperGuardian/utils"
	"github.com/bwmarrin/discordgo"
)

func SendGreetings(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	// Get Config
	config := utils.GetConfig()

	// Check If Greetings Is Enabled
	if config.Greetings.Enabled {
		memberCount := utils.GetGuildMemberCount(s, m.GuildID)

		// Create Embed
		embed := discordgo.MessageEmbed{
			Title:       fmt.Sprintf("%v, Has Joined!", m.User.Username),
			Color:       config.Colors.Main,
			Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: m.User.AvatarURL("")},
			Description: fmt.Sprintf("%v\n\n**Members:** %v", strings.Replace(config.Greetings.Message, "${tag}", m.Mention(), -1), memberCount),
			Timestamp:   time.Now().Format(time.RFC3339),
		}

		// Send Embed
		_, err := s.ChannelMessageSendEmbed(config.Greetings.ChannelId, &embed)
		if err != nil {
			fmt.Println("Error Greeting Member: " + err.Error())
			return
		}
	}
}
