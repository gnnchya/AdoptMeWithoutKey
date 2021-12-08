package implement

import (
	"context"
	"github.com/gnnchya/AdoptMeWithoutKey/post/domain"
	"github.com/gnnchya/AdoptMeWithoutKey/post/service/user/userin"
)

func (impl *implementation) SearchAdoptionPost(ctx context.Context, input *userin.SearchInputStruct) ([]domain.CreateAdoptionPostStruct, error) {
	a, err := impl.openSearch.SearchAdoptionPost(*input, ctx)
	if err != nil {
		return a, err
	}
	return a, nil
}

func (impl *implementation) SearchLostPetPost(ctx context.Context, input *userin.SearchInputStruct) ([]domain.CreateLostPetPostStruct, error) {
	a, err := impl.openSearch.SearchLostPetPost(*input, ctx)
	if err != nil {
		return a, err
	}
	return a, nil
}
