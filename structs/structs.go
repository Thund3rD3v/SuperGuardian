package structs

import "gorm.io/gorm"

type (
	AnyData map[string]Any
	Any     interface{}
)

type Config struct {
	GuildId string `json:"guildId"`
	OwnerId string `json:"ownerId"`

	Colors struct {
		Main    int `json:"main"`
		Success int `json:"success"`
		Error   int `json:"error"`
		Warning int `json:"warning"`
	} `json:"colors"`

	Greetings struct {
		Enabled   bool   `json:"enabled"`
		Message   string `json:"message"`
		ChannelId string `json:"channelId"`
	} `json:"greetings"`

	JoinRoles struct {
		Enabled     bool     `json:"enabled"`
		IncludeBots bool     `json:"includeBots"`
		Roles       []string `json:"roles"`
	}

	Levels struct {
		Enabled      bool   `json:"enabled"`
		ChannelId    string `json:"channelId"`
		Message      string `json:"message"`
		CoolDown     int    `json:"coolDown"`
		MinXp        int    `json:"minXp"`
		MaxXp        int    `json:"maxXp"`
		BaseXp       int    `json:"baseXp"`
		XpMultiplier int    `json:"xpMultiplier"`
	}
}

type Response struct {
	Success bool    `json:"success"`
	Message string  `json:"message,omitempty"`
	Data    AnyData `json:"data,omitempty"`
}

type Member struct {
	gorm.Model
	Id            string
	Level         int
	Xp            int
	MessageSentAt int64
}

type LeaderboardMember struct {
	Id       string `json:"id" gorm:"not null"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Level    int    `json:"level"`
	Xp       int    `json:"xp"`
	MaxXp    int    `json:"maxXp"`
}
