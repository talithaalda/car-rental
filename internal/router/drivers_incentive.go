package router

import (
	"car-rental/internal/handler"

	"github.com/gin-gonic/gin"
)

type DriverIncentiveRouter interface {
	Mount()
}

type driverIncentiveRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.DriverIncentiveHandler
}

func NewDriverIncentiveRouter(v *gin.RouterGroup, handler handler.DriverIncentiveHandler) DriverIncentiveRouter {
	return &driverIncentiveRouterImpl{v: v, handler: handler}
}

func (p *driverIncentiveRouterImpl) Mount() {
	p.v.GET("/:id", p.handler.GetDriverIncentiveByID)
	p.v.GET("", p.handler.GetDriversincentive)
	p.v.DELETE("/:id", p.handler.DeleteDriverIncentiveByID)
	p.v.PUT("/:id", p.handler.EditDriverIncentive)
	p.v.POST("", p.handler.CreateDriverIncentive)
}
