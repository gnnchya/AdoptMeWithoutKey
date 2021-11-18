package main

import (
	"github.com/gnnchya/AdoptMe/post/app"
	"github.com/gnnchya/AdoptMe/post/config"
	dynamo "github.com/gnnchya/AdoptMe/post/repository/dynamoDB"
	userService "github.com/gnnchya/AdoptMe/post/service/user/implement"
	"github.com/sirupsen/logrus"
	"log"
)

func newApp(appConfig *config.Config) *app.App {

	dynamoDBClient,_ := dynamo.New(appConfig.DynamoDBRegion)
	user := userService.New(dynamoDBClient)
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