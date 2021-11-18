package dynamoDB

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gnnchya/AdoptMe/post/domain"
	"github.com/gnnchya/AdoptMe/post/service/user/userin"
)

func (repo *Repository) ReadAllAdoptionPostByField (input userin.ReadAllAdoptionPostInputStruct) ([]domain.CreateAdoptionPostStruct, error) {
	q := &dynamodb.ScanInput{
		TableName: aws.String("AdoptionPost"),
		Limit: aws.Int64(int64(input.Limit)),
		ExpressionAttributeNames: map[string]*string{
			"#animal": aws.String("animal"),
			"#type": aws.String("type"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":a": {

				S: aws.String(input.Field),
			},
		},
		FilterExpression: aws.String("#animal.#type = :a"),
	}

	result, err := repo.Client.Scan(q)
	var posts []domain.CreateAdoptionPostStruct
	for _, i := range result.Items {
		post := domain.CreateAdoptionPostStruct{}

		err = dynamodbattribute.UnmarshalMap(i, &post)

		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}

	return posts,err
}

func (repo *Repository) ReadAllLostPetPostByField (input userin.ReadAllLostPetPostInputStruct) ([]domain.CreateLostPetPostStruct, error) {
	q := &dynamodb.ScanInput{
		TableName: aws.String("LostPetPost"),
		Limit: aws.Int64(int64(input.Limit)),
		ExpressionAttributeNames: map[string]*string{
			"#animal": aws.String("animal"),
			"#type": aws.String("type"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":a": {

				S: aws.String(input.Field),
			},
		},
		FilterExpression: aws.String("#animal.#type = :a"),
	}

	result, err := repo.Client.Scan(q)
	var posts []domain.CreateLostPetPostStruct
	for _, i := range result.Items {
		post := domain.CreateLostPetPostStruct{}

		err = dynamodbattribute.UnmarshalMap(i, &post)

		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}

	return posts,err
}

func (repo *Repository) ReadAllAdoptionPost (input userin.ReadAllAdoptionPostInputStruct) ([]domain.CreateAdoptionPostStruct, error) {
	q := &dynamodb.ScanInput{
		TableName: aws.String("AdoptionPost"),
		Limit: aws.Int64(int64(input.Limit)),
	}

	result, err := repo.Client.Scan(q)
	var posts []domain.CreateAdoptionPostStruct
	for _, i := range result.Items {
		post := domain.CreateAdoptionPostStruct{}

		err = dynamodbattribute.UnmarshalMap(i, &post)

		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}

	return posts,err
}

func (repo *Repository) ReadAllLostPetPost (input userin.ReadAllLostPetPostInputStruct) ([]domain.CreateLostPetPostStruct, error) {
	q := &dynamodb.ScanInput{
		TableName: aws.String("LostPetPost"),
		Limit: aws.Int64(int64(input.Limit)),
	}

	result, err := repo.Client.Scan(q)
	var posts []domain.CreateLostPetPostStruct
	for _, i := range result.Items {
		post := domain.CreateLostPetPostStruct{}

		err = dynamodbattribute.UnmarshalMap(i, &post)

		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}

	return posts,err
}
