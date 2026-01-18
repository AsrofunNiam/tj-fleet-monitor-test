package route

import (
	"github.com/AsrofunNiam/tj-fleet-monitor-test/controller"
	"github.com/AsrofunNiam/tj-fleet-monitor-test/repository"
	"github.com/AsrofunNiam/tj-fleet-monitor-test/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

func VehicleRoute(router *gin.Engine, db *gorm.DB, rabbitConn *amqp.Connection, validate *validator.Validate) {
	Vehicles := service.NewVehicleService(
		repository.NewVehicleRepository(),
		db,
		rabbitConn,
		validate,
	)
	productController := controller.NewVehicleController(Vehicles)
	router.GET("/vehicles/:vehicle_id/location", productController.FindLatestByVehicleID)
	router.GET("/vehicles/:vehicle_id/history", productController.GetHistory)
}
