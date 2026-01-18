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

type VehicleServiceImpl struct {
	VehicleRepository repository.VehicleRepository
	DB                *gorm.DB
	RabbitConn        *amqp.Connection
	Validate          *validator.Validate
}

func NewVehicleService(
	vehicleRepository repository.VehicleRepository,
	db *gorm.DB,
	rabbitConn *amqp.Connection,
	validate *validator.Validate,
) VehicleService {
	return &VehicleServiceImpl{
		VehicleRepository: vehicleRepository,
		DB:                db,
		RabbitConn:        rabbitConn,
		Validate:          validate,
	}
}

func (s *VehicleServiceImpl) SaveLocation(ctx context.Context, request web.VehicleLocationCreateRequest) (err error) {
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

	err = s.VehicleRepository.SaveLocation(ctx, tx, newVehicleLocation)
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

func (s *VehicleServiceImpl) publishToRabbitMQ(event web.GeofenceEvent) error {
	ch, err := s.RabbitConn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Declare Exchange
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

	// Declare Queue
	q, err := ch.QueueDeclare(
		"geofence_alerts", // name queue
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		return err
	}

	// Binding Exchange
	err = ch.QueueBind(
		q.Name,            // geofence_alerts
		"geofence_alerts", // routing key
		"fleet.events",    // exchange name
		false,
		nil,
	)
	if err != nil {
		return err
	}

	body, _ := json.Marshal(event)

	// Publish
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

func (s *VehicleServiceImpl) FindLatestByVehicleID(ctx context.Context, vehicleID string) (web.VehicleLocationResponse, error) {
	location, err := s.VehicleRepository.FindLatestByVehicleID(ctx, s.DB, vehicleID)
	if err != nil {
		return web.VehicleLocationResponse{}, err
	}

	return web.VehicleLocationResponse{
		VehicleID: location.VehicleID,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		Timestamp: location.Timestamp,
	}, nil
}

func (s *VehicleServiceImpl) GetHistory(ctx context.Context, vehicleID string, start, end int64) ([]web.VehicleLocationResponse, error) {
	locations, err := s.VehicleRepository.FindHistory(ctx, s.DB, vehicleID, start, end)
	if err != nil {
		return nil, err
	}

	var responses []web.VehicleLocationResponse
	for _, loc := range locations {
		responses = append(responses, web.VehicleLocationResponse{
			VehicleID: loc.VehicleID,
			Latitude:  loc.Latitude,
			Longitude: loc.Longitude,
			Timestamp: loc.Timestamp,
		})
	}
	return responses, nil
}
