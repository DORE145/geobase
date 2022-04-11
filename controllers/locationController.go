package controllers

import (
	"github.com/DORE145/geobase/models"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/DORE145/geobase/service"
)

type LocationController struct {
	LocationService service.LocationService
}

func NewLocationController(service service.LocationService) LocationController {
	return LocationController{
		LocationService: service,
	}
}

func (controller *LocationController) GetLocationByOrg(ctx *gin.Context) {
	org := ctx.Query("org")
	if org == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Query parameter org is not found",
		})
		return
	}
	location, err := controller.LocationService.GetLocationByOrg(org)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Location not found",
		})
		return
	}

	ctx.JSON(200, location.ToResponse())
}

func (controller *LocationController) GetLocationsByCity(ctx *gin.Context) {
	city := ctx.Query("city")
	if city == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Query parameter city is not found",
		})
		return
	}
	locations, err := controller.LocationService.GetLocationsByCity(city)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Location not found",
		})
		return
	}
	result := make([]models.LocationResp, 0, len(locations))
	for _, location := range locations {
		result = append(result, location.ToResponse())
	}
	ctx.JSON(200, result)
}

func (controller *LocationController) GetLocationsByPostal(ctx *gin.Context) {
	postal := ctx.Query("postal")
	if postal == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Query parameter postal is not found",
		})
		return
	}
	locations, err := controller.LocationService.GetLocationsByPostal(postal)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Location not found",
		})
		return
	}
	result := make([]models.LocationResp, 0, len(locations))
	for _, location := range locations {
		result = append(result, location.ToResponse())
	}
	ctx.JSON(200, result)
}

func (controller *LocationController) GetLocationsByRegion(ctx *gin.Context) {
	region := ctx.Query("region")
	if region == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Query parameter region is not found",
		})
		return
	}
	locations, err := controller.LocationService.GetLocationsByCity(region)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Location not found",
		})
		return
	}
	result := make([]models.LocationResp, 0, len(locations))
	for _, location := range locations {
		result = append(result, location.ToResponse())
	}
	ctx.JSON(200, result)
}

func (controller *LocationController) GetLocationsByCountry(ctx *gin.Context) {
	country := ctx.Query("country")
	if country == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Query parameter country is not found",
		})
		return
	}
	locations, err := controller.LocationService.GetLocationsByCity(country)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Location not found",
		})
		return
	}
	result := make([]models.LocationResp, 0, len(locations))
	for _, location := range locations {
		result = append(result, location.ToResponse())
	}
	ctx.JSON(200, result)
}

func (controller *LocationController) RegisterLocationRouts(group *gin.RouterGroup) {
	group.GET("/org/location", controller.GetLocationByOrg)
	group.GET("/city/locations", controller.GetLocationsByCity)
	group.GET("/postal/locations", controller.GetLocationsByPostal)
	group.GET("/region/locations", controller.GetLocationsByRegion)
	group.GET("/country/locations", controller.GetLocationsByCountry)
}
