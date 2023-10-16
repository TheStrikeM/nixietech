package config

type Config struct {
	Env      string `yaml:"env" env-default:"local"`
	MongoURI string `yaml:"mongo-uri" env-default:"mongodb://localhost:27017"`
}
