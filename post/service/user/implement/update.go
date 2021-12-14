package implement

import (
	"context"
	"github.com/gnnchya/AdoptMeWithoutKey/post/domain"
	"log"
)

func (impl *implementation) UpdateAdoptionPost(ctx context.Context, input domain.CreateAdoptionPostStruct) error {
	err := impl.dynamoDB.CreateAdoptionPost(input)
	if err != nil {
		log.Println("err2:", err)
	}
	err = impl.openSearch.UpdateAdoptionPost(ctx, &input)
	if err != nil {
		log.Println("err2:", err)
	}
	return err
}

func (impl *implementation) UpdateLostPetPost(ctx context.Context, input domain.CreateLostPetPostStruct) error {
	err := impl.dynamoDB.CreateLostPetPost(input)
	if err != nil {
		log.Println("err2:", err)
	}
	err = impl.openSearch.UpdateLostPetPost(ctx, &input)
	if err != nil {
		log.Println("err2:", err)
	}
	return err
}
