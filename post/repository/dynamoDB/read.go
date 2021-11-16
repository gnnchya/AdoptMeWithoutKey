package dynamoDB

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gnnchya/AdoptMe/post/domain"
)

func (repo *Repository) ReadAdoptionPostByID (id string) (domain.CreateAdoptionPostStruct, error){
	fmt.Println("id:",id)
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
	fmt.Println("item:",item)
	return item, err
}

func (repo *Repository) ReadLostPetPostByID (id string) (domain.CreateLostPetPostStruct, error){
	fmt.Println("id:",id)
	result, err := repo.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("LostPetPost"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	item := domain.CreateLostPetPostStruct{}
	fmt.Println("item:",item)

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	return item, err
}
