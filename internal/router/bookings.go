package router

import (
	"car-rental/internal/handler"

	"github.com/gin-gonic/gin"
)

type BookingRouter interface {
	Mount()
}

type BookingRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.BookingHandler
}

func NewBookingRouter(v *gin.RouterGroup, handler handler.BookingHandler) BookingRouter {
	return &BookingRouterImpl{v: v, handler: handler}
}

func (p *BookingRouterImpl) Mount() {
	p.v.GET("/:id", p.handler.GetBookingByID)
	p.v.GET("", p.handler.GetBookings)
	p.v.DELETE("/:id", p.handler.DeleteBookingByID)
	p.v.PUT("/:id", p.handler.EditBooking)
	p.v.POST("", p.handler.CreateBooking)
}
