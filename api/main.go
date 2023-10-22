package api

import (
	"encoding/json"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const Port = 3000

func Start(session *discordgo.Session, password string) {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization", // Add the headers your client may send
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS",
	}))

	app.Post("/login", LoginRoute(password))

	app.Get("/info", InfoRoute(session))
	app.Get("/info/channels", ValidatePasswordMiddleware(password), ChannelsRoute(session))

	app.Get("/greetings/info", ValidatePasswordMiddleware(password), GreetingsRoute(session))
	app.Patch("/greetings/edit", ValidatePasswordMiddleware(password), EditGreetingsRoute(session))

	fmt.Printf("Server Available At http://localhost:%v", Port)

	app.Listen(fmt.Sprintf("127.0.0.1:%v", Port))
}
