package opensearch

import (
	"context"
	"github.com/gnnchya/AdoptMeWithoutKey/post/domain"
	"github.com/gnnchya/AdoptMeWithoutKey/post/service/user/userin"
)

func (repo *Repository) SearchAdoptionPost(input userin.SearchInputStruct, ctx context.Context) ([]domain.CreateAdoptionPostStruct, error) {
	q, err := repo.queryAdoptionPost(ctx, buildSearchRequest(input.Page, input.Size, input.Keyword))
	if err != nil {
		return nil, err
	}
	result := IntoAdoptionPostStruct(q)
	return result, err
}

func (repo *Repository) SearchLostPetPost(input userin.SearchInputStruct, ctx context.Context) ([]domain.CreateLostPetPostStruct, error) {
	q, err := repo.queryLostPetPost(ctx, buildSearchRequest(input.Page, input.Size, input.Keyword))
	if err != nil {
		return nil, err
	}
	result := IntoLostPetPostStruct(q)
	return result, err
}
