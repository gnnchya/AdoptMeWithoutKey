package implement

import (
	"github.com/gnnchya/AdoptMe/post/service/user"
	"github.com/gnnchya/AdoptMe/post/service/util"

)

type implementation struct {
	dynamoDB  util.RepositoryDynamoDB
}

func New(dynamoDB  util.RepositoryDynamoDB) (service user.Service) {
	return &implementation{dynamoDB}
}