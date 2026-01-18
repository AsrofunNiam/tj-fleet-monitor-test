package route

import (
	"github.com/AsrofunNiam/tj-fleet-monitor-test/controller"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTRoute struct {
	VehicleController controller.VehicleMQTTController
}

func NewMQTTRoute(
	vehicleController controller.VehicleMQTTController,
) *MQTTRoute {
	return &MQTTRoute{
		VehicleController: vehicleController,
	}
}

func (r *MQTTRoute) Register(client mqtt.Client) {
	client.Subscribe(
		"/fleet/vehicle/+/location",
		1,
		r.VehicleController.HandleLocation,
	)
}
