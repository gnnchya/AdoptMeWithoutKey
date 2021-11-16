package user

import (
	"github.com/gnnchya/AdoptMe/post/domain"
	"github.com/gnnchya/AdoptMe/post/service/user/userin"
)


type Service interface {
	ReadAdoptionPost(id string) (domain.CreateAdoptionPostStruct, error)
	ReadLostPetPost(id string) (domain.CreateLostPetPostStruct, error)
	CreateAdoptionPost (input domain.CreateAdoptionPostStruct) error
	CreateLostPetPost (input domain.CreateLostPetPostStruct) error
	ReadAllAdoptionPost (input userin.ReadAllAdoptionPostStruct) ([]domain.CreateAdoptionPostStruct, error)
	ReadAllLostPetPost (input userin.ReadAllLostPetPostStruct) ([]domain.CreateLostPetPostStruct, error)
	DeleteAdoptionPost (id string) error
	DeleteLostPetPost (id string) error
}
