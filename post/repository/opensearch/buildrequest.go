package opensearch

import (
	"bytes"
	"encoding/json"
	"github.com/gnnchya/AdoptMeWithoutKey/post/domain"
	"log"
	"time"
)

func buildSearchRequest(page int, size int, keyword string) bytes.Buffer {
	var buf bytes.Buffer
	from := (page - 1) * size
	query := map[string]interface{}{
		"from": from,
		"size": size,
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query": "*" + keyword + "*",
				"fields": []interface{}{
					"animal.type",
					"animal.general_information",
					"animal.species",
					"animal.medical_condition",
				},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	return buf
}

func BuildCheckIDRequest(id string) (buf bytes.Buffer, err error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"_id": id,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return buf, err
	}
	return buf, err
}

func BuildUpdateAdoptionPostRequest(t *domain.CreateAdoptionPostStruct) (buf bytes.Buffer, err error) {
	query := map[string]interface{}{
		"doc": map[string]interface{}{
			"animal":   t.Animal,
			"adopt":    t.Adopt,
			"location": t.Location,
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return buf, err
	}
	return buf, err
}

func BuildUpdateLostPetPostRequest(t *domain.CreateLostPetPostStruct) (buf bytes.Buffer, err error) {
	query := map[string]interface{}{
		"doc": map[string]interface{}{
			"animal":   t.Animal,
			"found":    t.Found,
			"location": t.LostLocation,
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return buf, err
	}
	return buf, err
}

func buildViewRequest(id string) bytes.Buffer {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"_id": id,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	return buf
}

func buildViewAllRequest(page int, size int, field string) bytes.Buffer {
	var buf bytes.Buffer
	from := (page - 1) * size
	query := map[string]interface{}{
		"from": from,
		"size": size,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{"match": map[string]interface{}{
						field: false,
					}},
					{"match": map[string]interface{}{
						"delete_at": 0,
					}},
				},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	return buf
}

func buildReadAllByAnimalTypeRequest(page int, size int, keyword string, field string) bytes.Buffer {
	var buf bytes.Buffer
	from := (page - 1) * size
	query := map[string]interface{}{
		"from": from,
		"size": size,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{

					{"match": map[string]interface{}{
						"animal.type": keyword,
					}},
					{"match": map[string]interface{}{
						field: false,
					}},
					{"match": map[string]interface{}{
						"delete_at": 0,
					}},
				},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	return buf
}

func BuildSoftDeletePostRequest() (buf bytes.Buffer, err error) {
	query := map[string]interface{}{
		"doc": map[string]interface{}{
			"delete_at": time.Now().Unix(),
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return buf, err
	}
	return buf, err
}
