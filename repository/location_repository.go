package repository

import (
	"context"

	"github.com/AsrofunNiam/tj-fleet-monitor-test/model/domain"
	"gorm.io/gorm"
)

type LocationRepository interface {
	SaveLocation(ctx context.Context, tx *gorm.DB, location domain.VehicleLocation) error
}
