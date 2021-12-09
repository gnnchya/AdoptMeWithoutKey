package implement

import (
	"context"
	"fmt"
	"time"
)

func (impl *implementation) SoftDeleteLostPetPost(ctx context.Context, id string) error {
	input, err := impl.dynamoDB.ReadLostPetPostByID(id)
	input.DeleteAt = time.Now().Unix()
	err = impl.dynamoDB.CreateLostPetPost(input)
	if err != nil {
		fmt.Println("err2:", err)
	}
	err = impl.openSearch.SoftDeleteLostPetPost(ctx, input.ID)
	if err != nil {
		fmt.Println("err2:", err)
	}
	return err
}

func (impl *implementation) SoftDeleteAdoptionPost(ctx context.Context, id string) error {
	input, err := impl.dynamoDB.ReadAdoptionPostByID(id)
	input.DeleteAt = time.Now().Unix()
	err = impl.dynamoDB.CreateAdoptionPost(input)
	if err != nil {
		fmt.Println("err2:", err)
	}
	err = impl.openSearch.SoftDeleteAdoptionPost(ctx, input.ID)
	if err != nil {
		fmt.Println("err2:", err)
	}
	return err
}
