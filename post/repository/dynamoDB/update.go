package dynamoDB

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gnnchya/AdoptMe/post/domain"
	"log"
	"strconv"
)

func (repo *Repository) UpdateAdoptionPost (input domain.CreateAdoptionPostStruct) error {
	item := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(input.ID),
			},
			"animal": {
				M: map[string]*dynamodb.AttributeValue{
					"type": {
						S: aws.String(input.Animal.Type),
					},
					"age": {
						N: aws.String(strconv.Itoa(int(input.Animal.Age))),
					},
					"species": {
						S: aws.String(input.Animal.Species),
					},
					"gender": {
						S: aws.String(input.Animal.Gender),
					},
					"general_information": {
						S: aws.String(input.Animal.GeneralInformation),
					},
					"spay": {
						BOOL: aws.Bool(input.Animal.Spay),
					},
					"image": {
						S: aws.String(input.Animal.Image),
					},
					"medical_condition": {
						S: aws.String(input.Animal.MedicalCondition),
					},
				},
			},
			"adopt": {
				BOOL: aws.Bool(input.Adopt),
			},
			"uid": {
				S: aws.String(input.UID),
			},
			"location": {
				S: aws.String(input.Location),
			},
			"post_at": {
				N: aws.String(strconv.Itoa(int(input.PostAt))),
			},
			"update_at": {
				N: aws.String(strconv.Itoa(int(input.UpdateAt))),
			},
			"delete_at": {
				N: aws.String(strconv.Itoa(int(input.DeleteAt))),
			},
		},
		TableName: aws.String("AdoptionPost"),
	}
	_, err := repo.Client.PutItem(item)

	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}
	return err
}

func (repo *Repository) UpdateLostPetPost (input domain.CreateLostPetPostStruct) error {
	item := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(input.ID),
			},
			"animal": {
				M: map[string]*dynamodb.AttributeValue{
					"type": {
						S: aws.String(input.Animal.Type),
					},
					"age": {
						N: aws.String(strconv.Itoa(int(input.Animal.Age))),
					},
					"species": {
						S: aws.String(input.Animal.Species),
					},
					"gender": {
						S: aws.String(input.Animal.Gender),
					},
					"general_information": {
						S: aws.String(input.Animal.GeneralInformation),
					},
					"spay": {
						BOOL: aws.Bool(input.Animal.Spay),
					},
					"image": {
						S: aws.String(input.Animal.Image),
					},
					"medical_condition": {
						S: aws.String(input.Animal.MedicalCondition),
					},
				},
			},
			"found": {
				BOOL: aws.Bool(input.Found),
			},
			"uid": {
				S: aws.String(input.UID),
			},
			"lost_location": {
				S: aws.String(input.LostLocation),
			},
			"post_at": {
				N: aws.String(strconv.Itoa(int(input.PostAt))),
			},
			"update_at": {
				N: aws.String(strconv.Itoa(int(input.UpdateAt))),
			},
			"delete_at": {
				N: aws.String(strconv.Itoa(int(input.DeleteAt))),
			},
		},
		TableName: aws.String("LostPetPost"),
	}
	_, err := repo.Client.PutItem(item)

	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}
	return err
}