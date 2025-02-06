package service

import (
	"car-rental/internal/models"
	"car-rental/internal/repository"
	"context"
	"errors"
	"time"
)

type Bookingservice interface {
	GetBookings(ctx context.Context) ([]models.Booking, error)
	GetBookingsByID(ctx context.Context, id uint64) (models.Booking, error)
	CreateBooking(ctx context.Context, booking models.InputBooking) (models.Booking, error)
	EditBooking(ctx context.Context, id uint64, booking models.InputBooking) (models.Booking, error)
	DeleteBooking(ctx context.Context, id uint64) (models.Booking, error)
}
type bookingserviceImpl struct {
	bookingRepo repository.BookingsQuery
}

func NewBookingservice(bookingRepo repository.BookingsQuery) Bookingservice {
	return &bookingserviceImpl{bookingRepo: bookingRepo}
}


func (s *bookingserviceImpl) GetBookings(ctx context.Context) ([]models.Booking, error) {
	bookings, err := s.bookingRepo.GetBookings(ctx)
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (s *bookingserviceImpl) GetBookingsByID(ctx context.Context, id uint64) (models.Booking, error) {
	booking, err := s.bookingRepo.GetBookingsByID(ctx, id)
	if err != nil {
		return models.Booking{}, err
	}
	if booking.ID == 0 {
		return models.Booking{}, errors.New("booking not found")
	}
	return booking, nil
}

func (s *bookingserviceImpl) CreateBooking(ctx context.Context, booking models.InputBooking) (models.Booking, error) {
	NewBooking := models.Booking{}
	startRent, err := time.Parse("02/01/2006", booking.StartRent)
	if err != nil {
		return models.Booking{}, err
	}
	endRent, err := time.Parse("02/01/2006", booking.EndRent)
	if err != nil {
		return models.Booking{}, err
	}
	NewBooking.CustomerID = booking.CustomerID
	NewBooking.CarID = booking.CarID
	NewBooking.StartRent = startRent
	NewBooking.EndRent = endRent
	NewBooking.TotalCost = booking.TotalCost
	NewBooking.Finished = booking.Finished
	NewBooking.CreatedAt = time.Now()

	// Call repoBookingsitory to create booking
	createdBooking, err := s.bookingRepo.CreateBookings(ctx, NewBooking)
	if err != nil {
		return models.Booking{}, err
	}
	return createdBooking, nil
}

func (s *bookingserviceImpl) EditBooking(ctx context.Context, id uint64, booking models.InputBooking) (models.Booking, error) {
	updatedBooking := models.Booking{}
	startRent, err := time.Parse("02/01/2006", booking.StartRent)
	if err != nil {
		return models.Booking{}, err
	}
	endRent, err := time.Parse("02/01/2006", booking.EndRent)
	if err != nil {
		return models.Booking{}, err
	}

	updatedBooking.CarID = booking.CarID
	updatedBooking.StartRent = startRent
	updatedBooking.EndRent = endRent
	updatedBooking.TotalCost = booking.TotalCost
	updatedBooking.Finished = booking.Finished
	updatedBooking.UpdatedAt = time.Now()

	// Call repoBookingsitory to create booking
	updatedBooking, err = s.bookingRepo.EditBookings(ctx, id, updatedBooking)
	if err != nil {
		return models.Booking{}, err
	}
	return updatedBooking, nil
}

func (s *bookingserviceImpl) DeleteBooking(ctx context.Context, id uint64) (models.Booking, error) {
	booking, err := s.bookingRepo.GetBookingsByID(ctx, id)
	if err != nil {
		return models.Booking{}, err
	}
	// if booking doesn't exist, return
	if booking.ID == 0 {
		return models.Booking{}, nil
	}

	// delete booking by id
	err = s.bookingRepo.DeleteBookingsByID(ctx, id)
	if err != nil {
		return models.Booking{}, err
	}

	return booking, err
}
