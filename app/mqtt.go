package app

import (
	"github.com/AsrofunNiam/tj-fleet-monitor-test/route"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func NewMQTTClient(mqttRoute *route.MQTTRoute) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://mosquitto:1883")
	opts.SetClientID("fleet-backend")

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	mqttRoute.Register(client)

	return client
}
