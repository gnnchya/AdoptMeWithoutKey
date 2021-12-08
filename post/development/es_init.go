package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gnnchya/AdoptMeWithoutKey/post/config"
	"github.com/gnnchya/AdoptMeWithoutKey/post/domain"
	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

//func initDb(uri string, username string, password string) (*opensearch.Client, error) {
//	cfg := opensearch.Config{
//		Addresses: []string{
//			uri,
//		},
//		Username: username,
//		Password: password,
//	}
//	es, err := opensearch.NewClient(cfg)
//	return es, err
//}

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
			log.Fatalf("Error getting response: %s", err)
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
			log.Fatalf("Error getting response: %s", err)
		}
		err = res.Body.Close()
		if err != nil {
			return
		}

	}

}

func main() {
	configs := config.Get()
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
		log.Fatalf("Error creating the client: %s", err)
	}
	create(es)

}
