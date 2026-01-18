package app

import (
	"fmt"
	"runtime/debug"

	"github.com/AsrofunNiam/tj-fleet-monitor-test/exception"
	"github.com/AsrofunNiam/tj-fleet-monitor-test/route"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	amqp "github.com/rabbitmq/amqp091-go"

	"gorm.io/gorm"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
				exception.ErrorHandler(c, err)
			}
		}()
		c.Next()
	}
}

func NewRouter(db *gorm.DB, rabbitConn *amqp.Connection, validate *validator.Validate) *gin.Engine {

	router := gin.New()

	//  exception middleware
	router.Use(ErrorHandler())
	router.UseRawPath = true

	// route path
	route.VehicleRoute(router, db, rabbitConn, validate)

	return router
}
