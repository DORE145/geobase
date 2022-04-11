package controllers

import (
	"github.com/DORE145/geobase/service"
	"github.com/gin-gonic/gin"
)

type IPRangeController struct {
	ipService service.IPRangeService
}

func NewIPRangeController(service service.IPRangeService) IPRangeController {
	return IPRangeController{
		ipService: service,
	}
}

func (controller *IPRangeController) GetIPRange(ctx *gin.Context) {
	ctx.JSON(200, nil)
}

func (controller *IPRangeController) RegisterIPRangeRoutes(group *gin.RouterGroup) {
	group.GET("/ip/location", controller.GetIPRange)
}
