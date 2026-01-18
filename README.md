
# Fleet Monitoring System
A real-time fleet monitoring system for TransJakarta buses. This application integrates MQTT for IoT data ingestion, PostgreSQL for data persistence, and RabbitMQ for geofence alerting.

## Architecture Overview
IoT Simulator: A built-in goroutine that publishes mock GPS data to MQTT every 2 seconds.
MQTT Broker (Mosquitto): Handles incoming coordinates on the /fleet/vehicle/{vehicle_id}/location topic.

Backend Service (Go): * Consumes MQTT data and persists it to PostgreSQL.

Calculates real-time distance to a target (Monas).

Triggers an alert to RabbitMQ if the vehicle enters a 50m radius.

RabbitMQ: Stores geofence entry alerts in the geofence_alerts queue.

REST API: Provides endpoints for real-time tracking and historical data.

***

# Editing this README
When you're ready to make this README your own, just edit this file and use the handy template below (or feel free to structure it however you want - this is just a starting point!). Thank you to [makeareadme.com](https://www.makeareadme.com/) for this template.

## Suggestions for a good README
Every project is different, so consider which of these sections apply to yours. The sections used in the template are suggestions for most open source projects. Also keep in mind that while a README can be too long and detailed, too long is better than too short. If you think your README is too long, consider utilizing another form of documentation rather than cutting out information.

## Name
Choose a github.com/AsrofunNiam/tj-fleet-monitor-test name for your project.
 
## Badges
On some READMEs, you may see small images that convey metadata, such as whether or not all the tests are passing for the project. You can use Shields to add some to your README. Many services also have instructions for adding a badge.

## Getting Started
1. Prerequisites
Docker and Docker Compose installed.

Postman for API testing.

2. Run with Local 
You have to make sure file .env is plase on configuration/.env

Execute the following command to run the project:

```shellscript
go get
go run main.go
``` 

3. Run with Docker Compose
To start the entire stack (Database, Broker, RabbitMQ, and the Go App), run:

```shellscript
docker-compose up -d --build
```

4. Service Access Points
REST API: http://localhost:8089

RabbitMQ Management UI: http://localhost:15672 (User/Pass: guest / guest)

MQTT Broker: localhost:1883

## API Documentation
1. Get Last Location
Retrieve the most recent coordinates of a specific vehicle.
```json
GET /vehicles/:vehicle_id/location
```

Example: http://localhost:8089/vehicles/B1234XYZ/location

2. Get Location History
Retrieve all movement history within a specific time range.

```json
GET /vehicles/:vehicle_id/history
```

Query Params: * start: Unix timestamp (e.g., 1737216000)

end: Unix timestamp (e.g., 1737302400)

Example: http://localhost:8089/vehicles/B1234XYZ/history?start=1700000000&end=1800000000

## Testing the System
Automated Mock Data
Once the application starts, a simulator automatically sends data for vehicle B1234XYZ. You will see logs in the console: [SIMULATOR] Sent: {"latitude": -6.2088, "longitude": 106.8456, ...}

Verifying Geofence Alerts
Log in to the RabbitMQ Management UI.

Navigate to the Queues tab.

Locate the geofence_alerts queue.

As the simulator moves the vehicle near Monas, you will see the Message Ready count increase.

## Postman Collection (v1.0)
Copy this JSON and import it into Postman (Import > Raw Text):

```json

{
	"info": {
		"name": "TJ Fleet Monitor",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	},
	"item": [
		{
			"name": "Vehicle Last Location",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "http://localhost:8089/vehicles/B1234XYZ/location",
					"host": ["localhost"],
					"port": "8089",
					"path": ["vehicles", "B1234XYZ", "location"]
				}
			}
		},
		{
			"name": "Vehicle History",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "http://localhost:8089/vehicles/B1234XYZ/history?start=1600000000&end=1800000000",
					"host": ["localhost"],
					"port": "8089",
					"path": ["vehicles", "B1234XYZ", "history"],
					"query": [
						{ "key": "start", "value": "1600000000" },
						{ "key": "end", "value": "1800000000" }
					]
				}
			}
		}
	]
}
```

