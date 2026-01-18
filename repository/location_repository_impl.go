package repository

import (
	"context"

	"github.com/AsrofunNiam/tj-fleet-monitor-test/model/domain"
	"gorm.io/gorm"
)

type locationRepositoryImpl struct{}

func NewLocationRepository() LocationRepository {
	return &locationRepositoryImpl{}
}

func (r *locationRepositoryImpl) SaveLocation(ctx context.Context, tx *gorm.DB, loc domain.VehicleLocation) error {
	return tx.WithContext(ctx).Create(&loc).Error
}

func (r *locationRepositoryImpl) FindLatestByVehicleID(ctx context.Context, tx *gorm.DB, vehicleID string) (*domain.VehicleLocation, error) {
	var loc domain.VehicleLocation

	err := tx.WithContext(ctx).
		Where("vehicle_id = ?", vehicleID).
		Order("timestamp DESC").
		Limit(1).
		First(&loc).Error

	if err != nil {
		return nil, err
	}

	return &loc, nil
}

func (r *locationRepositoryImpl) FindHistory(ctx context.Context, tx *gorm.DB, vehicleID string, start int64, end int64) ([]domain.VehicleLocation, error) {
	var locations []domain.VehicleLocation

	err := tx.WithContext(ctx).
		Where(
			"vehicle_id = ? AND timestamp BETWEEN ? AND ?",
			vehicleID,
			start,
			end,
		).
		Order("timestamp ASC").
		Find(&locations).Error

	return locations, err
}
