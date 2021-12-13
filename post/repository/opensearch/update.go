package opensearch

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gnnchya/AdoptMeWithoutKey/post/domain"
	"io"
	"log"
)

func (repo *Repository) UpdateAdoptionPost(ctx context.Context, title *domain.CreateAdoptionPostStruct) error {
	if state, err := repo.CheckExistAdoptionPostID(ctx, title.ID); err != nil {
		return err
	} else if state == false {
		return fmt.Errorf("this ID does not exist")
	}
	buf, err := BuildUpdateAdoptionPostRequest(title)
	if err != nil {
		return err
	}
	res, err := repo.Client.Update(
		repo.AdoptionPostIndex, title.ID, &buf,
		repo.Client.Update.WithContext(ctx),
		repo.Client.Update.WithPretty())
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return err
		} else {
			// Print the response status and error information.
			log.Println("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	return err
}

func (repo *Repository) UpdateLostPetPost(ctx context.Context, title *domain.CreateLostPetPostStruct) error {
	if state, err := repo.CheckExistLostPetPostID(ctx, title.ID); err != nil {
		return err
	} else if state == false {
		return fmt.Errorf("this ID does not exist")
	}
	buf, err := BuildUpdateLostPetPostRequest(title)
	if err != nil {
		return err
	}
	res, err := repo.Client.Update(
		repo.LostPetPostIndex, title.ID, &buf,
		repo.Client.Update.WithContext(ctx),
		repo.Client.Update.WithPretty())
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return err
		} else {
			// Print the response status and error information.
			log.Println("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	return err
}
