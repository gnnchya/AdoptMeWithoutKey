package app

import (
	"github.com/gin-gonic/gin"

	// "touch/service/user"
	"github.com/gnnchya/AdoptMeWithoutKey/post/app/user"
	userService "github.com/gnnchya/AdoptMeWithoutKey/post/service/user"
)

type App struct {
	user *user.Controller
}

func New(userService userService.Service) *App {
	return &App{
		user: user.New(userService),
	}
}

func (app *App) RegisterRoute(router *gin.Engine) *App {
	//adminMiddleware := app.middle.Authorization(app.middle.Users)
	//adminRoute := router.Group("/pos", adminMiddleware)
	adminRoute := router.Group("/AdoptMe")
	{
		adminRoute.GET("/AdoptionPost/:id", app.user.ReadAdoptionPost)
		adminRoute.GET("/LostPetPost/:id", app.user.ReadLostPetPost)
		adminRoute.GET("/UserInfo/:id", app.user.ReadUserByID)
		adminRoute.POST("/Register", app.user.Register)
		adminRoute.POST("/AdoptionPost", app.user.CreateAdoptionPost)
		adminRoute.POST("/LostPetPost", app.user.CreateLostPetPost)
		adminRoute.GET("/AdoptionPost", app.user.ReadAllAdoptionPost)
		adminRoute.GET("/LostPetPost", app.user.ReadAllLostPetPost)
		adminRoute.DELETE("/AdoptionPost/:id", app.user.DeleteAdoptionPost)
		adminRoute.DELETE("/LostPetPost/:id", app.user.DeleteLostPetPost)
		adminRoute.PUT("/AdoptionPost", app.user.UpdateAdoptionPost)
		adminRoute.PUT("/LostPetPost", app.user.UpdateLostPetPost)

	}

	return app
}
