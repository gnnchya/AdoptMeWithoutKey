package userin

type ReadAllAdoptionPostInputStruct struct{
	Field string `json:"field"`
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type ReadAllLostPetPostInputStruct struct{
	Field string `json:"field"`
	Limit int `json:"limit"`
	Page int `json:"page"`
}
