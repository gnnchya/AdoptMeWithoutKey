package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/AdoptMe/post/app/view"
	"github.com/gnnchya/AdoptMe/post/service/user/userin"
	"strconv"
)

func (ctrl *Controller) ReadAllAdoptionPost(c *gin.Context) {
	query := c.Request.URL.Query()
	limit := 10
	page := 1
	keyword := "all"
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		case "keyword":
			keyword = queryValue

		}
	}

	input := &userin.ReadAllAdoptionPostInputStruct{}
	input.Limit = limit
	input.Field = keyword
	input.Page = page

	a, err := ctrl.service.ReadAllAdoptionPost(*input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, a)
}

func (ctrl *Controller) ReadAllLostPetPost(c *gin.Context) {
	query := c.Request.URL.Query()
	limit := 10
	page := 1
	keyword := "all"
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		case "keyword":
			keyword = queryValue

		}
	}

	input := &userin.ReadAllLostPetPostInputStruct{}
	input.Limit = limit
	input.Field = keyword
	input.Page = page

	a, err := ctrl.service.ReadAllLostPetPost(*input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}

	view.MakeSuccessResp(c, 200, a)
}
