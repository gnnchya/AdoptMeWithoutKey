package opensearch

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
)

func (repo *Repository) queryAdoptionPost(ctx context.Context, buf bytes.Buffer) (map[string]interface{}, error) {
	ops := repo.Client
	var r map[string]interface{}
	res, err := ops.Search(
		ops.Search.WithContext(ctx),
		ops.Search.WithIndex(repo.AdoptionPostIndex),
		ops.Search.WithBody(&buf),
		ops.Search.WithTrackTotalHits(true),
		ops.Search.WithPretty(),
	)
	if err != nil {
		log.Println("Error getting response: %s", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Println("Error parsing the response body: %s", err)
		} else {
			log.Println("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Println("Error parsing the response body: %s", err)
	}
	return r, err
}

func (repo *Repository) queryLostPetPost(ctx context.Context, buf bytes.Buffer) (map[string]interface{}, error) {
	ops := repo.Client
	var r map[string]interface{}
	res, err := ops.Search(
		ops.Search.WithContext(ctx),
		ops.Search.WithIndex(repo.LostPetPostIndex),
		ops.Search.WithBody(&buf),
		ops.Search.WithTrackTotalHits(true),
		ops.Search.WithPretty(),
	)
	if err != nil {
		log.Println("Error getting response: %s", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Println("Error parsing the response body: %s", err)
		} else {
			log.Println("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Println("Error parsing the response body: %s", err)
	}
	return r, err
}
