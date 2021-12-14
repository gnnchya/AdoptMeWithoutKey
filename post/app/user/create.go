package user

import (
	"github.com/gnnchya/AdoptMeWithoutKey/post/app/view"
	"github.com/gnnchya/AdoptMeWithoutKey/post/service/user/userin"
	"log"

	"github.com/gin-gonic/gin"

	goxid "github.com/touchtechnologies-product/xid"
)

func (ctrl *Controller) CreateLostPetPost(c *gin.Context) {
	input := &userin.CreatePostInputStruct{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c, 400, err)
		return
	}
	log.Println("create", input)
	initID := goxid.New()
	input.ID = initID.Gen()
	err := ctrl.service.CreateLostPetPost(c, *input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}
	view.MakeSuccessResp(c, 200, nil)
}

func (ctrl *Controller) CreateAdoptionPost(c *gin.Context) {
	input := &userin.CreatePostInputStruct{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c, 400, err)
		return
	}
	log.Println("create", input)
	initID := goxid.New()
	input.ID = initID.Gen()
	err := ctrl.service.CreateAdoptionPost(c, *input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}
	view.MakeSuccessResp(c, 200, nil)
}
