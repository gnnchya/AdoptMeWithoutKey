package main

import (
	"github.com/gnnchya/AdoptMeWithoutKey/post/app"
	"github.com/gnnchya/AdoptMeWithoutKey/post/config"
	dynamo "github.com/gnnchya/AdoptMeWithoutKey/post/repository/dynamoDB"
	openSearch "github.com/gnnchya/AdoptMeWithoutKey/post/repository/opensearch"
	userService "github.com/gnnchya/AdoptMeWithoutKey/post/service/user/implement"
	"github.com/sirupsen/logrus"
	"log"
)

func newApp(appConfig *config.Config) *app.App {
	openSearchClient, _ := openSearch.New(appConfig.OpenSearchDBEndpoint, appConfig.OpenSearchDBUsername, appConfig.OpenSearchDBPassword, "adoptionpost", "lostpetpost")
	dynamoDBClient, _ := dynamo.New(appConfig.DynamoDBRegion)
	user := userService.New(dynamoDBClient, openSearchClient)
	return app.New(user)
}

func panicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

const (
	NETWORK = "tcp"
)

func setupLog() *logrus.Logger {
	lr := logrus.New()
	lr.SetFormatter(&logrus.JSONFormatter{})

	return lr
}
