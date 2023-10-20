package main

import (
	"os"

	"github.com/Thund3rD3v/SuperGuardian/bot"
	"github.com/joho/godotenv"
)

func main() {
	// Load Env File
	err := godotenv.Load()
	if err != nil {
		panic("Error while loading .env file," + err.Error())
	}

	botToken := os.Getenv("BOT_TOKEN")

	bot.Initlize(botToken)
}
