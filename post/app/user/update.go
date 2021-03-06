package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/AdoptMeWithoutKey/post/app/view"
	"github.com/gnnchya/AdoptMeWithoutKey/post/domain"
	"log"
	"time"
)

func (ctrl *Controller) UpdateLostPetPost(c *gin.Context) {
	input := &domain.CreateLostPetPostStruct{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c, 400, err)
		return
	}
	input.UpdateAt = time.Now().Unix()
	log.Println("update input:", input)
	err := ctrl.service.UpdateLostPetPost(c, *input)
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
	log.Println("update input:", input)
	err := ctrl.service.UpdateAdoptionPost(c, *input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, nil)
}
