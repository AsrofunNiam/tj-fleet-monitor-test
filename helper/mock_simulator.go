package helper

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func RunVehicleSimulator(client mqtt.Client) {
	vehicleID := "B1234XYZ"
	baseLat := -6.2088
	baseLong := 106.8456

	fmt.Println("simulator Started: Sending data every 2 seconds...")

	for {
		offsetLat := (rand.Float64() - 0.5) * 0.001
		offsetLong := (rand.Float64() - 0.5) * 0.001

		payload := map[string]interface{}{
			"vehicle_id": vehicleID,
			"latitude":   baseLat + offsetLat,
			"longitude":  baseLong + offsetLong,
			"timestamp":  time.Now().Unix(),
		}

		data, _ := json.Marshal(payload)
		topic := fmt.Sprintf("/fleet/vehicle/%s/location", vehicleID)

		token := client.Publish(topic, 0, false, data)
		token.Wait()

		// Log ke console backend untuk memastikan simulator jalan
		fmt.Printf("[SIMULATOR] sent to %s: %s\n", topic, string(data))

		time.Sleep(2 * time.Second)
	}
}
