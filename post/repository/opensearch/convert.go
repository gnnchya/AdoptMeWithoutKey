package opensearch

import (
	"encoding/json"
	"github.com/gnnchya/AdoptMeWithoutKey/post/domain"
)

func IntoAdoptionPostStruct(r map[string]interface{}) []domain.CreateAdoptionPostStruct {
	var temp domain.CreateAdoptionPostStruct
	var result []domain.CreateAdoptionPostStruct
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		s := hit.(map[string]interface{})["_source"]
		t, _ := json.Marshal(s)
		_ = json.Unmarshal(t, &temp)
		result = append(result, temp)
	}
	return result
}

func IntoLostPetPostStruct(r map[string]interface{}) []domain.CreateLostPetPostStruct {
	var temp domain.CreateLostPetPostStruct
	var result []domain.CreateLostPetPostStruct
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		s := hit.(map[string]interface{})["_source"]
		t, _ := json.Marshal(s)
		_ = json.Unmarshal(t, &temp)
		result = append(result, temp)
	}
	return result
}
