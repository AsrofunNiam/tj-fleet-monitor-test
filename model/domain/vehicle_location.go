package domain

type VehicleLocation struct {
	ID        uint    `gorm:"primaryKey"`
	VehicleID string  `gorm:"index;not null"`
	Latitude  float64 `gorm:"type:decimal(10,7);not null"`
	Longitude float64 `gorm:"type:decimal(10,7);not null"`
	Timestamp int64   `gorm:"index;not null"`
}
