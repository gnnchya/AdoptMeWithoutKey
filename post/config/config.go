package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	AppName string `env:"APP_NAME" envDefault:"gogo_blueprint"`

	DynamoDBRegion string `env:"AWS_REGION" envDefault:""`
	DynamoDBAccessKey string `env:"AWS_ACCESS_KEY" envDefault:""`
	DynamoDBSecretKey string `env:"AWS_SECRET_ACCESS_KEY" envDefault:""`
	DynamoDBToken string `env:"AWS_SESSION_TOKEN" envDefault:""`

}

func Get() *Config {
	appConfig := &Config{}
	if err := env.Parse(appConfig); err != nil {
		panic(err)
	}

	return appConfig
}
