package router

import (
	"car-rental/internal/handler"

	"github.com/gin-gonic/gin"
)

type BookingTypeRouter interface {
	Mount()
}

type bookingTypeRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.BookingTypeHandler
}

func NewBookingTypeRouter(v *gin.RouterGroup, handler handler.BookingTypeHandler) BookingTypeRouter {
	return &bookingTypeRouterImpl{v: v, handler: handler}
}

func (p *bookingTypeRouterImpl) Mount() {
	p.v.GET("/:id", p.handler.GetBookingTypeByID)
	p.v.GET("", p.handler.GetBookingTypes)
	p.v.DELETE("/:id", p.handler.DeleteBookingTypeByID)
	p.v.PUT("/:id", p.handler.EditBookingType)
	p.v.POST("", p.handler.CreateBookingType)
}
