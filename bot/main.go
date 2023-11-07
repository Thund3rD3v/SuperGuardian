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

	// Add required Intents
	session.Identify.Intents = discordgo.IntentsAll

	// Open a websocket connection to Discord and begin listening.
	err := session.Open()
	if err != nil {
		panic("Error While Opening Connection With Discord Session: " + err.Error())
	}

	// Once Ready Send Owner The Password
	config := utils.GetConfig()

	channel, err := session.UserChannelCreate(config.OwnerId)
	if err != nil {
		fmt.Println("Error While Sending Password To Owner: ", err.Error())
	}

	embed := discordgo.MessageEmbed{
		Title:       fmt.Sprintf("%v Dashboard", session.State.User.Username),
		Color:       config.Colors.Main,
		Description: fmt.Sprintf("Your %v password is ```%v```", session.State.User.Username, password),
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	_, err = session.ChannelMessageSendEmbed(channel.ID, &embed)
	if err != nil {
		fmt.Println("Error While Sending Password To Owner: ", err.Error())
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
}
