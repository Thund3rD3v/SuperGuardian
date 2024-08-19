package bot

import (
	"github.com/Thund3rD3v/SuperGuardian/utils"
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

func HandleCommands(db *gorm.DB) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand || i.Member.User.Bot {
			return
		}

		switch i.ApplicationCommandData().Name {
		case "level":
			LevelCommand(db, s, i)
		}
	}
}

func RegisterCommands(s *discordgo.Session) {
	config := utils.GetConfig()

	s.ApplicationCommandBulkOverwrite(s.State.User.ID, config.GuildId, []*discordgo.ApplicationCommand{
		{
			Name:        "level",
			Description: "Use this command to get your level",
		},
	})
}
