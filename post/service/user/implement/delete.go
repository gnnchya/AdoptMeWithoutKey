package implement

import "fmt"

func (impl *implementation) DeleteAdoptionPost (id string) error {
	err := impl.dynamoDB.DeleteAdoptionPost(id)
	if err != nil{
		fmt.Println("err2:", err)
	}
	return err
}

func (impl *implementation) DeleteLostPetPost (id string) error {
	err := impl.dynamoDB.DeleteLostPetPost(id)
	if err != nil{
		fmt.Println("err2:", err)
	}
	return err
}