package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Thund3rD3v/SuperGuardian/logger"
	"github.com/bwmarrin/discordgo"
)

func Initlize(token string) {
	// Create a new Discord sesson
	client, err := discordgo.New("Bot " + token)
	if err != nil {
		panic("Erorr creating discord session," + err.Error())
	}

	// Add event listeners
	client.AddHandler(messageCreate)
	client.AddHandler(guildMemberAdd)

	// Add required Intents
	client.Identify.Intents = discordgo.IntentGuildMessages | discordgo.IntentGuildMembers

	// Open a websocket connection to Discord and begin listening.
	err = client.Open()
	if err != nil {
		panic("Error while opening connection with discord session," + err.Error())
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	client.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// In this example, we only care about messages that are "ping".
	if m.Content != "ping" {
		return
	}

	// We create the private channel with the user who sent the message.
	channel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		// If an error occurred, we failed to create the channel.
		//
		// Some common causes are:
		// 1. We don't share a server with the user (not possible here).
		// 2. We opened enough DM channels quickly enough for Discord to
		//    label us as abusing the endpoint, blocking us from opening
		//    new ones.
		fmt.Println("error creating channel:", err)
		s.ChannelMessageSend(
			m.ChannelID,
			"Something went wrong while sending the DM!",
		)
		return
	}
	// Then we send the message through the channel we created.
	_, err = s.ChannelMessageSend(channel.ID, "Pong!")
	if err != nil {
		// If an error occurred, we failed to send the message.
		//
		// It may occur either when we do not share a server with the
		// user (highly unlikely as we just received a message) or
		// the user disabled DM in their settings (more likely).
		fmt.Println("error sending DM message:", err)
		s.ChannelMessageSend(
			m.ChannelID,
			"Failed to send you a DM. "+
				"Did you disable DM in your privacy settings?",
		)
	}
}

func guildMemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	logger.LogMebmerAdd(s, m)
}
