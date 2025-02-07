package router

import (
	"car-rental/internal/handler"

	"github.com/gin-gonic/gin"
)

type CarRouter interface {
	Mount()
}

type carRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.CarHandler
}

func NewCarRouter(v *gin.RouterGroup, handler handler.CarHandler) CarRouter {
	return &carRouterImpl{v: v, handler: handler}
}

func (p *carRouterImpl) Mount() {
	p.v.GET("/:id", p.handler.GetCarByID)
	p.v.GET("", p.handler.GetCars)
	p.v.DELETE("/:id", p.handler.DeleteCarByID)
	p.v.PUT("/:id", p.handler.EditCar)
	p.v.POST("", p.handler.CreateCar)
}
