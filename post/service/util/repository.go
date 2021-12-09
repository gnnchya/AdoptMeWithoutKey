package util

import (
	"context"
	"github.com/gnnchya/AdoptMeWithoutKey/post/domain"
	"github.com/gnnchya/AdoptMeWithoutKey/post/service/user/userin"
)

//go:generate mockery --name=Repository

type RepositoryDynamoDB interface {
	ReadAdoptionPostByID(id string) (domain.CreateAdoptionPostStruct, error)
	ReadLostPetPostByID(id string) (domain.CreateLostPetPostStruct, error)
	CreateAdoptionPost(input domain.CreateAdoptionPostStruct) error
	CreateLostPetPost(input domain.CreateLostPetPostStruct) error
	UpdateAdoptionPost(input domain.CreateAdoptionPostStruct) error
	UpdateLostPetPost(input domain.CreateLostPetPostStruct) error
	ReadAllAdoptionPost(input userin.ReadAllAdoptionPostInputStruct) ([]domain.CreateAdoptionPostStruct, error)
	ReadAllLostPetPost(input userin.ReadAllLostPetPostInputStruct) ([]domain.CreateLostPetPostStruct, error)
	ReadAllAdoptionPostByField(input userin.ReadAllAdoptionPostInputStruct) ([]domain.CreateAdoptionPostStruct, error)
	ReadAllLostPetPostByField(input userin.ReadAllLostPetPostInputStruct) ([]domain.CreateLostPetPostStruct, error)
	DeleteAdoptionPost(id string) error
	DeleteLostPetPost(id string) error
	ReadUserByID(uid string) (domain.UserStruct, error)
	Register(input domain.UserStruct) error
}

type RepositoryOpenSearch interface {
	SearchAdoptionPost(input userin.SearchInputStruct, ctx context.Context) ([]domain.CreateAdoptionPostStruct, error)
	SearchLostPetPost(input userin.SearchInputStruct, ctx context.Context) ([]domain.CreateLostPetPostStruct, error)
	CreateAdoptionPost(ctx context.Context, title *domain.CreateAdoptionPostStruct) error
	CreateLostPetPost(ctx context.Context, title *domain.CreateLostPetPostStruct) error
	UpdateAdoptionPost(ctx context.Context, title *domain.CreateAdoptionPostStruct) error
	UpdateLostPetPost(ctx context.Context, title *domain.CreateLostPetPostStruct) error
	ReadAdoptionPost(id string, ctx context.Context) (result domain.CreateAdoptionPostStruct, err error)
	ReadLostPetPost(id string, ctx context.Context) (result domain.CreateLostPetPostStruct, err error)
	ReadAllAdoptionPost(page int, size int, ctx context.Context) ([]domain.CreateAdoptionPostStruct, error)
	ReadAllLostPetPost(page int, size int, ctx context.Context) ([]domain.CreateLostPetPostStruct, error)
	ReadAllAdoptionPostByType(input userin.ReadAllAdoptionPostInputStruct, ctx context.Context) ([]domain.CreateAdoptionPostStruct, error)
	ReadAllLostPetPostByType(input userin.ReadAllLostPetPostInputStruct, ctx context.Context) ([]domain.CreateLostPetPostStruct, error)
	SoftDeleteAdoptionPost(ctx context.Context, id string) error
	SoftDeleteLostPetPost(ctx context.Context, id string) error
	DeleteAdoptionPost(ctx context.Context, id string) error
	DeleteLostPetPost(ctx context.Context, id string) error
}
