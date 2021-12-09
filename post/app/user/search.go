package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/AdoptMeWithoutKey/post/app/view"
	"github.com/gnnchya/AdoptMeWithoutKey/post/service/user/userin"
	"strconv"
)

func (ctrl *Controller) SearchAdoptionPost(c *gin.Context) {
	query := c.Request.URL.Query()
	limit := 10
	page := 1
	keyword := ""
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

	input := &userin.SearchInputStruct{}
	input.Size = limit
	input.Keyword = keyword
	input.Page = page

	value, err := ctrl.service.SearchAdoptionPost(c, input)
	fmt.Println("value: ", value)
	if err != nil {
		view.MakeErrResp(c, 422, "error search")
		return
	}
	view.MakeSuccessResp(c, 200, value)
}

func (ctrl *Controller) SearchLostPetPost(c *gin.Context) {
	query := c.Request.URL.Query()
	limit := 10
	page := 1
	keyword := ""
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

	input := &userin.SearchInputStruct{}
	input.Size = limit
	input.Keyword = keyword
	input.Page = page

	value, err := ctrl.service.SearchLostPetPost(c, input)
	fmt.Println("value: ", value)
	if err != nil {
		view.MakeErrResp(c, 422, "error search")
		return
	}
	view.MakeSuccessResp(c, 200, value)
}
