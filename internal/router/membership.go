package router

import (
	"car-rental/internal/handler"

	"github.com/gin-gonic/gin"
)

type MembershipRouter interface {
	Mount()
}

type membershipRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.MembershipHandler
}

func NewMembershipRouter(v *gin.RouterGroup, handler handler.MembershipHandler) MembershipRouter {
	return &membershipRouterImpl{v: v, handler: handler}
}

func (p *membershipRouterImpl) Mount() {
	p.v.GET("/:id", p.handler.GetMembershipByID)
	p.v.GET("", p.handler.GetMemberships)
	p.v.DELETE("/:id", p.handler.DeleteMembershipByID)
	p.v.PUT("/:id", p.handler.EditMembership)
	p.v.POST("", p.handler.CreateMembership)
}
