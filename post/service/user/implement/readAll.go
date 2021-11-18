package implement

import (
	"fmt"
	"github.com/gnnchya/AdoptMe/post/domain"
	"github.com/gnnchya/AdoptMe/post/service/user/userin"
)

func (impl *implementation) ReadAllAdoptionPost (input userin.ReadAllAdoptionPostInputStruct) ([]domain.CreateAdoptionPostStruct, error){
	if input.Field == "all"{
		a,err := impl.dynamoDB.ReadAllAdoptionPost(input)
		if err != nil{
			fmt.Println("err2:", err)
		}
		return a, err
	} else {
		a,err := impl.dynamoDB.ReadAllAdoptionPostByField(input)
		if err != nil{
			fmt.Println("err2:", err)
		}
		return a, err
	}
}

func (impl *implementation) ReadAllLostPetPost (input userin.ReadAllLostPetPostInputStruct) ([]domain.CreateLostPetPostStruct, error){
	if input.Field == "all"{
		a,err := impl.dynamoDB.ReadAllLostPetPost(input)
		if err != nil{
			fmt.Println("err2:", err)
		}
		return a, err
	} else {
		a,err := impl.dynamoDB.ReadAllLostPetPostByField(input)
		if err != nil{
			fmt.Println("err2:", err)
		}
		return a, err
	}
}
