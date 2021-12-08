package domain

type AnimalStruct struct {
	Type               string `json:"type" bson:"type"`
	Age                int64  `json:"age" bson:"age"`
	Species            string `json:"species" bson:"species"`
	Gender             string `json:"gender" bson:"gender"`
	GeneralInformation string `json:"general_information" bson:"general_information"`
	Spay               bool   `json:"spay" bson:"spay"`
	Image              string `json:"image" bson:"image"`
	MedicalCondition   string `json:"medical_condition" bson:"medical_condition"`
}

type CreateAdoptionPostStruct struct {
	ID       string       `json:"id" bson:"id"`
	Animal   AnimalStruct `json:"animal" bson:"animal"`
	Adopt    bool         `json:"adopt" bson:"adopt"`
	UID      string       `json:"uid" json:"uid"`
	Location string       `json:"location" bson:"location"`
	PostAt   int64        `json:"post_at" bson:"post_at"`
	UpdateAt int64        `json:"update_at" bson:"update_at"`
	DeleteAt int64        `json:"delete_at" bson:"delete_at"`
}

type CreateLostPetPostStruct struct {
	ID           string       `json:"id" bson:"id"`
	Animal       AnimalStruct `json:"animal" bson:"animal"`
	Found        bool         `json:"found" bson:"found"`
	UID          string       `json:"uid" bson:"UID"`
	LostLocation string       `json:"lost_location" bson:"lost_location"`
	PostAt       int64        `json:"post_at" bson:"post_at"`
	UpdateAt     int64        `json:"update_at" bson:"update_at"`
	DeleteAt     int64        `json:"delete_at" bson:"delete_at"`
}

type UserStruct struct {
	UID       string `json:"uid"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Birthdate int64  `json:"birthdate"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
}

//type AnimalStruct struct {
//	Type               string `json:"type"`
//	Age                int64  `json:"age"`
//	Species            string `json:"species"`
//	Gender             string `json:"gender"`
//	GeneralInformation string `json:"general_information"`
//	Spay               bool   `json:"spay"`
//	Image              string `json:"image"`
//	MedicalCondition   string `json:"medical_condition"`
//}
//
//type CreateAdoptionPostStruct struct {
//	ID       string       `json:"id"`
//	Animal   AnimalStruct `json:"animal"`
//	Adopt    bool         `json:"adopt"`
//	UID      string       `json:"uid"`
//	Location string       `json:"location"`
//	PostAt   int64        `json:"post_at"`
//	UpdateAt int64        `json:"update_at"`
//	DeleteAt int64        `json:"delete_at"`
//}
//
//type CreateLostPetPostStruct struct {
//	ID           string       `json:"id"`
//	Animal       AnimalStruct `json:"animal"`
//	Found        bool         `json:"found"`
//	UID          string       `json:"uid"`
//	LostLocation string       `json:"lost_location"`
//	PostAt       int64        `json:"post_at"`
//	UpdateAt     int64        `json:"update_at"`
//	DeleteAt     int64        `json:"delete_at"`
//}
//
//type UserStruct struct {
//	UID       string `json:"uid"`
//	Username  string `json:"username"`
//	Name      string `json:"name"`
//	Address   string `json:"address"`
//	Birthdate int64  `json:"birthdate"`
//	Email     string `json:"email"`
//	Gender    string `json:"gender"`
//}
