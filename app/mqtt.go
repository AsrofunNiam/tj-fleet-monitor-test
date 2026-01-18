package app

import (
	"github.com/AsrofunNiam/tj-fleet-monitor-test/route"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func NewMQTTClient(mqttRoute *route.MQTTRoute, mqttBroker string) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(mqttBroker)
	opts.SetClientID("fleet-backend")

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	mqttRoute.Register(client)

	return client
}
