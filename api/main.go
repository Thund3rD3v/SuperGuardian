package api

import (
	"encoding/json"
	"fmt"

	"github.com/Thund3rD3v/SuperGuardian/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func Start(session *discordgo.Session, db *gorm.DB, password string) {
	config := utils.GetConfig()

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
	app.Get("/greetings/info", ValidatePasswordMiddleware(password), GreetingsRoute())
	app.Patch("/greetings/edit", ValidatePasswordMiddleware(password), EditGreetingsRoute())

	// Join Roles
	app.Get("/join-roles/info", ValidatePasswordMiddleware(password), JoinRolesRoute())
	app.Patch("/join-roles/edit", ValidatePasswordMiddleware(password), EditJoinRolesRoute())

	// Levels
	app.Get("/levels/info", ValidatePasswordMiddleware(password), LevelsRoute())
	app.Patch("/levels/edit", ValidatePasswordMiddleware(password), EditLevelsRoute())

	// Embed Sender
	app.Post("/send-embed", ValidatePasswordMiddleware(password), SendEmbedRoute(session))

	// Verification
	app.Post("/verify", VerifyRoute(session))

	fmt.Printf("Server Available At http://localhost:%v\n", config.Port)
	fmt.Println("Dashboard Password:", password)

	app.Listen(fmt.Sprintf("127.0.0.1:%v", config.Port))
}
