package implement

import (
	"context"
	"github.com/gnnchya/AdoptMeWithoutKey/post/domain"
)

func (impl *implementation) ReadAdoptionPost(ctx context.Context, id string) (a domain.CreateAdoptionPostStruct, err error) {

	a, err = impl.openSearch.ReadAdoptionPost(id, ctx)
	if err != nil {
		return a, err
	}
	return a, err
}

func (impl *implementation) ReadLostPetPost(ctx context.Context, id string) (a domain.CreateLostPetPostStruct, err error) {

	a, err = impl.openSearch.ReadLostPetPost(id, ctx)
	if err != nil {
		return a, err
	}
	return a, err
}

func (impl *implementation) ReadUserByID(uid string) (a domain.UserStruct, err error) {
	a, err = impl.dynamoDB.ReadUserByID(uid)
	if err != nil {
		return a, err
	}
	return a, err
}
