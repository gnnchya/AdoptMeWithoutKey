package main

import (
	//"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	//"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func main() {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Endpoint: aws.String(aws.Config{}),
	})

	svc := dynamodb.New(sess)

	//if err != nil {
	//	log.Fatalf("Got error get new session: %s", err)
	//}

	//tableName := "book"
	//bookId := "069"
	//name := "Genshim"


	//item := &dynamodb.PutItemInput{
	//	Item: map[string]*dynamodb.AttributeValue{
	//		"id": {
	//			S: aws.String("5"),
	//		},
	//		"name": {
	//			S: aws.String("No One You Know"),
	//		},
	//		"location": {
	//			S: aws.String("Call Me Today"),
	//		},
	//	},
	//	TableName:              aws.String("post"),
	//}
	item := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("123"),
			},
			"animal": {
				M: map[string]*dynamodb.AttributeValue{
					"type": {
						S: aws.String("dog"),
					},
					"age": {
						N: aws.String("123.345341324123423"),
					},
					"species": {
						S: aws.String("bulldog"),
					},
					"gender": {
						S: aws.String("male"),
					},
					"general_information": {
						S: aws.String("black white dot"),
					},
					"spay": {
						BOOL: aws.Bool(false),
					},
					"image": {
						S: aws.String("image.png"),
					},
					"medical_condition": {
						S: aws.String("no problem"),
					},
				},
			},
			"adopt": {
				S: aws.String("no problem"),
			},
			"uid": {
				S: aws.String("no problem"),
			},
			"location": {
				S: aws.String("no problem"),
			},
			"post_at": {
				S: aws.String("no problem"),
			},
			"update_at": {
				S: aws.String("no problem"),
			},
			"delete_at": {
				S: aws.String("no problem"),
			},
		},
		TableName: aws.String("AdoptionPost"),
	}
	_, err = svc.PutItem(item)

	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	//if result.Item == nil {
	//	msg := "Could not find '" + name + "'"
	//	fmt.Println(msg)
	//}
	//
	//// var item string
	//err = dynamodbattribute.Unmarshal(result.Item["bookId"], &item.bookId)
	//err = dynamodbattribute.Unmarshal(result.Item["name"], &item.name)


	//if err != nil {
	//	panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	//}

	fmt.Println("Found item:")
	// fmt.Println(item)
	//fmt.Println("bookId:  ", item.bookId)
	//fmt.Println("name: ", item.name)
}
