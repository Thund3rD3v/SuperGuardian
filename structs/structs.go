package structs

import "gorm.io/gorm"

type (
	AnyData map[string]Any
	Any     interface{}
)

type Config struct {
	GuildId string `json:"guildId"`
	OwnerId string `json:"ownerId"`

	Port         int    `json:"port"`
	DashboardUrl string `json:"dashboardUrl"`

	Colors struct {
		Main    int `json:"main"`
		Success int `json:"success"`
		Error   int `json:"error"`
		Warning int `json:"warning"`
	} `json:"colors"`

	Greetings struct {
		Enabled   bool   `json:"enabled"`
		Title     string `json:"title"`
		Message   string `json:"message"`
		ChannelId string `json:"channelId"`
	} `json:"greetings"`

	JoinRoles struct {
		Enabled     bool     `json:"enabled"`
		IncludeBots bool     `json:"includeBots"`
		Roles       []string `json:"roles"`
	} `json:"joinRoles"`

	Levels struct {
		Enabled      bool   `json:"enabled"`
		ChannelId    string `json:"channelId"`
		Title        string `json:"title"`
		Message      string `json:"message"`
		CoolDown     int    `json:"coolDown"`
		MinXp        int    `json:"minXp"`
		MaxXp        int    `json:"maxXp"`
		BaseXp       int    `json:"baseXp"`
		XpMultiplier int    `json:"xpMultiplier"`
	} `json:"levels"`

	Verification struct {
		Enabled bool `json:"enabled"`
		Embeds  struct {
			Main struct {
				Title   string `json:"title"`
				Message string `json:"message"`
			} `json:"main"`
			Confirmation struct {
				Title   string `json:"title"`
				Message string `json:"message"`
			} `json:"confirmation"`
		} `json:"embeds"`
		VerifiedText string   `json:"verifiedText"`
		ChannelId    string   `json:"channelId"`
		Roles        []string `json:"roles"`
	} `json:"verification"`

	ContentFiltering struct {
		Enabled          bool     `json:"enabled"`
		ProfanityChecker bool     `json:"profanityChecker"`
		BlackList        []string `json:"blacklist"`
	} `json:"contentFiltering"`
}

type Response struct {
	Success bool    `json:"success"`
	Message string  `json:"message,omitempty"`
	Data    AnyData `json:"data,omitempty"`
}

type HCaptchaResponse struct {
	Success     bool     `json:"success"`
	ChallengeTs string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes"`
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
