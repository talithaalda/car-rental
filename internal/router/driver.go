package router

import (
	"car-rental/internal/handler"

	"github.com/gin-gonic/gin"
)

type DriverRouter interface {
	Mount()
}

type driverRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.DriverHandler
}

func NewDriverRouter(v *gin.RouterGroup, handler handler.DriverHandler) DriverRouter {
	return &driverRouterImpl{v: v, handler: handler}
}

func (p *driverRouterImpl) Mount() {
	p.v.GET("/:id", p.handler.GetDriverByID)
	p.v.GET("", p.handler.GetDrivers)
	p.v.DELETE("/:id", p.handler.DeleteDriverByID)
	p.v.PUT("/:id", p.handler.EditDriver)
	p.v.POST("", p.handler.CreateDriver)
}
