package service

import (
	"context"

	"github.com/AsrofunNiam/tj-fleet-monitor-test/model/domain"
)

type LocationService interface {
	SaveLocation(ctx context.Context, location domain.VehicleLocation) error
}
