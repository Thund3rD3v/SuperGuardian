package bot

import (
	"fmt"
	"time"

	"github.com/Thund3rD3v/SuperGuardian/database"
	"github.com/Thund3rD3v/SuperGuardian/utils"
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

func HandleVerification(db *gorm.DB) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		config := utils.GetConfig()
		if config.Verification.Enabled {
			if i.ChannelID == config.Verification.ChannelId && !i.Member.User.Bot && i.MessageComponentData().CustomID == "verification_button" {

				// If Author Is Bot Don't Continue
				if i.Member.User.Bot {
					return
				}

				// Check If Member Exists
				memberExists := database.MemberExists(db, i.Member.User.ID)

				// Create Member If Member Does Not Exist
				if !memberExists {
					err := database.CreateMember(db, i.Member.User.ID, 0, 0)
					// Handle Error
					if err != nil {
						fmt.Println("Error Creating Member: ", err.Error())
						utils.RespondWithError(s, i)
						return
					}
				}

				guild, err := s.Guild(i.GuildID)
				if err != nil {
					fmt.Println("Error Getting Guild: " + err.Error())
					utils.RespondWithError(s, i)
					return
				}

				title, err := utils.FormatText(config.Verification.Embeds.Confirmation.Title, s, guild, i.Member.User)
				if err != nil {
					fmt.Println("Error Formatting Verification Confirmation Title: " + err.Error())
					utils.RespondWithError(s, i)
					return
				}

				message, err := utils.FormatText(config.Verification.Embeds.Confirmation.Message, s, guild, i.Member.User)
				if err != nil {
					fmt.Println("Error Formatting Verification Confirmation Message: " + err.Error())
					utils.RespondWithError(s, i)
					return
				}

				embed := discordgo.MessageEmbed{
					Title:       title,
					Description: message,
					Color:       config.Colors.Success,
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
		}
	}
}
