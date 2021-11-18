package user

import (
	"github.com/gnnchya/AdoptMe/post/domain"
	"github.com/gnnchya/AdoptMe/post/service/user/userin"
)


type Service interface {
	ReadAdoptionPost(id string) (domain.CreateAdoptionPostStruct, error)
	ReadLostPetPost(id string) (domain.CreateLostPetPostStruct, error)
	CreateAdoptionPost (input userin.CreatePostInputStruct) error
	CreateLostPetPost (input userin.CreatePostInputStruct) error
	UpdateAdoptionPost (input domain.CreateAdoptionPostStruct) error
	UpdateLostPetPost (input domain.CreateLostPetPostStruct) error
	ReadAllAdoptionPost (input userin.ReadAllAdoptionPostInputStruct) ([]domain.CreateAdoptionPostStruct, error)
	ReadAllLostPetPost (input userin.ReadAllLostPetPostInputStruct) ([]domain.CreateLostPetPostStruct, error)
	DeleteAdoptionPost (id string) error
	DeleteLostPetPost (id string) error
}
