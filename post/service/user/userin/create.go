package userin

import (
	"github.com/gnnchya/AdoptMe/post/domain"
	"time"
)

type CreatePostInputStruct struct{
	ID string `json:"id"`
	UID string `json:"uid"`
	Location string `json:"location"`
	Animal domain.AnimalStruct`json:"animal"`
}

func CreateAdoptionPostInput (input CreatePostInputStruct) domain.CreateAdoptionPostStruct {
	return domain.CreateAdoptionPostStruct{
		ID:	input.ID,
		Animal: input.Animal,
		Adopt: false,
		UID: input.UID,
		Location: input.Location,
		PostAt: time.Now().Unix(),
		UpdateAt: 0,
		DeleteAt: 0,
	}
}

func CreateLostPetPostPostInput (input CreatePostInputStruct) domain.CreateLostPetPostStruct {
	return domain.CreateLostPetPostStruct{
		ID:	input.ID,
		Animal: input.Animal,
		Found: false,
		UID: input.UID,
		LostLocation: input.Location,
		PostAt: time.Now().Unix(),
		UpdateAt: 0,
		DeleteAt: 0,
	}
}
