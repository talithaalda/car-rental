package service

import (
	"car-rental/internal/models"
	"car-rental/internal/repository"
	"context"
	"errors"
	"time"
)

type Membershipservice interface {
	GetMemberships(ctx context.Context) ([]models.Membership, error)
	GetMembershipsByID(ctx context.Context, id uint64) (models.Membership, error)
	CreateMembership(ctx context.Context, membership models.InputMembership) (models.Membership, error)
	EditMembership(ctx context.Context, id uint64, membership models.InputMembership) (models.Membership, error)
	DeleteMembership(ctx context.Context, id uint64) (models.Membership, error)
}
type membershipserviceImpl struct {
	membershipRepo repository.MembershipQuery
}

func NewMembershipservice(membershipRepo repository.MembershipQuery) Membershipservice {
	return &membershipserviceImpl{membershipRepo: membershipRepo}
}


func (s *membershipserviceImpl) GetMemberships(ctx context.Context) ([]models.Membership, error) {
	memberships, err := s.membershipRepo.GetMembership(ctx)
	if err != nil {
		return nil, err
	}
	return memberships, nil
}

func (s *membershipserviceImpl) GetMembershipsByID(ctx context.Context, id uint64) (models.Membership, error) {
	membership, err := s.membershipRepo.GetMembershipByID(ctx, id)
	if err != nil {
		return models.Membership{}, err
	}
	if membership.ID == 0 {
		return models.Membership{}, errors.New("membership not found")
	}
	return membership, nil
}

func (s *membershipserviceImpl) CreateMembership(ctx context.Context, membership models.InputMembership) (models.Membership, error) {
	NewMembership := models.Membership{}
	NewMembership.MembershipName = membership.MembershipName
	NewMembership.Discount = membership.Discount
	NewMembership.CreatedAt = time.Now()

	createdMembership, err := s.membershipRepo.CreateMembership(ctx, NewMembership)
	if err != nil {
		return models.Membership{}, err
	}
	return createdMembership, nil
}

func (s *membershipserviceImpl) EditMembership(ctx context.Context, id uint64, membership models.InputMembership) (models.Membership, error) {
	updatedMembership := models.Membership{}
	updatedMembership.MembershipName = membership.MembershipName
	updatedMembership.Discount = membership.Discount
	updatedMembership.UpdatedAt = time.Now()

	updatedMembership, err := s.membershipRepo.EditMembership(ctx, id, updatedMembership)
	if err != nil {
		return models.Membership{}, err
	}
	return updatedMembership, nil
}

func (s *membershipserviceImpl) DeleteMembership(ctx context.Context, id uint64) (models.Membership, error) {
	membership, err := s.membershipRepo.GetMembershipByID(ctx, id)
	if err != nil {
		return models.Membership{}, err
	}
	if membership.ID == 0 {
		return models.Membership{}, nil
	}
	err = s.membershipRepo.DeleteMembershipByID(ctx, id)
	if err != nil {
		return models.Membership{}, err
	}

	return membership, err
}
