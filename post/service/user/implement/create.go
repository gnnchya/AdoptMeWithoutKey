package implement

import (
	"context"
	"fmt"
	"github.com/gnnchya/AdoptMeWithoutKey/post/service/user/userin"
)

func (impl *implementation) CreateAdoptionPost(ctx context.Context, input userin.CreatePostInputStruct) error {
	data := userin.CreateAdoptionPostInput(input)
	err := impl.dynamoDB.CreateAdoptionPost(data)
	if err != nil {
		fmt.Println("err2:", err)
	}
	err = impl.openSearch.CreateAdoptionPost(ctx, &data)
	if err != nil {
		fmt.Println("err2:", err)
	}
	return err
}

func (impl *implementation) CreateLostPetPost(ctx context.Context, input userin.CreatePostInputStruct) error {
	data := userin.CreateLostPetPostPostInput(input)
	err := impl.dynamoDB.CreateLostPetPost(data)
	if err != nil {
		fmt.Println("err2:", err)
	}
	err = impl.openSearch.CreateLostPetPost(ctx, &data)
	if err != nil {
		fmt.Println("err2:", err)
	}
	return err
}
