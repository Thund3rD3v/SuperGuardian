package bot

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/Thund3rD3v/SuperGuardian/database"
	"github.com/Thund3rD3v/SuperGuardian/utils"
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

func Levels(db *gorm.DB) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		config := utils.GetConfig()

		if config.Levels.Enabled {
			// If Author Is Bot Don't Continue
			if m.Author.Bot {
				return
			}

			// Check If Member Exists
			memberExists := database.MemberExists(db, m.Author.ID)

			// Create Member If Member Does Not Exist
			if !memberExists {
				err := database.CreateMember(db, m.Author.ID, 0, 0)
				// Handle Error
				if err != nil {
					fmt.Println("Error Creating Member: ", err.Error())
					return
				}
			}

			// Get Member
			member, err := database.GetMember(db, m.Author.ID)
			if err != nil {
				fmt.Println("Error While Getting Member: ", err.Error())
				return
			}

			// Check Cool Down
			messageSentAt := time.UnixMilli(member.MessageSentAt)
			if int(time.Since(messageSentAt).Seconds()) > config.Levels.CoolDown {
				// Update Database
				member.MessageSentAt = time.Now().UnixMilli()

				reward := rand.Intn(config.Levels.MaxXp-config.Levels.MinXp+1) + config.Levels.MinXp

				// Check If You Should Progress To Next Level
				if member.Xp+reward >= member.Level*config.Levels.XpMultiplier+config.Levels.BaseXp {
					reward = 0
					member.Xp = 0
					member.Level += 1

					guild, err := s.Guild(m.GuildID)
					if err != nil {
						fmt.Println("Error Getting Guild: " + err.Error())
						return
					}

					title, err := utils.FormatText(config.Greetings.Title, s, guild, m.Author)
					if err != nil {
						fmt.Println("Error Formatting Levels Title: " + err.Error())
						return
					}

					message, err := utils.FormatText(config.Greetings.Message, s, guild, m.Author)
					if err != nil {
						fmt.Println("Error Formatting Levels Message: " + err.Error())
						return
					}

					embed := discordgo.MessageEmbed{
						Title:       title,
						Color:       config.Colors.Main,
						Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: m.Author.AvatarURL("")},
						Description: message,
						Timestamp:   time.Now().Format(time.RFC3339),
					}

					embedMessage := discordgo.MessageSend{
						Content: m.Author.Mention(),
						Embeds:  []*discordgo.MessageEmbed{&embed},
					}

					// Send Level Up Message
					s.ChannelMessageSendComplex(config.Levels.ChannelId, &embedMessage)
				}

				// Add Random Exp Between Min And Max Values Set In Config
				member.Xp += reward

				database.UpdateMember(db, &member)
			}
		}
	}
}

func LevelCommand(db *gorm.DB, s *discordgo.Session, i *discordgo.InteractionCreate) {
	config := utils.GetConfig()

	if config.Levels.Enabled {
		// Check If Member Exists
		memberExists := database.MemberExists(db, i.Member.User.ID)

		// Create Member If Member Does Not Exist
		if !memberExists {
			err := database.CreateMember(db, i.Message.Author.ID, 0, 0)
			// Handle Error
			if err != nil {
				fmt.Println("Error Creating Member: ", err.Error())
				utils.RespondWithError(s, i)
				return
			}
		}

		// Get Member
		member, err := database.GetMember(db, i.Member.User.ID)
		if err != nil {
			fmt.Println("Error While Getting Member: ", err.Error())
			utils.RespondWithError(s, i)
			return
		}

		fields := []*discordgo.MessageEmbedField{
			{
				Name:  "Level",
				Value: strconv.Itoa(member.Level),
			},
			{
				Name:  "Xp",
				Value: fmt.Sprintf("%v/%v", member.Xp, member.Level*config.Levels.XpMultiplier+config.Levels.BaseXp),
			},
		}

		embed := discordgo.MessageEmbed{
			Title:     fmt.Sprintf("%v's Level", i.Member.User.Username),
			Color:     config.Colors.Main,
			Thumbnail: &discordgo.MessageEmbedThumbnail{URL: i.Member.User.AvatarURL("")},
			Fields:    fields,
			Timestamp: time.Now().Format(time.RFC3339),
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{&embed},
			},
		})
	} else {
		embed := discordgo.MessageEmbed{
			Title:       "Disabled Module",
			Color:       config.Colors.Error,
			Description: "This module is disabled, you can enable it through the dashboard",
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
