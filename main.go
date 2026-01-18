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

	db := app.ConnectDatabase(
		cfg.User,
		cfg.Host,
		cfg.Password,
		cfg.PortDB,
		cfg.Db,
	)

	validate := validator.New()
	router := app.NewRouter(db, validate)

	go app.InitApplication(db)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	log.Printf("Server running on port %s", cfg.Port)
	log.Fatal(server.ListenAndServe())
}
