package userin

import (
	"github.com/gnnchya/AdoptMeWithoutKey/post/domain"
	"time"
)

type CreatePostInputStruct struct {
	ID       string              `bson:"id" json:"id"`
	UID      string              `bson:"uid" json:"uid"`
	Location string              `bson:"location" json:"location"`
	Animal   domain.AnimalStruct `bson:"animal" json:"animal"`
}

func CreateAdoptionPostInput(input CreatePostInputStruct) domain.CreateAdoptionPostStruct {
	return domain.CreateAdoptionPostStruct{
		ID:       input.ID,
		Animal:   input.Animal,
		Adopt:    false,
		UID:      input.UID,
		Location: input.Location,
		PostAt:   time.Now().Unix(),
		UpdateAt: 0,
		DeleteAt: 0,
	}
}

func CreateLostPetPostPostInput(input CreatePostInputStruct) domain.CreateLostPetPostStruct {
	return domain.CreateLostPetPostStruct{
		ID:           input.ID,
		Animal:       input.Animal,
		Found:        false,
		UID:          input.UID,
		LostLocation: input.Location,
		PostAt:       time.Now().Unix(),
		UpdateAt:     0,
		DeleteAt:     0,
	}
}
