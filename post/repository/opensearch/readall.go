package opensearch

import (
	"context"
	"github.com/gnnchya/AdoptMeWithoutKey/post/domain"
	"github.com/gnnchya/AdoptMeWithoutKey/post/service/user/userin"
)

func (repo *Repository) ReadAllAdoptionPost(page int, size int, ctx context.Context) ([]domain.CreateAdoptionPostStruct, error) {
	q, err := repo.queryAdoptionPost(ctx, buildViewAllRequest(page, size, "adopt"))
	result := IntoAdoptionPostStruct(q)
	return result, err
}

func (repo *Repository) ReadAllLostPetPost(page int, size int, ctx context.Context) ([]domain.CreateLostPetPostStruct, error) {
	q, err := repo.queryLostPetPost(ctx, buildViewAllRequest(page, size, "found"))
	result := IntoLostPetPostStruct(q)
	return result, err
}

func (repo *Repository) ReadAllAdoptionPostByType(input userin.ReadAllAdoptionPostInputStruct, ctx context.Context) ([]domain.CreateAdoptionPostStruct, error) {
	q, err := repo.queryAdoptionPost(ctx, buildReadAllByAnimalTypeRequest(input.Page, input.Limit, input.Field, "adopt"))
	result := IntoAdoptionPostStruct(q)
	return result, err
}

func (repo *Repository) ReadAllLostPetPostByType(input userin.ReadAllLostPetPostInputStruct, ctx context.Context) ([]domain.CreateLostPetPostStruct, error) {
	q, err := repo.queryLostPetPost(ctx, buildReadAllByAnimalTypeRequest(input.Page, input.Limit, input.Field, "found"))
	result := IntoLostPetPostStruct(q)
	return result, err
}
