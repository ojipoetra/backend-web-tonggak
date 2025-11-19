package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ojipoetra/backend-web-tonggak/models/web"
	"github.com/ojipoetra/backend-web-tonggak/services"
)

type KamarControllerImpl struct {
	kamarService services.KamarService
}

// Create implements KamarControllers.
func (controller *KamarControllerImpl) Create(ctx *gin.Context) {
	kamarCreateRequest := web.CreateKamarRequest{}
	if err := ctx.ShouldBindBodyWithJSON(&kamarCreateRequest); err != nil{
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Code: 400,
			Status: "BAD REQUEST",
		})
		return
	}
	kamarResponse := controller.kamarService.Create(ctx, kamarCreateRequest)
	ctx.JSON(http.StatusOK, web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: kamarResponse,
	})


}

// Delete implements KamarControllers.
func (controller *KamarControllerImpl) Delete(c *gin.Context) {
	panic("unimplemented")
}

// FindAll implements KamarControllers.
func (controller *KamarControllerImpl) FindAll(c *gin.Context) {
	panic("unimplemented")
}

// FindById implements KamarControllers.
func (controller *KamarControllerImpl) FindById(c *gin.Context) {
	panic("unimplemented")
}

// Update implements KamarControllers.
func (controller *KamarControllerImpl) Update(c *gin.Context) {
	panic("unimplemented")
}

func NewKamarController(kamarService services.KamarService) KamarControllers {
	return &KamarControllerImpl{
		kamarService: kamarService,
	}
}
