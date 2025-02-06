package handler

import (
	"net/http"
	"strconv"

	"car-rental/internal/models"
	"car-rental/internal/service"
	"car-rental/pkg"

	"github.com/gin-gonic/gin"
)

type BookingHandler interface {
	GetBookings(ctx *gin.Context)
	GetBookingByID(ctx *gin.Context)
	DeleteBookingByID(ctx *gin.Context)
	CreateBooking(ctx *gin.Context)
	EditBooking(ctx *gin.Context)
}

type bookingHandlerImpl struct {
	bookingservice service.Bookingservice
}

func NewBookingHandler(bookingservice service.Bookingservice) BookingHandler {
	return &bookingHandlerImpl{bookingservice: bookingservice}
}
// GetBookings godoc
// @Summary Retrieve list of bookings
// @Description Retrieve a list of all bookings.
// @Tags bookings
// @Accept json
// @Produce json
// @Success	200	{object} models.Booking
// @Success 200 {object} pkg.ErrorResponse "No booking found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /bookings [get]
func (p *bookingHandlerImpl) GetBookings(ctx *gin.Context) {
	bookings, err := p.bookingservice.GetBookings(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if len(bookings) == 0 {
        ctx.JSON(http.StatusOK, gin.H{"message": "No booking found"})
        return
    }
	ctx.JSON(http.StatusOK, bookings)
}
// GetBookingByID godoc
// @Summary Retrieve booking by ID
// @Description Retrieve a booking by its ID
// @Tags bookings
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Booking ID"
// @Success 200 {object} models.UpdateBooking "booking"
// @Failure 400 {object} pkg.ErrorResponse "Bad request"
// @Failure 404 {object} pkg.ErrorResponse "Booking not found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /bookings/{id} [get]
func (p *bookingHandlerImpl) GetBookingByID(ctx *gin.Context) {
	// get booking ID from path parameter
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid booking ID"})
		return
	}

	booking, err := p.bookingservice.GetBookingsByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if booking.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "booking not found"})
		return
	}

	ctx.JSON(http.StatusOK, booking)
}
// DeleteBookingByID godoc
// @Summary Delete booking by ID
// @Description Delete a booking by its ID.
// @Tags bookings
// @Accept json
// @Produce json
// @Param id path int true "Booking ID"
// @Security Bearer
// @Success	200	{object} models.UpdateBooking
// @Failure 404 {object} pkg.ErrorResponse "Booking not found"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /bookings/{id} [delete]
func (p *bookingHandlerImpl) DeleteBookingByID(ctx *gin.Context) {
	// Get booking ID from path parameter
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}

	booking, err := p.bookingservice.DeleteBooking(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	if booking.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Booking not found"})
		return
	}


	ctx.JSON(http.StatusOK, map[string]any{
		"booking":    booking,
		"message": "Your booking has been successfully deleted",
	})
}
// CreateBooking godoc
// @Summary Create a new booking
// @Description Create a new booking.
// @Tags bookings
// @Accept json
// @Produce json
// @Param booking body models.InputBooking true "Booking data"
// @Security Bearer
// @Success	200	{object} models.CreateBooking
// @Failure 400 {object} pkg.ErrorResponse "Bad request"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /bookings [post]
func (p *bookingHandlerImpl) CreateBooking(ctx *gin.Context) {
	booking := models.InputBooking{}
	if err := ctx.BindJSON(&booking); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	createdBooking, err := p.bookingservice.CreateBooking(ctx, booking)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdBooking)
}
// EditBooking godoc
// @Summary Update booking information
// @Description Update information of a booking.
// @Tags bookings
// @Accept json
// @Produce json
// @Param id path int true "Booking ID"
// @Param booking body models.InputBooking true "Booking data"
// @Security Bearer
// @Success	200	{object} models.UpdateBooking
// @Failure 400 {object} pkg.ErrorResponse "Bad request"
// @Failure 401 {object} pkg.ErrorResponse "Unauthorized"
// @Failure 500 {object} pkg.ErrorResponse "Internal server error"
// @Router /bookings/{id} [put]
func (p *bookingHandlerImpl) EditBooking(ctx *gin.Context) {
	
    id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}
	booking, err := p.bookingservice.GetBookingsByID(ctx, uint64(id))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
        return
    }

    if booking.ID == 0 {
        ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Booking not found"})
        return
    }
	
	startDate := booking.StartRent.Format("02/01/2006")
	endDate := booking.EndRent.Format("02/01/2006")
	inputBooking := models.InputBooking{}
	inputBooking.CustomerID = booking.CustomerID
	inputBooking.CarID = booking.CarID
	inputBooking.StartRent = startDate
	inputBooking.EndRent = endDate
	inputBooking.Finished = booking.Finished
	if err := ctx.ShouldBindJSON(&inputBooking); err != nil {
        ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "Invalid request body"})
        return
    }

    updatedBooking, err := p.bookingservice.EditBooking(ctx, uint64(id), inputBooking)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, updatedBooking)
}
