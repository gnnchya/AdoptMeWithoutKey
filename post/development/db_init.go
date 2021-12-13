package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gnnchya/AdoptMeWithoutKey/post/config"
	"github.com/gnnchya/AdoptMeWithoutKey/post/domain"
	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var Animal = []domain.AnimalStruct{
	{"bunny", 3, "Netherland Dwarf", "male", "eat a lot", true, "https://adoptme-pic-storage.s3.ap-southeast-1.amazonaws.com/bunny/4-Week-Old_Netherlands_Dwarf_Rabbit.jpeg", "vaccinate"},
	{"cat", 4, "Persian", "male", "like to play with people", true, "https://adoptme-pic-storage.s3.ap-southeast-1.amazonaws.com/cat/250px-Gatto_europeo4.jpeg", "very healthy"},
	{"dog", 1, "Bulldog", "female", "very friendly", false, "https://adoptme-pic-storage.s3.ap-southeast-1.amazonaws.com/dogs/90.jpeg", "no medical condition"},
	{"dog", 2, "Poodle", "male", "like to eat bonchon", true, "https://adoptme-pic-storage.s3.ap-southeast-1.amazonaws.com/dogs/file-20200310-61148-vllmgm.jpg", "insomnia"},
	{"dog", 7, "Golden Retriever", "female", "need to sleep with people", false, "https://adoptme-pic-storage.s3.ap-southeast-1.amazonaws.com/dogs/1628056310_dogdrinking-1200x800.png.jpeg", "good health"},
	{"cat", 9, "Birman", "female", "like to play with people", true, "https://adoptme-pic-storage.s3.ap-southeast-1.amazonaws.com/cat/3500.jpeg", "very healthy"},
	{"bunny", 2, "European", "male", "eat a lot", true, "https://adoptme-pic-storage.s3.ap-southeast-1.amazonaws.com/bunny/BUNNY_-Snowflake.jpeg", "vaccinate"},
	{"dog", 5, "Labrador", "female", "should provide a lot of time for her", true, "https://adoptme-pic-storage.s3.ap-southeast-1.amazonaws.com/dogs/golden-retriever-dog-breed-info.jpg", "very healthy"},
	{"cat", 6, "Persian", "male", "doesn’t like to shower", true, "https://adoptme-pic-storage.s3.ap-southeast-1.amazonaws.com/cat/images.jpeg", "has every month check up"},
	{"cat", 8, "Maine Coon", "male", "doesn’t like to eat veggie", false, "https://adoptme-pic-storage.s3.ap-southeast-1.amazonaws.com/cat/pexels-evg-culture-1170986.jpeg", "unhealthy"},
	{"cat", 3, "Maine Coon", "male", "eat a lot", true, "https://adoptme-pic-storage.s3.ap-southeast-1.amazonaws.com/cat/Screen+Shot+2021-11-15+at+15.24.07.png", "unhealthy"},
	{"dog", 4, "Labrador", "male", "like to play with people", true, "https://adoptme-pic-storage.s3.ap-southeast-1.amazonaws.com/dogs/imageForEntry18-8on.jpg", "very healthy"},
	{"dog", 1, "Poodle", "female", "very friendly", false, "https://adoptme-pic-storage.s3.ap-southeast-1.amazonaws.com/dogs/images-2.jpeg", "no medical condition"},
	{"dog", 2, "Siberian Husky", "male", "like to eat bonchon", true, "https://adoptme-pic-storage.s3.ap-southeast-1.amazonaws.com/dogs/images-3.jpeg", "insomnia"},
	{"bunny", 7, "European", "female", "need to sleep with people", false, "https://adoptme-pic-storage.s3.ap-southeast-1.amazonaws.com/bunny/easter-bunny-1615238549.jpeg", "good health"},
	{"dog", 9, "Golden Retriever", "female", "like to play with people", true, "https://adoptme-pic-storage.s3.ap-southeast-1.amazonaws.com/dogs/images.jpeg", "no medical condition"},
	{"cat", 2, "Bengal cat", "male", "eat a lot", true, "https://adoptme-pic-storage.s3.ap-southeast-1.amazonaws.com/cat/istockphoto-1266826333-612x612.jpg", "vaccinate"},
	{"dog", 5, "Siberian Husky", "female", "should provide a lot of time for her", true, "https://adoptme-pic-storage.s3.ap-southeast-1.amazonaws.com/dogs/Miniature-Poodle-walking-in-the-grass-in-the-fall.jpg", "very healthy"},
	{"dog", 6, "Labrador", "male", "doesn’t like to shower", true, "https://adoptme-pic-storage.s3.ap-southeast-1.amazonaws.com/dogs/Siberian-Husky-standing-outdoors-in-the-winter.jpg", "vaccinate"},
	{"bunny", 8, "Mini Lop", "male", "doesn’t like to eat veggie", false, "https://adoptme-pic-storage.s3.ap-southeast-1.amazonaws.com/bunny/how_much_is_a_pet_bunny.jpeg", "unhealthy"},
}

var Post = []domain.CreateAdoptionPostStruct{
	{"abc111", Animal[0], false, "94aafbb6-a846-4849-ba7a-79c01b473c21", "Bangkok", 1636361248, 0, 0},
	{"abc113", Animal[1], true, "94aafbb6-a846-4849-ba7a-79c01b473c21", "Chiangmai", 1634368848, 0, 0},
	{"abc114", Animal[2], false, "94aafbb6-a846-4849-ba7a-79c01b473c21", "Bangkok", 1636548848, 0, 0},
	{"abc115", Animal[3], true, "94aafbb6-a846-4849-ba7a-79c01b473c21", "Chonburi", 1631168848, 0, 0},
	{"abc116", Animal[4], false, "592b6739-b2a5-40fb-92b9-a2c19e56de5c", "Bangkok", 1622368848, 0, 0},
	{"abc117", Animal[5], false, "592b6739-b2a5-40fb-92b9-a2c19e56de5c", "Phuket", 1636368338, 0, 0},
	{"abc118", Animal[6], true, "592b6739-b2a5-40fb-92b9-a2c19e56de5c", "Bangkok", 123658848, 0, 0},
	{"abc119", Animal[7], false, "592b6739-b2a5-40fb-92b9-a2c19e56de5c", "Phuket", 1126368848, 0, 0},
	{"abc120", Animal[8], false, "b77ffffd-3f65-4293-a4b0-2b8c95909f6f", "Bangkok", 1236368848, 0, 0},
	{"abc121", Animal[9], false, "b77ffffd-3f65-4293-a4b0-2b8c95909f6f", "Phuket", 1636368811, 0, 0},
	{"abc122", Animal[10], true, "b77ffffd-3f65-4293-a4b0-2b8c95909f6f", "Nan", 1636368821, 0, 0},
	{"abc123", Animal[11], false, "b77ffffd-3f65-4293-a4b0-2b8c95909f6f", "Bangkok", 1636368845, 0, 0},
	{"abc124", Animal[12], false, "53496045-6f94-430c-81fe-84793623ccc5", "Krabi", 1622368848, 0, 0},
	{"abc125", Animal[13], false, "53496045-6f94-430c-81fe-84793623ccc5", "Bangkok", 1636322848, 0, 0},
	{"abc126", Animal[14], false, "53496045-6f94-430c-81fe-84793623ccc5", "Rayong", 1632268848, 0, 0},
	{"abc127", Animal[15], false, "53496045-6f94-430c-81fe-84793623ccc5", "Krabi", 1636368848, 0, 0},
	{"abc128", Animal[16], false, "819a739b-011a-4d35-8c6f-0adbd066ce7", "Chiangmai", 1633368848, 0, 0},
	{"abc129", Animal[17], false, "819a739b-011a-4d35-8c6f-0adbd066ce7", "Bangkok", 1625368848, 0, 0},
	{"abc130", Animal[18], false, "819a739b-011a-4d35-8c6f-0adbd066ce7", "Chiangmai", 1636368848, 0, 0},
	{"abc112", Animal[19], true, "819a739b-011a-4d35-8c6f-0adbd066ce7", "Bangkok", 1636323848, 0, 0},
}

var lostPost = []domain.CreateLostPetPostStruct{
	{"bbc111", Animal[0], true, "6fc526a4-0695-43f2-885e-35b2118ed4ff", "Bangkok", 1636361248, 0, 0},
	{"bbc113", Animal[1], false, "6fc526a4-0695-43f2-885e-35b2118ed4ff", "Chiangmai", 1634368848, 0, 0},
	{"bbc114", Animal[2], false, "6fc526a4-0695-43f2-885e-35b2118ed4ff", "Bangkok", 1636548848, 0, 0},
	{"bbc115", Animal[3], false, "6fc526a4-0695-43f2-885e-35b2118ed4ff", "Chonburi", 1631168848, 0, 0},
	{"bbc116", Animal[4], false, "0ec59c93-d7a9-4a0b-998f-8ff39e9e1826", "Bangkok", 1622368848, 0, 0},
	{"bbc117", Animal[5], false, "0ec59c93-d7a9-4a0b-998f-8ff39e9e1826", "Phuket", 1636368338, 0, 0},
	{"bbc118", Animal[6], true, "0ec59c93-d7a9-4a0b-998f-8ff39e9e1826", "Bangkok", 123658848, 0, 0},
	{"bbc119", Animal[7], false, "0ec59c93-d7a9-4a0b-998f-8ff39e9e1826", "Phuket", 1126368848, 0, 0},
	{"bbc120", Animal[8], false, "efe878ef-5b2d-4ada-9e08-c6f23c51cfb4", "Bangkok", 1236368848, 0, 0},
	{"bbc121", Animal[9], false, "efe878ef-5b2d-4ada-9e08-c6f23c51cfb4", "Nan", 1636368811, 0, 0},
	{"bbc122", Animal[10], true, "efe878ef-5b2d-4ada-9e08-c6f23c51cfb4", "Rayong", 1636368821, 0, 0},
	{"bbc123", Animal[11], false, "efe878ef-5b2d-4ada-9e08-c6f23c51cfb4", "Samutprakarn", 1636368845, 0, 0},
	{"bbc124", Animal[12], false, "8d13c2e7-4ec5-4367-8f5d-18e61de1e066", "Krabi", 1622368848, 0, 0},
	{"bbc125", Animal[13], false, "8d13c2e7-4ec5-4367-8f5d-18e61de1e066", "Bangkok", 1636322848, 0, 0},
	{"bbc126", Animal[14], false, "8d13c2e7-4ec5-4367-8f5d-18e61de1e066", "Krabi", 1632268848, 0, 0},
	{"bbc127", Animal[15], true, "8d13c2e7-4ec5-4367-8f5d-18e61de1e066", "Bangkok", 1636368848, 0, 0},
	{"bbc128", Animal[16], false, "bd115d76-9c00-4621-86b2-dc7fb1f567c7", "Bangkok", 1633368848, 0, 0},
	{"bbc129", Animal[17], false, "bd115d76-9c00-4621-86b2-dc7fb1f567c7", "Krabi", 1625368848, 0, 0},
	{"bbc130", Animal[18], true, "bd115d76-9c00-4621-86b2-dc7fb1f567c7", "Bangkok", 1636368848, 0, 0},
	{"bbc112", Animal[19], false, "bd115d76-9c00-4621-86b2-dc7fb1f567c7", "Bangkok", 1636323848, 0, 0},
}

func create(es *opensearch.Client) {
	//var c *gin.Context
	for k, v := range Post {
		fmt.Println("p", Post[k].ID)
		out, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}

		var b strings.Builder
		b.WriteString(string(out))

		// Set up the request object.
		req := opensearchapi.IndexRequest{
			Index:      "adoptionpost",
			DocumentID: Post[k].ID,
			Body:       strings.NewReader(b.String()),
			Refresh:    "true",
		}

		// Perform the request with the client.
		res, err := req.Do(nil, es)
		//res, err := es.Info()

		if err != nil {
			log.Println("Error getting response: %s", err)
		}
		err = res.Body.Close()
		if err != nil {
			return
		}

	}
	for k, v := range lostPost {
		fmt.Println("p", lostPost[k].ID)
		out, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}

		var b strings.Builder
		b.WriteString(string(out))

		// Set up the request object.
		req := opensearchapi.IndexRequest{
			Index:      "lostpetpost",
			DocumentID: lostPost[k].ID,
			Body:       strings.NewReader(b.String()),
			Refresh:    "true",
		}

		// Perform the request with the client.
		res, err := req.Do(nil, es)
		//res, err := es.Info()

		if err != nil {
			log.Println("Error getting response: %s", err)
		}
		err = res.Body.Close()
		if err != nil {
			return
		}

	}

}

func main() {

	configs := config.Get()

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(configs.DynamoDBRegion),
		Credentials: credentials.NewEnvCredentials(),
	})

	if err != nil {
		log.Println("Got error get new session: %s", err)
	}

	svc := dynamodb.New(sess)

	item := &dynamodb.PutItemInput{}
	for _, input := range Post {
		item = &dynamodb.PutItemInput{
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
		_, err = svc.PutItem(item)

		if err != nil {
			log.Println("Got error calling GetItem: %s", err)
		}
	}

	for _, input := range lostPost {
		item = &dynamodb.PutItemInput{
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
		_, err = svc.PutItem(item)

		if err != nil {
			log.Println("Got error calling GetItem: %s", err)
		}
	}

	es, err := opensearch.NewClient(opensearch.Config{
		Addresses: []string{
			configs.OpenSearchDBEndpoint,
		},
		Username: configs.OpenSearchDBUsername,
		Password: configs.OpenSearchDBPassword,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS11,
			},
		},
	})
	if err != nil {
		log.Println("Error creating the client: %s", err)
	}
	create(es)

}
