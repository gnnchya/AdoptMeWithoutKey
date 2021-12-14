package implement

import (
	"context"
	"log"
)

func (impl *implementation) DeleteAdoptionPost(ctx context.Context, id string) error {
	err := impl.dynamoDB.DeleteAdoptionPost(id)
	if err != nil {
		log.Println("err2:", err)
	}
	err = impl.openSearch.DeleteAdoptionPost(ctx, id)
	if err != nil {
		log.Println("err2:", err)
	}
	return err
}

func (impl *implementation) DeleteLostPetPost(ctx context.Context, id string) error {
	err := impl.dynamoDB.DeleteLostPetPost(id)
	if err != nil {
		log.Println("err2:", err)
	}
	err = impl.openSearch.DeleteLostPetPost(ctx, id)
	if err != nil {
		log.Println("err2:", err)
	}
	return err
}
