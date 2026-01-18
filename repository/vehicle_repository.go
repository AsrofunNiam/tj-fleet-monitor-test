package repository

import (
	"context"

	"github.com/AsrofunNiam/tj-fleet-monitor-test/model/domain"
	"gorm.io/gorm"
)

type VehicleRepository interface {
	SaveLocation(ctx context.Context, tx *gorm.DB, location domain.VehicleLocation) error
	FindHistory(ctx context.Context, tx *gorm.DB, vehicleID string, start int64, end int64) ([]domain.VehicleLocation, error)
	FindLatestByVehicleID(ctx context.Context, tx *gorm.DB, vehicleID string) (*domain.VehicleLocation, error)
}
