package dynamoDB

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gnnchya/AdoptMeWithoutKey/post/domain"
	"log"
	"strconv"
)

func (repo *Repository) Register(input domain.UserStruct) error {
	item := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"uid": {
				S: aws.String(input.UID),
			},
			"username": {
				S: aws.String(input.Username),
			},
			"name": {
				S: aws.String(input.Name),
			},
			"gender": {
				S: aws.String(input.Gender),
			},
			"address": {
				S: aws.String(input.Address),
			},
			"email": {
				S: aws.String(input.Email),
			},
			"birthdate": {
				N: aws.String(strconv.Itoa(int(input.Birthdate))),
			},
		},
		TableName: aws.String("AdoptMe-User"),
	}
	_, err := repo.Client.PutItem(item)

	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}
	return err
}
