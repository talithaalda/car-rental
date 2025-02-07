package handler

import (
	"net/http"
	"strconv"
	"time"

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
	customerservice service.CustomerService
	carservice service.Carservice
	driverservice service.Driverservice
	booktypeservice service.BookingTypeservice
}

func NewBookingHandler(
    bookingservice service.Bookingservice, 
    customerservice service.CustomerService, 
    carservice service.Carservice,
	driverservice service.Driverservice,
	booktypeservice service.BookingTypeservice,
) BookingHandler {
    return &bookingHandlerImpl{
        bookingservice: bookingservice,
        customerservice: customerservice,
        carservice: carservice,
		driverservice: driverservice,
		booktypeservice: booktypeservice,
    }
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
	_, err := p.customerservice.GetCustomersByID(ctx, uint64(booking.CustomerID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: "customer not found"})
		return
	}
	
	_, err = p.carservice.GetCarsByID(ctx, uint64(booking.CarID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: "car not found"})
		return
	}
	
	_, err = time.Parse("02/01/2006", booking.StartRent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "start date must be in format dd/mm/yyyy"})
		return
	}
	_, err = time.Parse("02/01/2006", booking.EndRent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "end date must be in format dd/mm/yyyy"})
		return
	}

	if booking.StartRent > booking.EndRent {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "start date must be before end date"})
		return
	}

	if booking.BookTypeID != nil && *booking.BookTypeID == 2 {
		if booking.DriverID == nil {
			ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "driver id should be provided for BookTypeID 2"})
			return
		} else {
			_, err := p.driverservice.GetDriversByID(ctx, uint64(*booking.DriverID))
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: "driver not found"})
				return
			}
		}
	}
	if booking.BookTypeID != nil && *booking.BookTypeID == 1 {
		if booking.DriverID != nil {
			ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "driver id should not be provided for BookTypeID 1"})
			return
		}
	}

	if booking.BookTypeID != nil {
		_, err := p.booktypeservice.GetBookingTypesByID(ctx, uint64(*booking.BookTypeID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: "booking type not found"})
			return
		}
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
	_, err = time.Parse("02/01/2006", inputBooking.StartRent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "start date must be in format dd/mm/yyyy"})
		return
	}
	_, err = time.Parse("02/01/2006", inputBooking.EndRent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "end date must be in format dd/mm/yyyy"})
		return
	}

	if inputBooking.StartRent > inputBooking.EndRent {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "start date must be before end date"})
		return
	}

	if inputBooking.BookTypeID != nil && *inputBooking.BookTypeID == 2 {
		if inputBooking.DriverID == nil {
			ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "driver id should be provided for BookTypeID 2"})
			return
		} else {
			_, err := p.driverservice.GetDriversByID(ctx, uint64(*inputBooking.DriverID))
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: "driver not found"})
				return
			}
		}
	}
	if inputBooking.BookTypeID != nil && *inputBooking.BookTypeID == 1 {
		if inputBooking.DriverID != nil {
			ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "driver id should not be provided for BookTypeID 1"})
			return
		}
	}

	if inputBooking.BookTypeID != nil {
		_, err := p.booktypeservice.GetBookingTypesByID(ctx, uint64(*booking.BookTypeID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: "booking type not found"})
			return
		}
	}

    updatedBooking, err := p.bookingservice.EditBooking(ctx, uint64(id), inputBooking)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, updatedBooking)
}
