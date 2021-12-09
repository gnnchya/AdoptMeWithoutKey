package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/AdoptMeWithoutKey/post/app/view"
)

func (ctrl *Controller) SoftDeleteLostPetPost(c *gin.Context) {
	id := c.Param("id")
	err := ctrl.service.SoftDeleteLostPetPost(c, id)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, nil)
}

func (ctrl *Controller) SoftDeleteAdoptionPost(c *gin.Context) {
	id := c.Param("id")
	err := ctrl.service.SoftDeleteAdoptionPost(c, id)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, nil)
}
