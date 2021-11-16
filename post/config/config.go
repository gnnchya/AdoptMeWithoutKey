package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	AppName string `env:"APP_NAME" envDefault:"gogo_blueprint"`

	DynamoDBRegion string `env:"aws_region" envDefault:""`
	DynamoDBAccessKey string `env:"aws_access_key_id" envDefault:""`
	DynamoDBSecretKey string `env:"aws_secret_access_key_id" envDefault:""`
	DynamoDBToken string `env:"aws_token_id" envDefault:""`

}

func Get() *Config {
	appConfig := &Config{}
	if err := env.Parse(appConfig); err != nil {
		panic(err)
	}

	return appConfig
}
