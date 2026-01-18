package controller

import (
	"context"
	"encoding/json"
	"log"

	"github.com/AsrofunNiam/tj-fleet-monitor-test/model/domain"
	"github.com/AsrofunNiam/tj-fleet-monitor-test/service"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type VehicleMQTTControllerImpl struct {
	LocationService service.LocationService
}

func NewVehicleMQTTController(locationService service.LocationService) VehicleMQTTController {
	return &VehicleMQTTControllerImpl{
		LocationService: locationService,
	}
}

func (c *VehicleMQTTControllerImpl) HandleLocation(client mqtt.Client, msg mqtt.Message) {
	var payload domain.VehicleLocation

	if err := json.Unmarshal(msg.Payload(), &payload); err != nil {
		log.Println("invalid payload:", err)
		return
	}

	err := c.LocationService.SaveLocation(context.Background(), payload)
	if err != nil {
		log.Println("failed save location:", err)
	}
}
