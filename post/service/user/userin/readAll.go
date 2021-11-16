package userin

type ReadAllAdoptionPostStruct struct{
	Field string `json:"field"`
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type ReadAllLostPetPostStruct struct{
	Field string `json:"field"`
	Limit int `json:"limit"`
	Page int `json:"page"`
}
