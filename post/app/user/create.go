package user

import (
	"fmt"
	"github.com/gnnchya/AdoptMe/post/app/view"
	"github.com/gnnchya/AdoptMe/post/domain"

	"github.com/gin-gonic/gin"

	goxid "github.com/touchtechnologies-product/xid"
)

func (ctrl *Controller) CreateLostPetPost(c *gin.Context) {
	input := domain.CreateLostPetPostStruct{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c, 400, err)
		return
	}

	initID := goxid.New()
	input.ID = initID.Gen()

	err := ctrl.service.CreateLostPetPost(input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, nil)
}

func (ctrl *Controller) CreateAdoptionPost(c *gin.Context) {
	input := &domain.CreateAdoptionPostStruct{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c, 400, err)
		return
	}
	fmt.Println("create")
	initID := goxid.New()
	input.ID = initID.Gen()

	err := ctrl.service.CreateAdoptionPost(*input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, nil)
}
