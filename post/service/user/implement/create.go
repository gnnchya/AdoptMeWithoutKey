package implement

import (
	"fmt"
	"github.com/gnnchya/AdoptMe/post/service/user/userin"
)

func (impl *implementation) CreateAdoptionPost (input userin.CreatePostInputStruct) error {
	data := userin.CreateAdoptionPostInput(input)
	err := impl.dynamoDB.CreateAdoptionPost(data)
	if err != nil{
		fmt.Println("err2:", err)
	}

	return err
}

func (impl *implementation) CreateLostPetPost (input userin.CreatePostInputStruct) error {
	data := userin.CreateLostPetPostPostInput(input)
	err := impl.dynamoDB.CreateLostPetPost(data)
	if err != nil{
		fmt.Println("err2:", err)
	}

	return err
}
