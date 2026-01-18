package controller

import (
	"net/http"
	"strconv"

	"github.com/AsrofunNiam/tj-fleet-monitor-test/model/web"
	"github.com/AsrofunNiam/tj-fleet-monitor-test/service"
	"github.com/gin-gonic/gin"
)

type VehicleControllerImpl struct {
	VehicleService service.VehicleService
}

func NewVehicleController(locationService service.VehicleService) VehicleController {
	return &VehicleControllerImpl{
		VehicleService: locationService,
	}
}

func (controller *VehicleControllerImpl) GetHistory(c *gin.Context) {
	vehicleID := c.Param("vehicle_id")

	startStr := c.Query("start")
	endStr := c.Query("end")

	start, _ := strconv.ParseInt(startStr, 10, 64)
	end, _ := strconv.ParseInt(endStr, 10, 64)

	historyResponses, err := controller.VehicleService.GetHistory(c.Request.Context(), vehicleID, start, end)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Success:   true,
		Message:   "Data found",
		TotalData: len(historyResponses),
		Data:      historyResponses,
	})
}

func (controller *VehicleControllerImpl) FindLatestByVehicleID(c *gin.Context) {
	vehicleID := c.Param("vehicle_id")

	locationResponse, err := controller.VehicleService.FindLatestByVehicleID(c.Request.Context(), vehicleID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Success: true,
		Message: "Data found",
		Data:    locationResponse,
	})
}
