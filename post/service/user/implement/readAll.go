package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/AdoptMeWithoutKey/post/domain"
	"github.com/gnnchya/AdoptMeWithoutKey/post/service/user/userin"
)

func (impl *implementation) ReadAllAdoptionPost(ctx context.Context, input userin.ReadAllAdoptionPostInputStruct) ([]domain.CreateAdoptionPostStruct, error) {
	if input.Field == "all" {
		a, err := impl.openSearch.ReadAllAdoptionPost(input.Page, input.Limit, ctx)
		if err != nil {
			fmt.Println("err2:", err)
		}
		return a, err
	} else {
		a, err := impl.openSearch.ReadAllAdoptionPostByType(input, ctx)
		if err != nil {
			fmt.Println("err2:", err)
		}
		return a, err
	}
}

func (impl *implementation) ReadAllLostPetPost(ctx context.Context, input userin.ReadAllLostPetPostInputStruct) ([]domain.CreateLostPetPostStruct, error) {
	if input.Field == "all" {
		a, err := impl.openSearch.ReadAllLostPetPost(input.Page, input.Limit, ctx)
		if err != nil {
			fmt.Println("err2:", err)
		}
		return a, err
	} else {
		a, err := impl.openSearch.ReadAllLostPetPostByType(input, ctx)
		if err != nil {
			fmt.Println("err2:", err)
		}
		return a, err
	}
}
