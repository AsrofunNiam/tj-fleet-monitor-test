package domain

import (
	"github.com/AsrofunNiam/tj-fleet-monitor-test/model/web"
)

type VehicleLocation struct {
	ID        uint    `gorm:"primaryKey"`
	VehicleID string  `gorm:"index;not null"`
	Latitude  float64 `gorm:"type:decimal(10,7);not null"`
	Longitude float64 `gorm:"type:decimal(10,7);not null"`
	Timestamp int64   `gorm:"index;not null"`
}

func ToVehicleLocationResponse(location VehicleLocation) web.VehicleLocationResponse {
	return web.VehicleLocationResponse{
		VehicleID: location.VehicleID,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		Timestamp: location.Timestamp,
	}
}

func ToVehicleLocationResponses(locations []VehicleLocation) []web.VehicleLocationResponse {
	var responses []web.VehicleLocationResponse
	for _, loc := range locations {
		responses = append(responses, ToVehicleLocationResponse(loc))
	}
	return responses
}
