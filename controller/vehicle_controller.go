package controller

import (
	"github.com/gin-gonic/gin"
)

type VehicleController interface {
	FindLatestByVehicleID(context *gin.Context)
	GetHistory(context *gin.Context)
}
