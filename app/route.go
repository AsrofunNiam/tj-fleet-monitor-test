package app

import (
	"fmt"
	"runtime/debug"

	// "github.com/AsrofunNiam/tj-fleet-monitor-test/route"
	"github.com/AsrofunNiam/tj-fleet-monitor-test/exception"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"gorm.io/gorm"
)

//  Handle  mqqt route

// ErrorHandler
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

func NewRouter(db *gorm.DB, validate *validator.Validate) *gin.Engine {

	router := gin.New()

	//  exception middleware
	router.Use(ErrorHandler())
	router.UseRawPath = true

	// route path
	// route.UserRoute(router, db, validate)

	return router
}
