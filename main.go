package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go_restfull_api/configs"
	"go_restfull_api/database"
	"go_restfull_api/dj"
	"go_restfull_api/docs"
	"go_restfull_api/routes"
	"go_restfull_api/utils"
)

/**
 * @author : Donald Trieu
 * @created : 8/25/21, Wednesday
**/
func main() {
	godotenv.Load()
	environment := utils.GetEnvString("RUN_MODEL", "development")
	configs.Init(environment)

	database.Init()
	_ = database.MigrateMysql()
	dj.InitDjContainer()
	utils.InitLogger()
	// programmatically set swagger info
	address := fmt.Sprintf("%s:%s", configs.GetConfig().GetString("server.addr"), configs.GetConfig().GetString("server.port"))

	docs.SwaggerInfo.Title = "RestAPI"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = address
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()

	if environment == "development" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	v1 := router.Group("api/v1")
	{
		cameraGroup := v1.Group("cameras")
		routes.SetupCameraRoute(cameraGroup)
	}
	_ = router.Run(address)
}
