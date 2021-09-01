package routes

import (
	"github.com/gin-gonic/gin"
	"go_restfull_api/controllers"
	"go_restfull_api/services"
	"go_restfull_api/utils"
)

/**
 * @author : Donald Trieu
 * @created : 9/1/21, Wednesday
**/
func SetupCameraRoute(router *gin.RouterGroup) {
	c := controllers.CameraController{
		Service: utils.Container.GetFromBindAble(&services.CameraService{}).(*services.CameraService),
	}
	router.GET("/", c.GetAllCamera)
	router.POST("/", c.InsertCamera)
	router.GET("/:id", c.GetCameraSingle)
	router.PUT("/:id", c.UpdateCamera)
}
