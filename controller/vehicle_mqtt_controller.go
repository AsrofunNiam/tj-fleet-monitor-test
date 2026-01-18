package controller

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type VehicleMQTTController interface {
	HandleLocation(client mqtt.Client, msg mqtt.Message)
}
