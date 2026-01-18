package service

import (
	"context"
	"errors"

	"github.com/AsrofunNiam/tj-fleet-monitor-test/helper"
	"github.com/AsrofunNiam/tj-fleet-monitor-test/model/domain"
	"github.com/AsrofunNiam/tj-fleet-monitor-test/repository"
	"gorm.io/gorm"
)

type LocationServiceImpl struct {
	Repo repository.LocationRepository
	DB   *gorm.DB
}

func NewLocationService(
	repo repository.LocationRepository,
	db *gorm.DB,
) LocationService {
	return &LocationServiceImpl{
		Repo: repo,
		DB:   db,
	}
}

func (s *LocationServiceImpl) SaveLocation(ctx context.Context, loc domain.VehicleLocation) error {
	tx := s.DB.WithContext(ctx).Begin()
	err := tx.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	if loc.VehicleID == "" {
		return errors.New("vehicle_id is required")
	}
	if loc.Timestamp == 0 {
		return errors.New("timestamp is required")
	}

	return s.Repo.SaveLocation(ctx, tx, loc)
}
