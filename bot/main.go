package bot

import (
	"fmt"
	"time"

	"github.com/Thund3rD3v/SuperGuardian/utils"
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

func Start(session *discordgo.Session, db *gorm.DB, password string) {

	// Add event listeners
	session.AddHandler(SendGreetings)
	session.AddHandler(JoinRoles)
	session.AddHandler(Levels(db))
	session.AddHandler(HandleCommands(db))
	session.AddHandler(HandleVerification(db))

	// Add required Intents
	session.Identify.Intents = discordgo.IntentsAll

	// Open a websocket connection to Discord and begin listening.
	err := session.Open()
	if err != nil {
		panic("Error While Opening Connection With Discord Session: " + err.Error())
	}

	// Register the commands
	RegisterCommands(session)

	// Once ready send owner the password
	config := utils.GetConfig()
	channel, err := session.UserChannelCreate(config.OwnerId)
	if err != nil {
		fmt.Println("Error While Sending Password To Owner: ", err.Error())
	}
	embed := discordgo.MessageEmbed{
		Title:       fmt.Sprintf("🖥️ %v Dashboard Access", session.State.User.Username),
		Color:       config.Colors.Main,
		Description: fmt.Sprintf("Here's your dashboard password: ```%v```", password),
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	// Send verification embed
	if config.Verification.Enabled {
		err = SendVerificationEmbed(session)
		if err != nil {
			fmt.Println("Error While Sending Verification Embed: ", err.Error())
		}
	}

	_, err = session.ChannelMessageSendEmbed(channel.ID, &embed)
	if err != nil {
		fmt.Println("Error While Sending Password To Owner: ", err.Error())
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
}

func SendVerificationEmbed(s *discordgo.Session) error {
	config := utils.GetConfig()

	guild, err := s.Guild(config.GuildId)
	if err != nil {
		return err
	}

	user, err := s.User(config.OwnerId)
	if err != nil {
		return err
	}

	title, err := utils.FormatText(config.Verification.Embeds.Main.Title, s, guild, user)
	if err != nil {
		return err
	}

	message, err := utils.FormatText(config.Verification.Embeds.Main.Message, s, guild, user)
	if err != nil {
		return err
	}

	embed2 := &discordgo.MessageEmbed{
		Title:       title,
		Description: message,
		Color:       config.Colors.Main,
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	button := discordgo.Button{
		Label:    "Verify",
		Emoji:    discordgo.ComponentEmoji{Name: "✅"},
		Style:    discordgo.PrimaryButton,
		CustomID: "verification_button",
	}

	msgSend := &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{embed2},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{button},
			},
		},
	}

	_, err = s.ChannelMessageSendComplex(config.Verification.ChannelId, msgSend)
	if err != nil {
		// fmt.Println("Error While Sending Verification Embed: ", err.Error())
		return err
	}

	return nil
}
