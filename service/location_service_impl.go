package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/AsrofunNiam/tj-fleet-monitor-test/helper"
	"github.com/AsrofunNiam/tj-fleet-monitor-test/model/domain"
	"github.com/AsrofunNiam/tj-fleet-monitor-test/model/web"
	"github.com/AsrofunNiam/tj-fleet-monitor-test/repository"
	"github.com/go-playground/validator/v10"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type LocationServiceImpl struct {
	Repo       repository.LocationRepository
	DB         *gorm.DB
	RabbitConn *amqp.Connection
	Validate   *validator.Validate
}

func NewLocationService(
	repo repository.LocationRepository,
	db *gorm.DB,
	rabbitConn *amqp.Connection,
	validate *validator.Validate,
) LocationService {
	return &LocationServiceImpl{
		Repo:       repo,
		DB:         db,
		RabbitConn: rabbitConn,
		Validate:   validate,
	}
}

func (s *LocationServiceImpl) SaveLocation(ctx context.Context, request web.VehicleLocationCreateRequest) (err error) {
	if err = s.Validate.Struct(request); err != nil {
		return err
	}

	tx := s.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer helper.CommitOrRollback(tx, &err)

	newVehicleLocation := domain.VehicleLocation{
		VehicleID: request.VehicleID,
		Latitude:  request.Latitude,
		Longitude: request.Longitude,
		Timestamp: request.Timestamp,
	}

	err = s.Repo.SaveLocation(ctx, tx, newVehicleLocation)
	if err != nil {
		return err
	}

	const (
		//  Geofence can be list of coordinates
		targetLat  = -6.2088
		targetLong = 106.8456
		maxRadius  = 50.0
	)

	distance := helper.CalculateDistance(request.Latitude, request.Longitude, targetLat, targetLong)
	fmt.Printf("Distance: %f\n", distance)
	if distance <= maxRadius {
		eventPayload := web.GeofenceEvent{
			VehicleID: request.VehicleID,
			Event:     "geofence_entry",
			Location: web.GeofenceLocation{
				Latitude:  request.Latitude,
				Longitude: request.Longitude,
			},
			Timestamp: request.Timestamp,
		}

		// Publish ke RabbitMQ
		err = s.publishToRabbitMQ(eventPayload)
		if err != nil {
			log.Printf("Failed to trigger geofence event: %v", err)
		}
	}
	return nil
}

func (s *LocationServiceImpl) publishToRabbitMQ(event web.GeofenceEvent) error {
	ch, err := s.RabbitConn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"fleet.events", // name
		"direct",       // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		return err
	}

	body, _ := json.Marshal(event)

	return ch.Publish(
		"fleet.events",    // exchange
		"geofence_alerts", // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
