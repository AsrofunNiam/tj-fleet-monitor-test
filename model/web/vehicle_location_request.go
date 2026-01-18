package web

type VehicleLocationCreateRequest struct {
	VehicleID string  `json:"vehicle_id" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
	Timestamp int64   `json:"timestamp" validate:"required"`
}
