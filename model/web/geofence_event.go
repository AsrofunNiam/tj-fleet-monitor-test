package web

type GeofenceLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GeofenceEvent struct {
	VehicleID string           `json:"vehicle_id"`
	Event     string           `json:"event"`
	Location  GeofenceLocation `json:"location"`
	Timestamp int64            `json:"timestamp"`
}
