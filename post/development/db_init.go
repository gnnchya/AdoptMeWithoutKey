package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gnnchya/AdoptMe/post/config"
	"github.com/gnnchya/AdoptMe/post/domain"
	"log"
	"strconv"
)

func main() {

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
		{"abc111", Animal[0], true, "u111", "Bangkok", 1636361248, 0, 0},
		{"abc113", Animal[1], true, "u113", "Chiangmai", 1634368848, 0, 0},
		{"abc114", Animal[2], true, "u114", "Bangkok", 1636548848, 0, 0},
		{"abc115", Animal[3], true, "u115", "Chonburi", 1631168848, 0, 0},
		{"abc116", Animal[4], true, "u116", "Bangkok", 1622368848, 0, 0},
		{"abc117", Animal[5], true, "u117", "Phuket", 1636368338, 0, 0},
		{"abc118", Animal[6], true, "u118", "Bangkok", 123658848, 0, 0},
		{"abc119", Animal[7], true, "u119", "Phuket", 1126368848, 0, 0},
		{"abc120", Animal[8], true, "u120", "Bangkok", 1236368848, 0, 0},
		{"abc121", Animal[9], true, "u123", "Phuket", 1636368811, 0, 0},
		{"abc122", Animal[10], true, "u122", "Nan", 1636368821, 0, 0},
		{"abc123", Animal[11], true, "u123", "Bangkok", 1636368845, 0, 0},
		{"abc124", Animal[12], true, "u124", "Krabi", 1622368848, 0, 0},
		{"abc125", Animal[13], true, "u125", "Bangkok", 1636322848, 0, 0},
		{"abc126", Animal[14], true, "u126", "Rayong", 1632268848, 0, 0},
		{"abc127", Animal[15], true, "u127", "Krabi", 1636368848, 0, 0},
		{"abc128", Animal[16], true, "u128", "Chiangmai", 1633368848, 0, 0},
		{"abc129", Animal[17], true, "u129", "Bangkok", 1625368848, 0, 0},
		{"abc130", Animal[18], true, "u130", "Chiangmai", 1636368848, 0, 0},
		{"abc112", Animal[19], true, "u112", "Bangkok", 1636323848, 0, 0},
	}

	var lostPost = []domain.CreateLostPetPostStruct{
		{"bbc111", Animal[0], true, "u011", "Bangkok", 1636361248, 0, 0},
		{"bbc113", Animal[1], false, "u013", "Chiangmai", 1634368848, 0, 0},
		{"bbc114", Animal[2], true, "u014", "Bangkok", 1636548848, 0, 0},
		{"bbc115", Animal[3], true, "u015", "Chonburi", 1631168848, 0, 0},
		{"bbc116", Animal[4], false, "u016", "Bangkok", 1622368848, 0, 0},
		{"bbc117", Animal[5], false, "u017", "Phuket", 1636368338, 0, 0},
		{"bbc118", Animal[6], true, "u018", "Bangkok", 123658848, 0, 0},
		{"bbc119", Animal[7], false, "u019", "Phuket", 1126368848, 0, 0},
		{"bbc120", Animal[8], false, "u020", "Bangkok", 1236368848, 0, 0},
		{"bbc121", Animal[9], false, "u023", "Nan", 1636368811, 0, 0},
		{"bbc122", Animal[10], true, "u022", "Rayong", 1636368821, 0, 0},
		{"bbc123", Animal[11], false, "u023", "Samutprakarn", 1636368845, 0, 0},
		{"bbc124", Animal[12], true, "u024", "Krabi", 1622368848, 0, 0},
		{"bbc125", Animal[13], false, "u025", "Bangkok", 1636322848, 0, 0},
		{"bbc126", Animal[14], true, "u026", "Krabi", 1632268848, 0, 0},
		{"bbc127", Animal[15], true, "u027", "Bangkok", 1636368848, 0, 0},
		{"bbc128", Animal[16], false, "u028", "Bangkok", 1633368848, 0, 0},
		{"bbc129", Animal[17], true, "u029", "Krabi", 1625368848, 0, 0},
		{"bbc130", Animal[18], true, "u030", "Bangkok", 1636368848, 0, 0},
		{"bbc112", Animal[19], false, "u012", "Bangkok", 1636323848, 0, 0},
	}

	configs := config.Get()

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(configs.DynamoDBRegion),
		Credentials: credentials.NewEnvCredentials(),
	})

	if err != nil {
		log.Fatalf("Got error get new session: %s", err)
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
			log.Fatalf("Got error calling GetItem: %s", err)
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
			log.Fatalf("Got error calling GetItem: %s", err)
		}
	}

}
