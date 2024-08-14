package config

type Config struct {
	APP_PORT string `mapstructure:"APP_PORT"`
}

var envs = []string{
	"APP_PORT",
}
