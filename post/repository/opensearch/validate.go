package opensearch

import "context"

func (repo *Repository) CheckExistAdoptionPostID(ctx context.Context, id string) (bool, error) {
	buf, err := BuildCheckIDRequest(id)
	if err != nil {
		return false, err
	}
	result, err := repo.queryAdoptionPost(ctx, buf)
	found := int((result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"]).(float64))
	if found == 0 {
		return false, nil
	}
	return true, err
}

func (repo *Repository) CheckExistLostPetPostID(ctx context.Context, id string) (bool, error) {
	buf, err := BuildCheckIDRequest(id)
	if err != nil {
		return false, err
	}
	result, err := repo.queryLostPetPost(ctx, buf)
	found := int((result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"]).(float64))
	if found == 0 {
		return false, nil
	}
	return true, err
}
