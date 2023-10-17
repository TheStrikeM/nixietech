package config

type Config struct {
	Env              string      `yaml:"env" env-default:"local"`
	MongoURI         string      `yaml:"mongo-uri" env-default:"mongodb://localhost:27017"`
	TelegramApiToken string      `yaml:"telegram-api-token"`
	Admins           []string    `yaml:"admins"`
	BotMessages      BotMessages `yaml:"bot-messages"`
}

type BotMessages struct {
	InitialHelp string `yaml:"initial-help"`
}
