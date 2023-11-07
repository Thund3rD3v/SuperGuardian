package api

import (
	"encoding/json"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

const Port = 3000

func Start(session *discordgo.Session, db *gorm.DB, password string) {
	app := fiber.New(fiber.Config{
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		DisableStartupMessage: true,
	})

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization", // Add the headers your client may send
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS",
	}))

	app.Post("/login", LoginRoute(password))

	// Info
	app.Get("/info", InfoRoute(session))
	app.Get("/info/channels", ValidatePasswordMiddleware(password), ChannelsRoute(session))
	app.Get("/info/roles", ValidatePasswordMiddleware(password), RolesRoute(session))
	app.Get("/info/leaderboard/:cursor", ValidatePasswordMiddleware(password), LeaderboardRoute(db, session))

	// Greetings
	app.Get("/greetings/info", ValidatePasswordMiddleware(password), GreetingsRoute(session))
	app.Patch("/greetings/edit", ValidatePasswordMiddleware(password), EditGreetingsRoute(session))

	// Join Roles
	app.Get("/join-roles/info", ValidatePasswordMiddleware(password), JoinRolesRoute(session))
	app.Patch("/join-roles/edit", ValidatePasswordMiddleware(password), EditJoinRolesRoute(session))

	// Levels
	app.Get("/levels/info", ValidatePasswordMiddleware(password), LevelsRoute(session))
	app.Patch("/levels/edit", ValidatePasswordMiddleware(password), EditLevelsRoute(session))

	// Embed Sender
	app.Post("/send-embed", ValidatePasswordMiddleware(password), SendEmbedRoute(session))

	fmt.Printf("Server Available At http://localhost:%v\n", Port)
	fmt.Println("Dashboard Password:", password)

	app.Listen(fmt.Sprintf("127.0.0.1:%v", Port))
}
