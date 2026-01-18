package service

import (
	"context"

	"github.com/AsrofunNiam/tj-fleet-monitor-test/model/web"
)

type LocationService interface {
	SaveLocation(ctx context.Context, location web.VehicleLocationCreateRequest) error
}
