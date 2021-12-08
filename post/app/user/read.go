package user

import (
	"github.com/gnnchya/AdoptMeWithoutKey/post/app/view"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) ReadAdoptionPost(c *gin.Context) {
	id := c.Param("id")
	a, err := ctrl.service.ReadAdoptionPost(c, id)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, a)
}

func (ctrl *Controller) ReadLostPetPost(c *gin.Context) {
	id := c.Param("id")
	a, err := ctrl.service.ReadLostPetPost(c, id)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, a)
}

func (ctrl *Controller) ReadUserByID(c *gin.Context) {
	id := c.Param("id")
	a, err := ctrl.service.ReadUserByID(id)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, a)
}
