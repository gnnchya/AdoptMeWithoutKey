package dynamoDB

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

type Repository struct {
	Client *dynamodb.DynamoDB
}

func New(region string) (repo *Repository, err error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	//sess, err = session.NewSession(&aws.Config{
	//	Region:      aws.String(region),
	//	Credentials: credentials.NewEnvCredentials(),
	//})
	log.Println("error cant connect to aws session:", err)
	svc := dynamodb.New(sess)
	repo = &Repository{}
	repo.Client = svc
	return repo, err
}
