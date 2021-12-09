package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/AdoptMeWithoutKey/post/app/view"
)

func (ctrl *Controller) DeleteAdoptionPost(c *gin.Context) {
	id := c.Param("id")
	err := ctrl.service.DeleteAdoptionPost(c, id)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}
	view.MakeSuccessResp(c, 200, nil)
}

func (ctrl *Controller) DeleteLostPetPost(c *gin.Context) {
	id := c.Param("id")
	err := ctrl.service.DeleteLostPetPost(c, id)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}
	view.MakeSuccessResp(c, 200, nil)
}
