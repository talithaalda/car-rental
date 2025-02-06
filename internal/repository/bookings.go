package repository

import (
	"context"
	"fmt"

	"car-rental/internal/infrastructure"
	"car-rental/internal/models"

	"gorm.io/gorm"
)

type BookingsQuery interface {
	GetBookings(ctx context.Context) ([]models.Booking, error)
	GetBookingsByID(ctx context.Context, id uint64) (models.Booking, error)
	EditBookings(ctx context.Context, id uint64, bookings models.Booking) (models.Booking, error)
	DeleteBookingsByID(ctx context.Context, id uint64) error
	CreateBookings(ctx context.Context, bookings models.Booking) (models.Booking, error)
}

type BookingsCommand interface {
	CreateBookings(ctx context.Context, bookings models.Booking) (models.Booking, error)
}

type bookingsQueryImpl struct {
	db infrastructure.GormPostgres
}

func NewBookingsQuery(db infrastructure.GormPostgres) BookingsQuery {
	return &bookingsQueryImpl{db: db}
}

func (u *bookingsQueryImpl) GetBookings(ctx context.Context) ([]models.Booking, error) {
	db := u.db.GetConnection()
	bookings := []models.Booking{}

	if err := db.WithContext(ctx).
		Preload("Customer"). 
		Preload("Car").      
		Find(&bookings).Error; err != nil {
		return nil, err
	}

	return bookings, nil
}


func (u *bookingsQueryImpl) GetBookingsByID(ctx context.Context, id uint64) (models.Booking, error) {
	db := u.db.GetConnection()
	bookings := models.Booking{}

	if err := db.WithContext(ctx).
		Preload("Customer"). 
		Preload("Car").     
		First(&bookings, id).Error; err != nil { 
		if err == gorm.ErrRecordNotFound {
			return models.Booking{}, nil
		}
		return models.Booking{}, err
	}

	return bookings, nil
}


func (u *bookingsQueryImpl) DeleteBookingsByID(ctx context.Context, id uint64) error {
	fmt.Println("repo")
	db := u.db.GetConnection()

	booking := models.Booking{}
	if err := db.WithContext(ctx).
		Preload("Customer").
		Preload("Car").
		First(&booking, id).Error; err != nil {
		return err
	}

	// Hapus booking
	if err := db.WithContext(ctx).
		Delete(&booking).Error; err != nil {
		return err
	}

	return nil
}


func (u *bookingsQueryImpl) CreateBookings(ctx context.Context, bookings models.Booking) (models.Booking, error) {
	db := u.db.GetConnection()

	if err := db.WithContext(ctx).Table("bookings").Save(&bookings).Error; err != nil {
		return models.Booking{}, err
	}

	if err := db.WithContext(ctx).
		Preload("Customer").
		Preload("Car").
		First(&bookings, bookings.ID).Error; err != nil {
		return models.Booking{}, err
	}

	return bookings, nil
}
func (u *bookingsQueryImpl) EditBookings(ctx context.Context, id uint64, booking models.Booking) (models.Booking, error) {
	db := u.db.GetConnection()

	if err := db.WithContext(ctx).
		Model(&models.Booking{}).
		Where("id = ?", id).
		Updates(booking).Error; err != nil {
		return models.Booking{}, err
	}

	updatedBooking := models.Booking{}
	if err := db.WithContext(ctx).
		Preload("Customer").
		Preload("Car").
		First(&updatedBooking, id).Error; err != nil {
		return models.Booking{}, err
	}

	return updatedBooking, nil
}

