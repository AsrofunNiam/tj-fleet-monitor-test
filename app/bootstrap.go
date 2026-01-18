package app

import (
	"log"

	"github.com/AsrofunNiam/tj-fleet-monitor-test/controller"
	"github.com/AsrofunNiam/tj-fleet-monitor-test/repository"
	route "github.com/AsrofunNiam/tj-fleet-monitor-test/route"
	"github.com/AsrofunNiam/tj-fleet-monitor-test/service"
	"gorm.io/gorm"
)

func InitApplication(db *gorm.DB) {
	log.Println("InitApplication started")

	if db == nil {
		log.Fatal("DB is nil in InitApplication")
	}

	log.Println("Creating LocationService")
	vehicleMQTTController := controller.NewVehicleMQTTController(
		service.NewLocationService(
			repository.NewLocationRepository(),
			db,
		))

	log.Println("Creating MQTTRoute")
	mqttRoute := route.NewMQTTRoute(vehicleMQTTController)

	log.Println("Connecting MQTT")
	NewMQTTClient(mqttRoute)

	log.Println("MQTT connected")
}
