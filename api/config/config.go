package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func LoadConfig() (*Config, error) {
	var err error

	once.Do(func() {
		var cfg Config

		viper.AddConfigPath("./")
		viper.SetConfigFile(".env")
		err = viper.ReadInConfig()

		if err != nil {
			return
		}

		for _, env := range envs {
			if bindErr := viper.BindEnv(env); bindErr != nil {
				err = bindErr
				return
			}
		}

		if unmarshalErr := viper.Unmarshal(&cfg); unmarshalErr != nil {
			err = unmarshalErr
			return
		}

		if validateErr := validator.New().Struct(&cfg); validateErr != nil {
			err = validateErr
			return
		}

		config = &cfg
	})

	return config, err
}
