package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Thund3rD3v/SuperGuardian/api"
	"github.com/Thund3rD3v/SuperGuardian/bot"
	"github.com/Thund3rD3v/SuperGuardian/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// Load Env File
	err := godotenv.Load()
	if err != nil {
		panic("Error while loading .env file," + err.Error())
	}

	// Generate Password
	password := utils.GeneratePassword(64)

	// Create a new Discord sesson
	session, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		panic("Erorr creating discord session," + err.Error())
	}

	// Start The Bot
	bot.Start(session, password)

	// Start The Api
	go api.Start(session, password)

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	session.Close()
}
