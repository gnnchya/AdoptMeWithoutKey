package util

import(
	"github.com/gnnchya/AdoptMe/post/domain"
	"github.com/gnnchya/AdoptMe/post/service/user/userin"
)

//go:generate mockery --name=Repository

type RepositoryDynamoDB interface{
	ReadAdoptionPostByID(id string) (domain.CreateAdoptionPostStruct,error)
	ReadLostPetPostByID(id string) (domain.CreateLostPetPostStruct, error)
	CreateAdoptionPost (input domain.CreateAdoptionPostStruct) error
	CreateLostPetPost (input domain.CreateLostPetPostStruct) error
	UpdateAdoptionPost (input domain.CreateAdoptionPostStruct) error
	UpdateLostPetPost (input domain.CreateLostPetPostStruct) error
	ReadAllAdoptionPost (input userin.ReadAllAdoptionPostInputStruct) ([]domain.CreateAdoptionPostStruct, error)
	ReadAllLostPetPost (input userin.ReadAllLostPetPostInputStruct) ([]domain.CreateLostPetPostStruct, error)
	ReadAllAdoptionPostByField (input userin.ReadAllAdoptionPostInputStruct) ([]domain.CreateAdoptionPostStruct, error)
	ReadAllLostPetPostByField (input userin.ReadAllLostPetPostInputStruct) ([]domain.CreateLostPetPostStruct, error)
	DeleteAdoptionPost (id string) error
	DeleteLostPetPost (id string) error
}
