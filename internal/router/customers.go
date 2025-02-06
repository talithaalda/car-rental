package router

import (
	"car-rental/internal/handler"

	"github.com/gin-gonic/gin"
)

type CustomerRouter interface {
	Mount()
}

type CustomerRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.CustomerHandler
}

func NewCustomerRouter(v *gin.RouterGroup, handler handler.CustomerHandler) CustomerRouter {
	return &CustomerRouterImpl{v: v, handler: handler}
}

func (p *CustomerRouterImpl) Mount() {
	// p.v.GET("", p.handler.GetCustomers)
	p.v.GET("/:id", p.handler.GetCustomerByID)
	p.v.GET("", p.handler.GetCustomers)
	p.v.DELETE("/:id", p.handler.DeleteCustomerByID)
	p.v.PUT("/:id", p.handler.EditCustomer)
	p.v.POST("", p.handler.CreateCustomer)
}
