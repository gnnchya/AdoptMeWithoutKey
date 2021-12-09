package opensearch

import (
	"context"
	"fmt"
	"github.com/gnnchya/AdoptMeWithoutKey/post/domain"
)

func (repo *Repository) ReadAdoptionPost(id string, ctx context.Context) (result domain.CreateAdoptionPostStruct, err error) {
	if found, err := repo.CheckExistAdoptionPostID(ctx, id); found == false {
		return result, fmt.Errorf("error : ID does not exist")
	} else if err != nil {
		return result, err
	}
	q, err := repo.queryAdoptionPost(ctx, buildViewRequest(id))
	fmt.Println(q)
	results := IntoAdoptionPostStruct(q)
	fmt.Println(results[0])
	return results[0], err
}

func (repo *Repository) ReadLostPetPost(id string, ctx context.Context) (result domain.CreateLostPetPostStruct, err error) {
	if found, err := repo.CheckExistLostPetPostID(ctx, id); found == false {
		return result, fmt.Errorf("error : ID does not exist")
	} else if err != nil {
		return result, err
	}
	q, err := repo.queryLostPetPost(ctx, buildViewRequest(id))
	fmt.Println(q)
	results := IntoLostPetPostStruct(q)
	fmt.Println(results[0])
	return results[0], err
}
