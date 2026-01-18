package main

import (
	"log"
	"net/http"

	"github.com/AsrofunNiam/tj-fleet-monitor-test/app"
	c "github.com/AsrofunNiam/tj-fleet-monitor-test/configuration"
	"github.com/go-playground/validator/v10"
)

func main() {
	cfg, err := c.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db := app.ConnectDatabase(cfg.User, cfg.Host, cfg.Password, cfg.PortDB, cfg.Db)

	// RabbitMQ
	rabbitConn := app.NewRabbitMQConnection(cfg.RabbitMQURL)
	defer rabbitConn.Close()

	validate := validator.New()

	// Gin
	router := app.NewRouter(db, rabbitConn, validate)

	// MQTT
	go app.InitApplication(db, rabbitConn, validate, cfg.MQTTBroker)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	log.Printf("Server running on port %s", cfg.Port)
	log.Fatal(server.ListenAndServe())
}
