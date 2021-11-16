package user

import (
	"github.com/gnnchya/AdoptMe/post/service/user"
)

type Controller struct {
	service user.Service
}

func New(service user.Service) (ctrl *Controller) {
	return &Controller{service}
}
