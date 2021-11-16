package implement

import (
	"fmt"
	"github.com/gnnchya/AdoptMe/post/domain"
)

func (impl *implementation) CreateAdoptionPost (input domain.CreateAdoptionPostStruct) error {
	err := impl.dynamoDB.CreateAdoptionPost(input)
	if err != nil{
		fmt.Println("err2:", err)
	}

	return err
}

func (impl *implementation) CreateLostPetPost (input domain.CreateLostPetPostStruct) error {
	err := impl.dynamoDB.CreateLostPetPost(input)
	if err != nil{
		fmt.Println("err2:", err)
	}

	return err
}
