package repository

import (
	"context"

	"car-rental/internal/infrastructure"
	"car-rental/internal/models"

	"gorm.io/gorm"
)

type MembershipQuery interface {
	GetMembership(ctx context.Context) ([]models.Membership, error)
	GetMembershipByID(ctx context.Context, id uint64) (models.Membership, error)
	EditMembership(ctx context.Context, id uint64, membership models.Membership) (models.Membership, error)
	DeleteMembershipByID(ctx context.Context, id uint64) error
	CreateMembership(ctx context.Context, membership models.Membership) (models.Membership, error)
}

type MembershipCommand interface {
	CreateMembership(ctx context.Context, membership models.Membership) (models.Membership, error)
}

type membershipQueryImpl struct {
	db infrastructure.GormPostgres
}

func NewMembershipQuery(db infrastructure.GormPostgres) MembershipQuery {
	return &membershipQueryImpl{db: db}
}

func (u *membershipQueryImpl) GetMembership(ctx context.Context) ([]models.Membership, error) {
	db := u.db.GetConnection()
	membership := []models.Membership{}
	if err := db.
		WithContext(ctx).
		Table("memberships").
		Find(&membership).Error; err != nil {
		return nil, err
	}
	return membership, nil
}

func (u *membershipQueryImpl) GetMembershipByID(ctx context.Context, id uint64) (models.Membership, error) {
	db := u.db.GetConnection()
	membership := models.Membership{}
	if err := db.
		WithContext(ctx).
		Table("memberships").
		Where("id = ?", id).
		Find(&membership).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Membership{}, nil
		}

		return models.Membership{}, err
	}
	return membership, nil
}

func (u *membershipQueryImpl) DeleteMembershipByID(ctx context.Context, id uint64) error {
	db := u.db.GetConnection()
	if err := db.
		WithContext(ctx).
		Table("memberships").
		Delete(&models.Membership{ID: uint(id)}).
		Error; err != nil {
		return err
	}
	return nil
}

func (u *membershipQueryImpl) CreateMembership(ctx context.Context, membership models.Membership) (models.Membership, error) {
	
	db := u.db.GetConnection()
	if err := db.
		WithContext(ctx).
		Table("memberships").
		Save(&membership).Error; err != nil {
		return models.Membership{}, err
	}
	return membership, nil
}
func (u *membershipQueryImpl) EditMembership(ctx context.Context, id uint64, membership models.Membership) (models.Membership, error) {
	db := u.db.GetConnection()
	updatedMembership := models.Membership{}
	if err := db.
		WithContext(ctx).
		Table("memberships").
		Where("id = ?", id).Updates(&membership).First(&updatedMembership).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return models.Membership{}, nil
			}
		}
	return updatedMembership, nil
}
