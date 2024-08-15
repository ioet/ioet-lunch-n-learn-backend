package config

type Config struct {
	AppPort string `mapstructure:"APP_PORT"`
}

var envs = []string{
	"APP_PORT",
}
