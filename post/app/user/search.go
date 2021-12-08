package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/AdoptMeWithoutKey/post/app/view"
	"github.com/gnnchya/AdoptMeWithoutKey/post/service/user/userin"
)

func (ctrl *Controller) SearchAdoptionPost(c *gin.Context) {
	input := &userin.SearchInputStruct{}
	fmt.Println("input in app search:", input)
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c, 400, err)
		fmt.Println("error")
		return
	}
	value, err := ctrl.service.SearchAdoptionPost(c, input)
	fmt.Println("value: ", value)
	if err != nil {
		view.MakeErrResp(c, 422, "error search")
		return
	}
	view.MakeSuccessResp(c, 200, value)
}

func (ctrl *Controller) SearchLostPetPost(c *gin.Context) {
	input := &userin.SearchInputStruct{}
	fmt.Println("input in app search:", input)
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c, 400, err)
		fmt.Println("error")
		return
	}
	value, err := ctrl.service.SearchLostPetPost(c, input)
	fmt.Println("value: ", value)
	if err != nil {
		view.MakeErrResp(c, 422, "error search")
		return
	}
	view.MakeSuccessResp(c, 200, value)
}
