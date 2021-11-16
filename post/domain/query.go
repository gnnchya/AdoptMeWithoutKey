package domain

type AnimalStruct struct{
	Type string `json:"type"`
	Age int64 `json:"age"`
	Species string `json:"species"`
	Gender string `json:"gender"`
	GeneralInformation string `json:"general_information"`
	Spay bool `json:"spay"`
	Image string `json:"image"`
	MedicalCondition string `json:"medical_condition"`
}

type CreateAdoptionPostStruct struct {
	ID string `json:"id"`
	Animal AnimalStruct `json:"animal"`
	Adopt bool `json:"adopt"`
	UID string `json:"uid"`
	Location string `json:"location"`
	PostAt int64 `json:"post_at"`
	UpdateAt int64 `json:"update_at"`
	DeleteAt int64 `json:"delete_at"`
}

type CreateLostPetPostStruct struct {
	ID string `json:"id"`
	Animal AnimalStruct `json:"animal"`
	Found bool `json:"found"`
	UID string `json:"uid"`
	LostLocation string `json:"lost_location"`
	PostAt int64 `json:"post_at"`
	UpdateAt int64 `json:"update_at"`
	DeleteAt int64 `json:"delete_at"`
}