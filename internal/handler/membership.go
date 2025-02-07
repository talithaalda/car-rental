package handler

import (
	"net/http"
	"strconv"

	"car-rental/internal/models"
	"car-rental/internal/service"
	"car-rental/pkg"

	"github.com/gin-gonic/gin"
)

type MembershipHandler interface {
	GetMemberships(ctx *gin.Context)
	GetMembershipByID(ctx *gin.Context)
	DeleteMembershipByID(ctx *gin.Context)
	CreateMembership(ctx *gin.Context)
	EditMembership(ctx *gin.Context)
}

type membershipHandlerImpl struct {
	membershipservice service.Membershipservice
}

func NewMembershipHandler(membershipservice service.Membershipservice) MembershipHandler {
	return &membershipHandlerImpl{membershipservice: membershipservice}
}
// GetMemberships godoc
// @Summary Retrieve list of memberships
// @Description Retrieve a list of all memberships.
// @Tags memberships
// @Accept json
// @Produce json
// @Success 200 {array} models.Membership "List of memberships"
// @Success 404 {object} pkg.ErrorResponse "No membership found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /memberships [get]
func (p *membershipHandlerImpl) GetMemberships(ctx *gin.Context) {
	memberships, err := p.membershipservice.GetMemberships(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if len(memberships) == 0 {
        ctx.JSON(http.StatusOK, gin.H{"message": "No membership found"})
        return
    }
	ctx.JSON(http.StatusOK, memberships)
}
// GetMembershipByID godoc
// @Summary Retrieve membership by ID
// @Description Retrieve a membership by its ID
// @Tags memberships
// @Accept json
// @Produce json
// @Param id path int true "Membership ID"
// @Success 200 {object} models.Membership "Membership data"
// @Failure 400 {object} pkg.ErrorResponse "Invalid membership ID"
// @Failure 404 {object} pkg.ErrorResponse "Membership not found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /memberships/{id} [get]
func (p *membershipHandlerImpl) GetMembershipByID(ctx *gin.Context) {
	// get membership ID from path parameter
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid membership ID"})
		return
	}

	membership, err := p.membershipservice.GetMembershipsByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if membership.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "membership not found"})
		return
	}

	ctx.JSON(http.StatusOK, membership)
}
// DeleteMembershipByID godoc
// @Summary Delete membership by ID
// @Description Delete a membership by its ID.
// @Tags memberships
// @Accept json
// @Produce json
// @Param id path int true "Membership ID"
// @Success 200 {object} map[string]any "Success message and deleted membership data"
// @Failure 400 {object} pkg.ErrorResponse "Invalid membership ID"
// @Failure 404 {object} pkg.ErrorResponse "Membership not found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /memberships/{id} [delete]
func (p *membershipHandlerImpl) DeleteMembershipByID(ctx *gin.Context) {
	// Get membership ID from path parameter
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}

	// Delete membership by ID
	membership, err := p.membershipservice.DeleteMembership(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	// Check if the membership exists
	if membership.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Membership not found"})
		return
	}


	ctx.JSON(http.StatusOK, map[string]any{
		"membership":    membership,
		"message": "Your membership has been successfully deleted",
	})
}
// CreateMembership godoc
// @Summary Create a new membership
// @Description Create a new membership.
// @Tags memberships
// @Accept json
// @Produce json
// @Param membership body models.InputMembership true "Membership data"
// @Success 201 {object} models.Membership "Created membership data"
// @Failure 400 {object} pkg.ErrorResponse "Invalid request body"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /memberships [post]
func (p *membershipHandlerImpl) CreateMembership(ctx *gin.Context) {
	membership := models.InputMembership{}
	if err := ctx.BindJSON(&membership); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	createdMembership, err := p.membershipservice.CreateMembership(ctx, membership)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdMembership)
}
// EditMembership godoc
// @Summary Update membership information
// @Description Update information of a membership.
// @Tags memberships
// @Accept json
// @Produce json
// @Param id path int true "Membership ID"
// @Param membership body models.InputMembership true "Updated membership data"
// @Success 200 {object} models.Membership "Updated membership data"
// @Failure 400 {object} pkg.ErrorResponse "Invalid membership ID or request body"
// @Failure 404 {object} pkg.ErrorResponse "Membership not found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /memberships/{id} [put]
func (p *membershipHandlerImpl) EditMembership(ctx *gin.Context) {
	
    id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}
	membership, err := p.membershipservice.GetMembershipsByID(ctx, uint64(id))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
        return
    }
    if membership.ID == 0 {
        ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Membership not found"})
        return
    }

    
    if err := ctx.ShouldBindJSON(&membership); err != nil {
        ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "Invalid request body"})
        return
    }
	inputMembership := models.InputMembership{}
	inputMembership.MembershipName = membership.MembershipName
	inputMembership.Discount = membership.Discount
    updatedMembership, err := p.membershipservice.EditMembership(ctx, uint64(id), inputMembership)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, updatedMembership)
}
