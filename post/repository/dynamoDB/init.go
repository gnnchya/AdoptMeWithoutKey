package dynamoDB

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Repository struct{
	Client 	*dynamodb.DynamoDB
}

func New(region string, accessKey string, secretKey string, token string)(repo *Repository, err error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey ,token),

	})

	svc := dynamodb.New(sess)
	repo = &Repository{}
	repo.Client = svc
	return repo , err
}
