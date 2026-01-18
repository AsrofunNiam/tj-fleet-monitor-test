package service

import (
	"context"

	"github.com/AsrofunNiam/tj-fleet-monitor-test/model/web"
)

type VehicleService interface {
	SaveLocation(ctx context.Context, location web.VehicleLocationCreateRequest) error
	FindLatestByVehicleID(ctx context.Context, vehicleID string) (web.VehicleLocationResponse, error)
	GetHistory(ctx context.Context, vehicleID string, start int64, end int64) ([]web.VehicleLocationResponse, error)
}
