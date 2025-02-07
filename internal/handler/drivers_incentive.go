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
	GetTotalDriversIncentiveByDriverID(ctx *gin.Context)
	GetDriverIncentivesByDriverID(ctx *gin.Context)
}

type driverIncentiveHandlerImpl struct {
	driversincentiveervice service.DriversIncentiveservice
	bookingservice       service.Bookingservice
	driverservice       service.Driverservice
}

func NewDriverIncentiveHandler(driversincentiveervice service.DriversIncentiveservice, bookingservice service.Bookingservice, driverservice service.Driverservice) DriverIncentiveHandler {
	return &driverIncentiveHandlerImpl{driversincentiveervice: driversincentiveervice, bookingservice: bookingservice, driverservice: driverservice}
}
// GetDriversincentive godoc
// @Summary Retrieve list of driver incentives
// @Description Retrieve a list of all driver incentives.
// @Tags driverIncentives
// @Accept json
// @Produce json
// @Success 200 {array} models.DriverIncentive "List of driver incentives"
// @Success 404 {object} pkg.ErrorResponse "No driver incentive found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /driver-incentives [get]
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
// @Summary Retrieve a driver incentive by ID
// @Description Retrieve a specific driver incentive by its ID.
// @Tags driverIncentives
// @Accept json
// @Produce json
// @Param id path int true "Driver Incentive ID"
// @Success 200 {object} models.DriverIncentive "Driver incentive details"
// @Failure 400 {object} pkg.ErrorResponse "Invalid driver incentive ID"
// @Failure 404 {object} pkg.ErrorResponse "Driver incentive not found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /driver-incentives/{id} [get]
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
// @Summary Delete a driver incentive by ID
// @Description Remove a driver incentive using its ID.
// @Tags driverIncentives
// @Accept json
// @Produce json
// @Param id path int true "Driver Incentive ID"
// @Success 200 {object} map[string]any "Successful deletion message"
// @Failure 400 {object} pkg.ErrorResponse "Invalid required param"
// @Failure 404 {object} pkg.ErrorResponse "Driver incentive not found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /driver-incentives/{id} [delete]
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
// @Summary Create a new driver incentive
// @Description Add a new driver incentive record.
// @Tags driverIncentives
// @Accept json
// @Produce json
// @Param driverIncentive body models.InputDriverIncentive true "Driver Incentive Data"
// @Success 201 {object} models.DriverIncentive "Created driver incentive"
// @Failure 400 {object} pkg.ErrorResponse "Bad request"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /driver-incentives [post]
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
// @Summary Update a driver incentive
// @Description Modify details of an existing driver incentive.
// @Tags driverIncentives
// @Accept json
// @Produce json
// @Param id path int true "Driver Incentive ID"
// @Param driverIncentive body models.InputDriverIncentive true "Updated driver incentive data"
// @Success 200 {object} models.DriverIncentive "Updated driver incentive"
// @Failure 400 {object} pkg.ErrorResponse "Invalid required param or request body"
// @Failure 404 {object} pkg.ErrorResponse "Driver incentive not found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /driver-incentives/{id} [put]
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

// @Summary Get Driver Incentives List
// @Description Retrieve a list of incentives for a given driver ID
// @Tags Drivers Incentive
// @Produce json
// @Param id path int true "Driver ID"
// @Success 200 {array} models.DriverIncentive "List of Incentives"
// @Success 404 {object} map[string]string "No driverIncentive found"
// @Failure 400 {object} pkg.ErrorResponse "Invalid Driver ID"
// @Failure 500 {object} pkg.ErrorResponse "Driver not found / Internal Server Error"
// @Router /driver-incentives/driver/{id} [get]
func (p *driverIncentiveHandlerImpl) GetDriverIncentivesByDriverID(ctx *gin.Context) {
	
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}

	_, err = p.driverservice.GetDriversByID(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: "Driver not found"})
		return
	}

	driverIncentive, err := p.driversincentiveervice.GetDriverIncentivesByDriverID(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	if len(driverIncentive) == 0 {
        ctx.JSON(http.StatusOK, gin.H{"message": "No driverIncentive found"})
        return
    }

	ctx.JSON(http.StatusOK, driverIncentive)
}

// @Summary Get Total Driver Incentives
// @Description Retrieve the total incentive amount for a given driver ID
// @Tags Drivers Incentive
// @Produce json
// @Param id path int true "Driver ID"
// @Success 200 {object} map[string]interface{} "Total Incentive for Driver"
// @Failure 400 {object} pkg.ErrorResponse "Invalid Driver ID"
// @Failure 500 {object} pkg.ErrorResponse "Driver not found / Internal Server Error"
// @Router /driver-incentives/driver/{id}/total [get]
func (p *driverIncentiveHandlerImpl) GetTotalDriversIncentiveByDriverID(ctx *gin.Context) {
	
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}

	total, err := p.driversincentiveervice.GetTotalDriversIncentiveByDriverID(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	driver, err := p.driverservice.GetDriversByID(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: "Driver not found"})
		return
	}
	response := map[string]interface{}{
		"driver_id": id,
		"driver_name": driver.Name,
		"total_incentive": total,
		"currency": "IDR", // Tambahan opsional jika perlu
	}

	ctx.JSON(http.StatusOK, response)
}