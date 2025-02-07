package handler

import (
	"net/http"
	"strconv"

	"car-rental/internal/models"
	"car-rental/internal/service"
	"car-rental/pkg"

	"github.com/gin-gonic/gin"
)

type BookingTypeHandler interface {
	GetBookingTypes(ctx *gin.Context)
	GetBookingTypeByID(ctx *gin.Context)
	DeleteBookingTypeByID(ctx *gin.Context)
	CreateBookingType(ctx *gin.Context)
	EditBookingType(ctx *gin.Context)
}

type bookingTypeHandlerImpl struct {
	bookingTypeservice service.BookingTypeservice
}

func NewBookingTypeHandler(bookingTypeservice service.BookingTypeservice) BookingTypeHandler {
	return &bookingTypeHandlerImpl{bookingTypeservice: bookingTypeservice}
}
// GetBookingTypes godoc
// @Summary Retrieve list of bookingTypes
// @Description Retrieve a list of all bookingTypes.
// @Tags bookingTypes
// @Accept json
// @Produce json
// @Success	200	{object} models.BookingType
// @Success 200 {object} pkg.ErrorResponse "No bookingType found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /bookingTypes [get]
func (p *bookingTypeHandlerImpl) GetBookingTypes(ctx *gin.Context) {
	bookingTypes, err := p.bookingTypeservice.GetBookingTypes(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if len(bookingTypes) == 0 {
        ctx.JSON(http.StatusOK, gin.H{"message": "No bookingType found"})
        return
    }
	ctx.JSON(http.StatusOK, bookingTypes)
}
// GetBookingTypeByID godoc
// @Summary Retrieve bookingType by ID
// @Description Retrieve a bookingType by its ID
// @Tags bookingTypes
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "BookingType ID"
// @Success 200 {object} models.UpdateBookingType "bookingType"
// @Failure 400 {object} pkg.ErrorResponse "Bad request"
// @Failure 404 {object} pkg.ErrorResponse "BookingType not found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /bookingTypes/{id} [get]
func (p *bookingTypeHandlerImpl) GetBookingTypeByID(ctx *gin.Context) {
	// get bookingType ID from path parameter
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid bookingType ID"})
		return
	}

	bookingType, err := p.bookingTypeservice.GetBookingTypesByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if bookingType.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "bookingType not found"})
		return
	}

	ctx.JSON(http.StatusOK, bookingType)
}
// DeleteBookingTypeByID godoc
// @Summary Delete bookingType by ID
// @Description Delete a bookingType by its ID.
// @Tags bookingTypes
// @Accept json
// @Produce json
// @Param id path int true "BookingType ID"
// @Security Bearer
// @Success	200	{object} models.UpdateBookingType
// @Failure 404 {object} pkg.ErrorResponse "BookingType not found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /bookingTypes/{id} [delete]
func (p *bookingTypeHandlerImpl) DeleteBookingTypeByID(ctx *gin.Context) {
	// Get bookingType ID from path parameter
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}

	// Delete bookingType by ID
	bookingType, err := p.bookingTypeservice.DeleteBookingType(ctx, uint64(id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	// Check if the bookingType exists
	if bookingType.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "BookingType not found"})
		return
	}


	ctx.JSON(http.StatusOK, map[string]any{
		"bookingType":    bookingType,
		"message": "Your bookingType has been successfully deleted",
	})
}
// CreateBookingType godoc
// @Summary Create a new bookingType
// @Description Create a new bookingType.
// @Tags bookingTypes
// @Accept json
// @Produce json
// @Param bookingType body models.InputBookingType true "BookingType data"
// @Security Bearer
// @Success	200	{object} models.CreateBookingType
// @Failure 400 {object} pkg.ErrorResponse "Bad request"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /bookingTypes [post]
func (p *bookingTypeHandlerImpl) CreateBookingType(ctx *gin.Context) {
	bookingType := models.InputBookingType{}
	if err := ctx.BindJSON(&bookingType); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	createdBookingType, err := p.bookingTypeservice.CreateBookingType(ctx, bookingType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdBookingType)
}
// EditBookingType godoc
// @Summary Update bookingType information
// @Description Update information of a bookingType.
// @Tags bookingTypes
// @Accept json
// @Produce json
// @Param id path int true "BookingType ID"
// @Param bookingType body models.InputBookingType true "BookingType data"
// @Security Bearer
// @Success	200	{object} models.UpdateBookingType
// @Failure 400 {object} pkg.ErrorResponse "Bad request"
// @Failure 401 {object} pkg.ErrorResponse "Unauthorized"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /bookingTypes/{id} [put]
func (p *bookingTypeHandlerImpl) EditBookingType(ctx *gin.Context) {
	
    id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}
	bookingType, err := p.bookingTypeservice.GetBookingTypesByID(ctx, uint64(id))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
        return
    }
    if bookingType.ID == 0 {
        ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "BookingType not found"})
        return
    }

    
    if err := ctx.ShouldBindJSON(&bookingType); err != nil {
        ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "Invalid request body"})
        return
    }
	inputBookingType := models.InputBookingType{}
	inputBookingType.BookingType = bookingType.BookingType
	inputBookingType.Description = bookingType.Description

    updatedBookingType, err := p.bookingTypeservice.EditBookingType(ctx, uint64(id), inputBookingType)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, updatedBookingType)
}
