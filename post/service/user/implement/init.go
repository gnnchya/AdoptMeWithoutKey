package implement

import (
	"github.com/gnnchya/AdoptMeWithoutKey/post/service/user"
	"github.com/gnnchya/AdoptMeWithoutKey/post/service/util"
)

type implementation struct {
	dynamoDB   util.RepositoryDynamoDB
	openSearch util.RepositoryOpenSearch
}

func New(dynamoDB util.RepositoryDynamoDB, openSearch util.RepositoryOpenSearch) (service user.Service) {
	return &implementation{dynamoDB, openSearch}
}
