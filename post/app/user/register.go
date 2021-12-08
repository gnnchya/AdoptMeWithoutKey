package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/AdoptMeWithoutKey/post/app/view"
	"github.com/gnnchya/AdoptMeWithoutKey/post/domain"
)

func (ctrl *Controller) Register(c *gin.Context) {
	input := &domain.UserStruct{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c, 400, err)
		return
	}
	err := ctrl.service.Register(*input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}
	view.MakeSuccessResp(c, 200, nil)
}
