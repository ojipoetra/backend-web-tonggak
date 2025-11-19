package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ojipoetra/backend-web-tonggak/controllers"
)
func NewRouter(kamarController controllers.KamarControllers) *gin.Engine{
	router :=  gin.Default()
	api := router.Group("/api")
	{
		api.GET("/kamar",kamarController.FindAll)
		api.POST("/kamar",kamarController.Create)
	}
	return router
}