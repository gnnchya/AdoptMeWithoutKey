package implement

import (
	"github.com/gnnchya/AdoptMe/post/domain"
)

func (impl *implementation) Register (input domain.UserStruct) error {
	err := impl.dynamoDB.Register(input)
	if err != nil{
		return err
	}
	return err
}
