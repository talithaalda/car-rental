package handler

import (
	"net/http"
	"strconv"

	"car-rental/internal/models"
	"car-rental/internal/service"
	"car-rental/pkg"

	"github.com/gin-gonic/gin"
)

type DriverHandler interface {
	GetDrivers(ctx *gin.Context)
	GetDriverByID(ctx *gin.Context)
	DeleteDriverByID(ctx *gin.Context)
	CreateDriver(ctx *gin.Context)
	EditDriver(ctx *gin.Context)
}

type driverHandlerImpl struct {
	driverservice service.Driverservice
}

func NewDriverHandler(driverservice service.Driverservice) DriverHandler {
	return &driverHandlerImpl{driverservice: driverservice}
}
// GetDrivers godoc
// @Summary Retrieve list of drivers
// @Description Retrieve a list of all drivers.
// @Tags drivers
// @Accept json
// @Produce json
// @Success	200	{object} models.Driver
// @Success 200 {object} pkg.ErrorResponse "No driver found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /drivers [get]
func (p *driverHandlerImpl) GetDrivers(ctx *gin.Context) {
	drivers, err := p.driverservice.GetDrivers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if len(drivers) == 0 {
        ctx.JSON(http.StatusOK, gin.H{"message": "No driver found"})
        return
    }
	ctx.JSON(http.StatusOK, drivers)
}
// GetDriverByID godoc
// @Summary Retrieve driver by ID
// @Description Retrieve a driver by its ID
// @Tags drivers
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Driver ID"
// @Success 200 {object} models.UpdateDriver "driver"
// @Failure 400 {object} pkg.ErrorResponse "Bad request"
// @Failure 404 {object} pkg.ErrorResponse "Driver not found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /drivers/{id} [get]
func (p *driverHandlerImpl) GetDriverByID(ctx *gin.Context) {
	// get driver ID from path parameter
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid driver ID"})
		return
	}

	driver, err := p.driverservice.GetDriversByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if driver.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "driver not found"})
		return
	}

	ctx.JSON(http.StatusOK, driver)
}
// DeleteDriverByID godoc
// @Summary Delete driver by ID
// @Description Delete a driver by its ID.
// @Tags drivers
// @Accept json
// @Produce json
// @Param id path int true "Driver ID"
// @Security Bearer
// @Success	200	{object} models.UpdateDriver
// @Failure 404 {object} pkg.ErrorResponse "Driver not found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /drivers/{id} [delete]
func (p *driverHandlerImpl) DeleteDriverByID(ctx *gin.Context) {
	// Get driver ID from path parameter
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}

	driver, err := p.driverservice.DeleteDriver(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	if driver.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Driver not found"})
		return
	}


	ctx.JSON(http.StatusOK, map[string]any{
		"driver":    driver,
		"message": "Your driver has been successfully deleted",
	})
}
// CreateDriver godoc
// @Summary Create a new driver
// @Description Create a new driver.
// @Tags drivers
// @Accept json
// @Produce json
// @Param driver body models.InputDriver true "Driver data"
// @Security Bearer
// @Success	200	{object} models.CreateDriver
// @Failure 400 {object} pkg.ErrorResponse "Bad request"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /drivers [post]
func (p *driverHandlerImpl) CreateDriver(ctx *gin.Context) {
	driver := models.InputDriver{}
	if err := ctx.BindJSON(&driver); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	createdDriver, err := p.driverservice.CreateDriver(ctx, driver)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdDriver)
}
// EditDriver godoc
// @Summary Update driver information
// @Description Update information of a driver.
// @Tags drivers
// @Accept json
// @Produce json
// @Param id path int true "Driver ID"
// @Param driver body models.InputDriver true "Driver data"
// @Security Bearer
// @Success	200	{object} models.UpdateDriver
// @Failure 400 {object} pkg.ErrorResponse "Bad request"
// @Failure 401 {object} pkg.ErrorResponse "Unauthorized"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /drivers/{id} [put]
func (p *driverHandlerImpl) EditDriver(ctx *gin.Context) {
	
    id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}
	driver, err := p.driverservice.GetDriversByID(ctx, uint64(id))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
        return
    }
    if driver.ID == 0 {
        ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Driver not found"})
        return
    }

    
    if err := ctx.ShouldBindJSON(&driver); err != nil {
        ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "Invalid request body"})
        return
    }
	inputDriver := models.InputDriver{}
	inputDriver.Name = driver.Name
	inputDriver.NIK = driver.NIK
	inputDriver.Phone = driver.Phone
	inputDriver.DailyCost = driver.DailyCost

    updatedDriver, err := p.driverservice.EditDriver(ctx, uint64(id), inputDriver)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, updatedDriver)
}
