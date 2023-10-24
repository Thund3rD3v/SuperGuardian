package structs

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
}

type Response struct {
	Success bool    `json:"success"`
	Message string  `json:"message,omitempty"`
	Data    AnyData `json:"data,omitempty"`
}
