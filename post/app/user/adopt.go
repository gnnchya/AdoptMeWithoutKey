package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/AdoptMeWithoutKey/post/app/view"
	"github.com/gnnchya/AdoptMeWithoutKey/post/service/user/userin"
)

func (ctrl *Controller) Adopt(c *gin.Context) {
	input := &userin.AdoptInputStruct{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c, 400, err)
		return
	}
	err := ctrl.service.Adopted(c, input.ID, input.UID)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}
	view.MakeSuccessResp(c, 200, nil)
}
