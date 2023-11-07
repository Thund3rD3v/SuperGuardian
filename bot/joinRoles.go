package bot

import (
	"fmt"

	"github.com/Thund3rD3v/SuperGuardian/utils"
	"github.com/bwmarrin/discordgo"
)

func JoinRoles(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	// Get Config
	config := utils.GetConfig()

	if config.JoinRoles.Enabled {
		// Loop Through Roles
		for i := 0; i < len(config.JoinRoles.Roles); i++ {
			// Add that role to user
			err := s.GuildMemberRoleAdd(config.GuildId, m.User.ID, config.JoinRoles.Roles[i])
			if err != nil {
				fmt.Println(fmt.Sprintf("Error Adding Role '%v' To Member: ", config.JoinRoles.Roles[i]), err.Error())
			}
		}
	}
}
