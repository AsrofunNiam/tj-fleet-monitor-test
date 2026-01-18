package app

import (
	"log"

	"github.com/AsrofunNiam/tj-fleet-monitor-test/controller"
	"github.com/AsrofunNiam/tj-fleet-monitor-test/repository"
	route "github.com/AsrofunNiam/tj-fleet-monitor-test/route"
	"github.com/AsrofunNiam/tj-fleet-monitor-test/service"
	"github.com/go-playground/validator/v10"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

func InitApplication(db *gorm.DB, rabbitConn *amqp.Connection, validate *validator.Validate, mqqtBroker string) {
	log.Println("InitApplication started")

	if db == nil {
		log.Fatal("DB is nil in InitApplication")
	}

	log.Println("Creating LocationService")
	vehicleMQTTController := controller.NewVehicleMQTTController(
		service.NewVehicleService(
			repository.NewVehicleRepository(),
			db,
			rabbitConn,
			validate,
		))

	log.Println("Creating MQTTRoute")
	mqttRoute := route.NewMQTTRoute(vehicleMQTTController)

	log.Println("Connecting MQTT")
	NewMQTTClient(mqttRoute, mqqtBroker)

	log.Println("MQTT connected")
}
