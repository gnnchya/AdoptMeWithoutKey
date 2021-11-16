package dynamoDB

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (repo *Repository) DeleteAdoptionPost (id string) error {
	_, err := repo.Client.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String("AdoptionPost"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	return err
}

func (repo *Repository) DeleteLostPetPost (id string) error {
	_, err := repo.Client.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String("LostPetPost"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	return err
}
