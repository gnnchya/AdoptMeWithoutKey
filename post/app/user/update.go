package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/AdoptMe/post/app/view"
	"github.com/gnnchya/AdoptMe/post/domain"
	"time"
)

func (ctrl *Controller) UpdateLostPetPost(c *gin.Context) {
	input := &domain.CreateLostPetPostStruct{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c, 400, err)
		return
	}
	input.UpdateAt = time.Now().Unix()
	fmt.Println("update input:",input)
	err := ctrl.service.UpdateLostPetPost(*input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, nil)
}

func (ctrl *Controller) UpdateAdoptionPost(c *gin.Context) {
	input := &domain.CreateAdoptionPostStruct{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c, 400, err)
		return
	}
	input.UpdateAt = time.Now().Unix()
	fmt.Println("update input:",input)
	err := ctrl.service.UpdateAdoptionPost(*input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, nil)
}
