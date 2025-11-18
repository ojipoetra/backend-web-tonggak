package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ojipoetra/backend-web-tonggak/services"
)

type KamarControllerImpl struct {
	kamarService services.KamarService
}

// Create implements KamarControllers.
func (controller *KamarControllerImpl) Create(c *gin.Context) {
	// kamarCreateRequest := web.CreateKamarRequest{}
	// if err := c.ShouldBindBodyWithJSON(&kamarCreateRequest)

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
