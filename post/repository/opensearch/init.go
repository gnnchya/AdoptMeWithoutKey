package opensearch

import (
	"crypto/tls"
	"github.com/opensearch-project/opensearch-go"
	"log"
	"net"
	"net/http"
	"time"
)

type Repository struct {
	Client            *opensearch.Client
	LostPetPostIndex  string
	AdoptionPostIndex string
}

func New(uri string, username string, password string, AdoptionPostIndex string, LostPetPostIndex string) (repo *Repository, err error) {

	log.Println(uri, username, password)
	openSearchCli, err := opensearch.NewClient(opensearch.Config{
		Addresses: []string{
			uri,
		},
		Username: username,
		Password: password,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS11,
			},
		},
	})
	repo = &Repository{}
	repo.Client = openSearchCli
	repo.LostPetPostIndex = LostPetPostIndex
	repo.AdoptionPostIndex = AdoptionPostIndex
	return repo, err
}
