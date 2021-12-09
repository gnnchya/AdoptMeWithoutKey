package user

import (
	"context"
	"github.com/gnnchya/AdoptMeWithoutKey/post/domain"
	"github.com/gnnchya/AdoptMeWithoutKey/post/service/user/userin"
)

type Service interface {
	ReadAdoptionPost(ctx context.Context, id string) (domain.CreateAdoptionPostStruct, error)
	ReadLostPetPost(ctx context.Context, id string) (domain.CreateLostPetPostStruct, error)
	CreateAdoptionPost(ctx context.Context, input userin.CreatePostInputStruct) error
	CreateLostPetPost(ctx context.Context, input userin.CreatePostInputStruct) error
	UpdateAdoptionPost(ctx context.Context, input domain.CreateAdoptionPostStruct) error
	UpdateLostPetPost(ctx context.Context, input domain.CreateLostPetPostStruct) error
	ReadAllAdoptionPost(ctx context.Context, input userin.ReadAllAdoptionPostInputStruct) ([]domain.CreateAdoptionPostStruct, error)
	ReadAllLostPetPost(ctx context.Context, input userin.ReadAllLostPetPostInputStruct) ([]domain.CreateLostPetPostStruct, error)
	DeleteAdoptionPost(ctx context.Context, id string) error
	DeleteLostPetPost(ctx context.Context, id string) error
	Register(input domain.UserStruct) error
	ReadUserByID(uid string) (a domain.UserStruct, err error)
	SearchAdoptionPost(ctx context.Context, input *userin.SearchInputStruct) ([]domain.CreateAdoptionPostStruct, error)
	SearchLostPetPost(ctx context.Context, input *userin.SearchInputStruct) ([]domain.CreateLostPetPostStruct, error)
	Found(ctx context.Context, id string, uid string) error
	Adopted(ctx context.Context, id string, uid string) error
	SoftDeleteAdoptionPost(ctx context.Context, id string) error
	SoftDeleteLostPetPost(ctx context.Context, id string) error
}
