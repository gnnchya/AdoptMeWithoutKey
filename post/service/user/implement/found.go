package implement

import (
	"context"
	"github.com/gnnchya/AdoptMeWithoutKey/post/service/user/email"
	"log"
)

func (impl *implementation) Found(ctx context.Context, id string, uid string) error {
	adopter, err := impl.ReadUserByID(uid)
	if err != nil {
		log.Fatal(err)
		//return err
	}
	postData, err := impl.ReadLostPetPost(ctx, id)
	if err != nil {
		log.Fatal(err)
		//return err
	}
	owner, err := impl.ReadUserByID(postData.UID)
	if err != nil {
		log.Fatal(err)
		//return err
	}
	err = email.SendFoundEmail(owner.Email, adopter)
	if err != nil {
		log.Fatal(err)
		//return err
	}
	return err
}
