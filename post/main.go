package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/AdoptMe/post/config"
	//"github.com/gin-contrib/cors"
	//cors "github.com/rs/cors/wrapper/gin"
	ginLogRus "github.com/toorop/gin-logrus"
	//"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {

	// Load config
	appConfig := config.Get()
	log := setupLog()
	// Gin setup
	router := gin.New()
	router.Use(ginLogRus.Logger(log), gin.Recovery())
	//router.Use(cors.New(cors.Options{
	//	AllowedOrigins: []string{"http://localhost:3000/"},
	//	AllowedMethods: []string{
	//		http.MethodHead,
	//		http.MethodGet,
	//		http.MethodPost,
	//		http.MethodPut,
	//		http.MethodPatch,
	//		http.MethodDelete},
	//	ExposedHeaders:   []string{"X-Content-Length"},
	//	AllowedHeaders:   []string{"*"},
	//	AllowCredentials: true,
	//}))

	router.Use(CORSMiddleware())


	// old
	//router.Use(cors.AllowAll())
	// Register route to gin
	_ = newApp(appConfig).RegisterRoute(router)

	// Gin start listen
	_ = router.Run()
}
