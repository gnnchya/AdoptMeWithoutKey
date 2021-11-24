package implement

import (
	"github.com/gnnchya/AdoptMe/post/domain"
)

func (impl *implementation) ReadAdoptionPost( id string) (a domain.CreateAdoptionPostStruct, err error) {

	a,err = impl.dynamoDB.ReadAdoptionPostByID(id)
	if err != nil{
		return a, err
	}
	return a, err
}

func (impl *implementation) ReadLostPetPost( id string) (a domain.CreateLostPetPostStruct, err error) {

	a,err = impl.dynamoDB.ReadLostPetPostByID(id)
	if err != nil{
		return a, err
	}
	return a, err
}

func (impl *implementation) ReadUserByID( uid string) (a domain.UserStruct, err error) {
	a,err = impl.dynamoDB.ReadUserByID(uid)
	if err != nil{
		return a, err
	}
	return a, err
}