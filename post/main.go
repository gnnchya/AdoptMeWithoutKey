package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/AdoptMeWithoutKey/post/config"
	ginLogRus "github.com/toorop/gin-logrus"
	"log"
	"time"
	//"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://18.140.154.104:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		} else {
			c.Next()
		}

		//c.Next()

		//c.JSON(http.StatusOK, gin.H{
		//	"message": "this response has custom headers",
		//})
	}
}
func init() { log.SetFlags(log.Lshortfile | log.LstdFlags) }

func main() {

	// Load config
	appConfig := config.Get()
	logs := setupLog()
	// Gin setup
	router := gin.New()
	//gin.SetMode(gin.ReleaseMode)
	router.Use(ginLogRus.Logger(logs), gin.Recovery())
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
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	// old
	//router.Use(cors.AllowAll())
	// Register route to gin
	_ = newApp(appConfig).RegisterRoute(router)

	// Gin start listen
	_ = router.Run()
}
