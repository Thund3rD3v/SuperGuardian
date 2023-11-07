package bot

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
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

					embed := discordgo.MessageEmbed{
						Title:       fmt.Sprintf("%v, Has Advanced To Level %v", m.Author.Username, member.Level),
						Color:       config.Colors.Main,
						Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: m.Author.AvatarURL("")},
						Description: strings.Replace(strings.Replace(config.Levels.Message, "${name}", m.Author.Username, -1), "${level}", strconv.Itoa(member.Level), -1),
						Timestamp:   time.Now().Format(time.RFC3339),
					}

					message := discordgo.MessageSend{
						Content: m.Author.Mention(),
						Embeds:  []*discordgo.MessageEmbed{&embed},
					}

					// Send Level Up Message
					s.ChannelMessageSendComplex(config.Levels.ChannelId, &message)
				}

				// Add Random Exp Between Min And Max Values Set In Config
				member.Xp += reward

				database.UpdateMember(db, &member)
			}
		}
	}
}
