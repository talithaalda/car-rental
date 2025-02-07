package service

import (
	"car-rental/internal/models"
	"car-rental/internal/repository"
	"context"
	"errors"
	"fmt"
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
	carRepo     repository.CarsQuery
	customerRepo repository.CustomersQuery
	driverRepo  repository.DriversQuery
	driverIncentiveRepo repository.DriversIncentiveQuery

}

func NewBookingservice(bookingRepo repository.BookingsQuery, 
	carRepo repository.CarsQuery, 
	customerRepo repository.CustomersQuery,
	driverRepo repository.DriversQuery,
	driverIncentiveRepo repository.DriversIncentiveQuery) Bookingservice {
	return &bookingserviceImpl{bookingRepo: bookingRepo, 
		carRepo: carRepo, 
		customerRepo: customerRepo,
		driverRepo: driverRepo,
		driverIncentiveRepo: driverIncentiveRepo,
	}
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
	daysOfRent := int(endRent.Sub(startRent).Hours() / 24) + 1
	if daysOfRent <= 0 {
		return models.Booking{}, fmt.Errorf("EndRent must be after StartRent")
	}

	car, err := s.carRepo.GetCarsByID(ctx, uint64(booking.CarID))
	if err != nil {
		return models.Booking{}, err
	}
	
	totalCost := daysOfRent * car.DailyRent

	NewBooking.CustomerID = booking.CustomerID
	NewBooking.CarID = booking.CarID
	NewBooking.BookTypeID = booking.BookTypeID
	NewBooking.DriverID = booking.DriverID
	NewBooking.StartRent = startRent
	NewBooking.EndRent = endRent
	NewBooking.TotalCost = totalCost
	NewBooking.Finished = booking.Finished
	NewBooking.CreatedAt = time.Now()

	customer, err := s.customerRepo.GetCustomersByID(ctx, uint64(booking.CustomerID))
	if err != nil {
		return models.Booking{}, err
	}
	if customer.MembershipID != nil {
		membershipDiscount := customer.Membership.Discount
		discount := totalCost * membershipDiscount/100
		NewBooking.Discount = discount
	}
	if booking.BookTypeID != nil && *booking.BookTypeID == 2 {
		driver, err := s.driverRepo.GetDriversByID(ctx, uint64(*booking.DriverID))
		if err != nil {
			return models.Booking{}, err
		}
		driverCost := driver.DailyCost
		totalDriverCost := daysOfRent * driverCost
		NewBooking.TotalDriverCost = totalDriverCost
	}

	createdBooking, err := s.bookingRepo.CreateBookings(ctx, NewBooking)
	if err == nil {
		incentiveDriver := (daysOfRent * car.DailyRent)*5/100
		newIncentive := models.DriverIncentive{}
		newIncentive.Incentive = incentiveDriver
		newIncentive.BookingID = &createdBooking.ID

		_, err = s.driverIncentiveRepo.CreateDriversIncentive(ctx, newIncentive)
		if err != nil {
			return models.Booking{}, err
		}
	}

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

	daysOfRent := int(endRent.Sub(startRent).Hours() / 24) + 1
	if daysOfRent <= 0 {
		return models.Booking{}, fmt.Errorf("EndRent must be after StartRent")
	}

	car, err := s.carRepo.GetCarsByID(ctx, uint64(booking.CarID))
	if err != nil {
		return models.Booking{}, err
	}
	
	totalCost := daysOfRent * car.DailyRent
	updatedBooking.TotalCost = totalCost
	updatedBooking.CarID = booking.CarID
	updatedBooking.StartRent = startRent
	updatedBooking.EndRent = endRent
	updatedBooking.Finished = booking.Finished
	updatedBooking.UpdatedAt = time.Now()

	customer, err := s.customerRepo.GetCustomersByID(ctx, uint64(booking.CustomerID))
	if err != nil {
		return models.Booking{}, err
	}
	if customer.MembershipID != nil {
		membershipDiscount := customer.Membership.Discount
		discount := totalCost * membershipDiscount/100
		updatedBooking.Discount = discount
	}
	if booking.BookTypeID != nil && *booking.BookTypeID == 2 {
		driver, err := s.driverRepo.GetDriversByID(ctx, uint64(*booking.DriverID))
		if err != nil {
			return models.Booking{}, err
		}
		driverCost := driver.DailyCost
		totalDriverCost := daysOfRent * driverCost
		updatedBooking.TotalDriverCost = totalDriverCost
	}

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
	if booking.ID == 0 {
		return models.Booking{}, nil
	}

	err = s.bookingRepo.DeleteBookingsByID(ctx, id)
	if err != nil {
		return models.Booking{}, err
	}

	return booking, err
}
