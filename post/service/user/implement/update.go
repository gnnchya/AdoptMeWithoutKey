package implement

import (
	"fmt"
	"github.com/gnnchya/AdoptMe/post/domain"
)

func (impl *implementation) UpdateAdoptionPost (input domain.CreateAdoptionPostStruct) error {
	err := impl.dynamoDB.CreateAdoptionPost(input)
	if err != nil{
		fmt.Println("err2:", err)
	}

	return err
}

func (impl *implementation) UpdateLostPetPost (input domain.CreateLostPetPostStruct) error {
	err := impl.dynamoDB.CreateLostPetPost(input)
	if err != nil{
		fmt.Println("err2:", err)
	}

	return err
}
