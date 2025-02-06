package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"car-rental/internal/models"
	"car-rental/internal/service"
	"car-rental/pkg"

	"github.com/gin-gonic/gin"
)

type CarHandler interface {
	GetCars(ctx *gin.Context)
	GetCarByID(ctx *gin.Context)
	DeleteCarByID(ctx *gin.Context)
	CreateCar(ctx *gin.Context)
	EditCar(ctx *gin.Context)
}

type carHandlerImpl struct {
	carservice service.Carservice
}

func NewCarHandler(carservice service.Carservice) CarHandler {
	return &carHandlerImpl{carservice: carservice}
}
// GetCars godoc
// @Summary Retrieve list of cars
// @Description Retrieve a list of all cars.
// @Tags cars
// @Accept json
// @Produce json
// @Success	200	{object} models.Car
// @Success 200 {object} pkg.ErrorResponse "No car found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /cars [get]
func (p *carHandlerImpl) GetCars(ctx *gin.Context) {
	cars, err := p.carservice.GetCars(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if len(cars) == 0 {
        ctx.JSON(http.StatusOK, gin.H{"message": "No car found"})
        return
    }
	ctx.JSON(http.StatusOK, cars)
}
// GetCarByID godoc
// @Summary Retrieve car by ID
// @Description Retrieve a car by its ID
// @Tags cars
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Car ID"
// @Success 200 {object} models.UpdateCar "car"
// @Failure 400 {object} pkg.ErrorResponse "Bad request"
// @Failure 404 {object} pkg.ErrorResponse "Car not found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /cars/{id} [get]
func (p *carHandlerImpl) GetCarByID(ctx *gin.Context) {
	// get car ID from path parameter
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid car ID"})
		return
	}

	car, err := p.carservice.GetCarsByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if car.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "car not found"})
		return
	}

	ctx.JSON(http.StatusOK, car)
}
// DeleteCarByID godoc
// @Summary Delete car by ID
// @Description Delete a car by its ID.
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "Car ID"
// @Security Bearer
// @Success	200	{object} models.UpdateCar
// @Failure 404 {object} pkg.ErrorResponse "Car not found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /cars/{id} [delete]
func (p *carHandlerImpl) DeleteCarByID(ctx *gin.Context) {
	// Get car ID from path parameter
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}

	// Delete car by ID
	car, err := p.carservice.DeleteCar(ctx, uint64(id))
	fmt.Println(car)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	// Check if the car exists
	if car.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Car not found"})
		return
	}


	ctx.JSON(http.StatusOK, map[string]any{
		"car":    car,
		"message": "Your car has been successfully deleted",
	})
}
// CreateCar godoc
// @Summary Create a new car
// @Description Create a new car.
// @Tags cars
// @Accept json
// @Produce json
// @Param car body models.InputCar true "Car data"
// @Security Bearer
// @Success	200	{object} models.CreateCar
// @Failure 400 {object} pkg.ErrorResponse "Bad request"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /cars [post]
func (p *carHandlerImpl) CreateCar(ctx *gin.Context) {
	car := models.InputCar{}
	if err := ctx.BindJSON(&car); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	createdCar, err := p.carservice.CreateCar(ctx, car)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdCar)
}
// EditCar godoc
// @Summary Update car information
// @Description Update information of a car.
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "Car ID"
// @Param car body models.InputCar true "Car data"
// @Security Bearer
// @Success	200	{object} models.UpdateCar
// @Failure 400 {object} pkg.ErrorResponse "Bad request"
// @Failure 401 {object} pkg.ErrorResponse "Unauthorized"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /cars/{id} [put]
func (p *carHandlerImpl) EditCar(ctx *gin.Context) {
	
    id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}
	car, err := p.carservice.GetCarsByID(ctx, uint64(id))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
        return
    }
    if car.ID == 0 {
        ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Car not found"})
        return
    }

    
    if err := ctx.ShouldBindJSON(&car); err != nil {
        ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "Invalid request body"})
        return
    }
	inputCar := models.InputCar{}
	inputCar.Name = car.Name
	inputCar.Stock = car.Stock
	inputCar.DailyRent = car.DailyRent
    updatedCar, err := p.carservice.EditCar(ctx, uint64(id), inputCar)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, updatedCar)
}
