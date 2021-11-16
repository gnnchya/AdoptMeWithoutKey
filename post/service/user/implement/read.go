package implement

import (
	"fmt"
	"github.com/gnnchya/AdoptMe/post/domain"
)

func (impl *implementation) ReadAdoptionPost( id string) (a domain.CreateAdoptionPostStruct, err error) {

	a,err = impl.dynamoDB.ReadAdoptionPostByID(id)
	if err != nil{
		fmt.Println("err2:", err)
	}
	fmt.Println("a:", a)
	return a, err
}

func (impl *implementation) ReadLostPetPost( id string) (a domain.CreateLostPetPostStruct, err error) {

	a,err = impl.dynamoDB.ReadLostPetPostByID(id)
	if err != nil{
		fmt.Println("err2:", err)
	}
	fmt.Println("a:", a)
	return a, err
}
