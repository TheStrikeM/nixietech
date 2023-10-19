package config

type Config struct {
	Env              string      `yaml:"env" env-default:"local"`
	MongoURI         string      `yaml:"mongo-uri" env-default:"mongodb://localhost:27017"`
	TelegramApiToken string      `yaml:"telegram-api-token"`
	Admins           []string    `yaml:"admins"`
	BotMessages      BotMessages `yaml:"bot-messages"`
}
type BotMessages struct {
	InitialHelpMessage      string   `yaml:"initial-help-message"`
	ShowMenuMessage         string   `yaml:"show-menu-message"`
	ClockMenuMessage        string   `yaml:"clock-menu-message"`
	OrderMenuMessage        string   `yaml:"order-menu-message"`
	PermissionPrefixFull    string   `yaml:"permission-prefix-full"`
	PermissionPrefixMinimum string   `yaml:"permission-prefix-minimum"`
	ClockPrefix             string   `yaml:"clock-prefix"`
	OrderPrefix             string   `yaml:"order-prefix"`
	GlobalPrefix            string   `yaml:"global-prefix"`
	PermissionGroups        []string `yaml:"permission-groups"`
}
