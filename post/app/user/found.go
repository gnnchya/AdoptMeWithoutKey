package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/AdoptMeWithoutKey/post/app/view"
	"github.com/gnnchya/AdoptMeWithoutKey/post/service/user/userin"
)

func (ctrl *Controller) Found(c *gin.Context) {
	input := &userin.FoundInputStruct{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c, 400, err)
		return
	}
	err := ctrl.service.Found(c, input.ID, input.UID)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}
	view.MakeSuccessResp(c, 200, nil)
}
