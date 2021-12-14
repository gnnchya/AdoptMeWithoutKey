package opensearch

import (
	"context"
	"fmt"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"io"
	"log"
)

func (repo *Repository) DeleteAdoptionPost(ctx context.Context, id string) error {
	if state, err := repo.CheckExistAdoptionPostID(ctx, id); err != nil {
		return err
	} else if state == false {
		return fmt.Errorf("this ID does not exist")
	}
	req := opensearchapi.DeleteRequest{
		Index:      repo.AdoptionPostIndex,
		DocumentID: id,
		Refresh:    "true",
	}

	res, err := req.Do(ctx, repo.Client)
	log.Println("delete  :", res)
	if err != nil {
		log.Println("Error getting response: %s", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	return err
}

func (repo *Repository) DeleteLostPetPost(ctx context.Context, id string) error {
	if state, err := repo.CheckExistLostPetPostID(ctx, id); err != nil {
		return err
	} else if state == false {
		return fmt.Errorf("this ID does not exist")
	}
	req := opensearchapi.DeleteRequest{
		Index:      repo.LostPetPostIndex,
		DocumentID: id,
		Refresh:    "true",
	}

	res, err := req.Do(ctx, repo.Client)
	log.Println("delete  :", res)
	if err != nil {
		log.Println("Error getting response: %s", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	return err
}
