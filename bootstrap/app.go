package bootstrap

import (
	"blog-api/app/repositories/services"
	"blog-api/config"
	"blog-api/routes"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitApp() *http.Server {
	app := gin.New()
	//gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode)

	//route.Use(cors.Default())
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	//config.GetConfig()
	route := routes.InitRouter(app)
	c := config.GetConfig("./app.yaml")
	services.InitMySQL("./app.yaml")

	server := &http.Server{
		Addr:           c.Server.Host + ":" + c.Server.Port,
		Handler:        route,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return server
}
