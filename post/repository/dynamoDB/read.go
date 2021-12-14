package dynamoDB

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gnnchya/AdoptMeWithoutKey/post/domain"
	"log"
)

func (repo *Repository) ReadAdoptionPostByID(id string) (domain.CreateAdoptionPostStruct, error) {
	log.Println("id:", id)
	result, err := repo.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("AdoptionPost"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	item := domain.CreateAdoptionPostStruct{}
	// var item string
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	log.Println("item:", item)
	return item, err
}

func (repo *Repository) ReadLostPetPostByID(id string) (domain.CreateLostPetPostStruct, error) {
	log.Println("id:", id)
	result, err := repo.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("LostPetPost"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	item := domain.CreateLostPetPostStruct{}
	log.Println("item:", item)

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	return item, err
}

func (repo *Repository) ReadUserByID(uid string) (domain.UserStruct, error) {
	result, err := repo.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("AdoptMe-User"),
		Key: map[string]*dynamodb.AttributeValue{
			"uid": {
				S: aws.String(uid),
			},
		},
	})
	item := domain.UserStruct{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	return item, err
}
