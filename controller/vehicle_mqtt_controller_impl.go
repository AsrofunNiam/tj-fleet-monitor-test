package controller

import (
	"context"
	"encoding/json"
	"log"

	"github.com/AsrofunNiam/tj-fleet-monitor-test/model/web"
	"github.com/AsrofunNiam/tj-fleet-monitor-test/service"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type VehicleMQTTControllerImpl struct {
	VehicleService service.VehicleService
}

func NewVehicleMQTTController(locationService service.VehicleService) VehicleMQTTController {
	return &VehicleMQTTControllerImpl{
		VehicleService: locationService,
	}
}

func (c *VehicleMQTTControllerImpl) HandleLocation(client mqtt.Client, msg mqtt.Message) {
	var payload web.VehicleLocationCreateRequest

	if err := json.Unmarshal(msg.Payload(), &payload); err != nil {
		log.Println("invalid payload:", err)
		return
	}

	err := c.VehicleService.SaveLocation(context.Background(), payload)
	if err != nil {
		log.Println("failed save location:", err)
	}
}
