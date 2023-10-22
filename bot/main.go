package bot

import (
	"fmt"
	"time"

	"github.com/Thund3rD3v/SuperGuardian/utils"
	"github.com/bwmarrin/discordgo"
)

func Start(session *discordgo.Session, password string) {

	// Add event listeners
	session.AddHandler(SendGreetings)

	// Add required Intents
	session.Identify.Intents = discordgo.IntentsAll

	// Open a websocket connection to Discord and begin listening.
	err := session.Open()
	if err != nil {
		panic("Error while opening connection with discord session," + err.Error())
	}

	// Once Ready Send Owner The Password
	config := utils.GetConfig()

	channel, err := session.UserChannelCreate(config.OwnerId)
	if err != nil {
		fmt.Println("Error while sending password to owner,", err.Error())
	}

	embed := discordgo.MessageEmbed{
		Title:       fmt.Sprintf("%v Dashboard", session.State.User.Username),
		Color:       config.Colors.Main,
		Description: fmt.Sprintf("Your %v password is ``%v``", session.State.User.Username, password),
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	_, err = session.ChannelMessageSendEmbed(channel.ID, &embed)
	if err != nil {
		fmt.Println("Error while sending password to owner,", err.Error())
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
}
