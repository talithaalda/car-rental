package handler

import (
	"net/http"
	"strconv"

	"car-rental/internal/models"
	"car-rental/internal/service"
	"car-rental/pkg"

	"github.com/gin-gonic/gin"
)

type DriverIncentiveHandler interface {
	GetDriversincentive(ctx *gin.Context)
	GetDriverIncentiveByID(ctx *gin.Context)
	DeleteDriverIncentiveByID(ctx *gin.Context)
	CreateDriverIncentive(ctx *gin.Context)
	EditDriverIncentive(ctx *gin.Context)
}

type driverIncentiveHandlerImpl struct {
	driversincentiveervice service.DriversIncentiveservice
	bookingservice       service.Bookingservice
}

func NewDriverIncentiveHandler(driversincentiveervice service.DriversIncentiveservice, bookingservice service.Bookingservice) DriverIncentiveHandler {
	return &driverIncentiveHandlerImpl{driversincentiveervice: driversincentiveervice, bookingservice: bookingservice}
}
// GetDriversincentive godoc
// @Summary Retrieve list of driversincentive
// @Description Retrieve a list of all driversincentive.
// @Tags driversincentive
// @Accept json
// @Produce json
// @Success	200	{object} models.DriverIncentive
// @Success 200 {object} pkg.ErrorResponse "No driverIncentive found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /driversincentive [get]
func (p *driverIncentiveHandlerImpl) GetDriversincentive(ctx *gin.Context) {
	driversincentive, err := p.driversincentiveervice.GetDriversIncentive(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if len(driversincentive) == 0 {
        ctx.JSON(http.StatusOK, gin.H{"message": "No driverIncentive found"})
        return
    }
	ctx.JSON(http.StatusOK, driversincentive)
}
// GetDriverIncentiveByID godoc
// @Summary Retrieve driverIncentive by ID
// @Description Retrieve a driverIncentive by its ID
// @Tags driversincentive
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "DriverIncentive ID"
// @Success 200 {object} models.UpdateDriverIncentive "driverIncentive"
// @Failure 400 {object} pkg.ErrorResponse "Bad request"
// @Failure 404 {object} pkg.ErrorResponse "DriverIncentive not found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /driversincentive/{id} [get]
func (p *driverIncentiveHandlerImpl) GetDriverIncentiveByID(ctx *gin.Context) {
	// get driverIncentive ID from path parameter
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid driverIncentive ID"})
		return
	}

	driverIncentive, err := p.driversincentiveervice.GetDriversIncentiveByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if driverIncentive.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "driverIncentive not found"})
		return
	}

	ctx.JSON(http.StatusOK, driverIncentive)
}
// DeleteDriverIncentiveByID godoc
// @Summary Delete driverIncentive by ID
// @Description Delete a driverIncentive by its ID.
// @Tags driversincentive
// @Accept json
// @Produce json
// @Param id path int true "DriverIncentive ID"
// @Security Bearer
// @Success	200	{object} models.UpdateDriverIncentive
// @Failure 404 {object} pkg.ErrorResponse "DriverIncentive not found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /driversincentive/{id} [delete]
func (p *driverIncentiveHandlerImpl) DeleteDriverIncentiveByID(ctx *gin.Context) {
	// Get driverIncentive ID from path parameter
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}

	// Delete driverIncentive by ID
	driverIncentive, err := p.driversincentiveervice.DeleteDriverIncentive(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	// Check if the driverIncentive exists
	if driverIncentive.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "DriverIncentive not found"})
		return
	}


	ctx.JSON(http.StatusOK, map[string]any{
		"driverIncentive":    driverIncentive,
		"message": "Your driverIncentive has been successfully deleted",
	})
}
// CreateDriverIncentive godoc
// @Summary Create a new driverIncentive
// @Description Create a new driverIncentive.
// @Tags driversincentive
// @Accept json
// @Produce json
// @Param driverIncentive body models.InputDriverIncentive true "DriverIncentive data"
// @Security Bearer
// @Success	200	{object} models.CreateDriverIncentive
// @Failure 400 {object} pkg.ErrorResponse "Bad request"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /driversincentive [post]
func (p *driverIncentiveHandlerImpl) CreateDriverIncentive(ctx *gin.Context) {
	driverIncentive := models.InputDriverIncentive{}
	if err := ctx.BindJSON(&driverIncentive); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	_, err := p.bookingservice.GetBookingsByID(ctx, uint64(driverIncentive.BookingID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "Booking not found"})
		return
	}
	createdDriverIncentive, err := p.driversincentiveervice.CreateDriverIncentive(ctx, driverIncentive)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdDriverIncentive)
}
// EditDriverIncentive godoc
// @Summary Update driverIncentive information
// @Description Update information of a driverIncentive.
// @Tags driversincentive
// @Accept json
// @Produce json
// @Param id path int true "DriverIncentive ID"
// @Param driverIncentive body models.InputDriverIncentive true "DriverIncentive data"
// @Security Bearer
// @Success	200	{object} models.UpdateDriverIncentive
// @Failure 400 {object} pkg.ErrorResponse "Bad request"
// @Failure 401 {object} pkg.ErrorResponse "Unauthorized"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /driversincentive/{id} [put]
func (p *driverIncentiveHandlerImpl) EditDriverIncentive(ctx *gin.Context) {
	
    id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}
	driverIncentive, err := p.driversincentiveervice.GetDriversIncentiveByID(ctx, uint64(id))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
        return
    }
	
    if driverIncentive.ID == 0 {
        ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "DriverIncentive not found"})
        return
    }

    
    if err := ctx.ShouldBindJSON(&driverIncentive); err != nil {
        ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "Invalid request body"})
        return
    }
	_, err = p.bookingservice.GetBookingsByID(ctx, uint64(*driverIncentive.BookingID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "Booking not found"})
		return
	}
	inputDriverIncentive := models.InputDriverIncentive{}
	inputDriverIncentive.BookingID = *driverIncentive.BookingID
	inputDriverIncentive.Incentive = driverIncentive.Incentive
    updatedDriverIncentive, err := p.driversincentiveervice.EditDriverIncentive(ctx, uint64(id), inputDriverIncentive)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, updatedDriverIncentive)
}
