package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type AnimalStruct struct {
	Type               string `json:"type"`
	Age                int64  `json:"age"`
	Species            string `json:"species"`
	Gender             string `json:"gender"`
	GeneralInformation string `json:"general_information"`
	Spay               bool   `json:"spay"`
	Image              string `json:"image"`
	MedicalCondition   string `json:"medical_condition"`
}

type CreateAdoptionPostStruct struct {
	ID       string       `json:"id"`
	Animal   AnimalStruct `json:"animal"`
	Adopt    bool         `json:"adopt"`
	UID      string       `json:"uid"`
	Location string       `json:"location"`
	PostAt   int64        `json:"post_at"`
	UpdateAt int64        `json:"update_at"`
	DeleteAt int64        `json:"delete_at"`
}

func main() {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewEnvCredentials(),
	})

	svc := dynamodb.New(sess)

	//id := "5"
	//name := "nametest"
	result, err := svc.Scan(&dynamodb.ScanInput{
		TableName: aws.String("AdoptionPost"),
		ExpressionAttributeNames: map[string]*string{
			"#animal": aws.String("animal"),
			"#type": aws.String("type"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":a": {

				S: aws.String("dog"),
			},
		},
		FilterExpression: aws.String("#animal.#type = :a"),
	})

	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	fmt.Println(result)
	var posts []CreateAdoptionPostStruct

	for _, i := range result.Items {
		post := CreateAdoptionPostStruct{}

		err = dynamodbattribute.UnmarshalMap(i, &post)
		fmt.Println("asd", post)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}

	//if result.Item == nil {
	//	msg := "Could not find '" + name + "'"
	//	fmt.Println(msg)
	//}

	//item := Item{}
	// var item string
	//err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	//err = dynamodbattribute.Unmarshal(result.Item["name"], &item.name)
	//err = dynamodbattribute.Unmarshal(result.Item["location"], &item.location)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("Found item:")
	//fmt.Println(posts)
	//fmt.Println("ID:  ", item.ID)
	//fmt.Println("name: ", item.name)
	//fmt.Println("name: ", item.name)

}
