package controllers

import (
	"github.com/DORE145/geobase/service"
	"github.com/gin-gonic/gin"
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
	ctx.JSON(200, nil)
}

// RegisterIPRangeRoutes registers all routes related to ip ranges
func (controller *IPRangeController) RegisterIPRangeRoutes(group *gin.RouterGroup) {
	group.GET("/ip/location", controller.GetIPRange)
}
