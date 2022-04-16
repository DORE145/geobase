package controllers

import (
	"github.com/DORE145/geobase/service"
	"github.com/DORE145/geobase/utils/conversion"
	"github.com/gin-gonic/gin"
	"net/http"
)

// IPRangeController prepares responses for ip ranges endpoints
type IPRangeController struct {
	ipService service.IPRangeService
}

// NewIPRangeController returns new IPRangeController
func NewIPRangeController(service service.IPRangeService) IPRangeController {
	return IPRangeController{
		ipService: service,
	}
}

// GetIPRange is a handler that serves /ip/location route
func (controller *IPRangeController) GetIPRange(ctx *gin.Context) {
	ipString := ctx.Query("ip")
	if ipString == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Query parameter ip is not found",
		})
		return
	}
	ip, err := conversion.IPStringToUint32(ipString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Unparsable ip address provided",
		})
		return
	}
	location, err := controller.ipService.GetLocationByIP(ip)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Location not found",
		})
		return
	}

	ctx.JSON(200, location.ToResponse())

}

// RegisterIPRangeRoutes registers all routes related to ip ranges
func (controller *IPRangeController) RegisterIPRangeRoutes(group *gin.RouterGroup) {
	group.GET("/ip/location", controller.GetIPRange)
}
