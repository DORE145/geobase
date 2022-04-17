package controllers

import (
	"github.com/DORE145/geobase/models"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/DORE145/geobase/service"
)

// LocationController prepares responses for location endpoints
type LocationController struct {
	LocationService service.LocationService
}

// NewLocationController creates new LocationController
func NewLocationController(service service.LocationService) LocationController {
	return LocationController{
		LocationService: service,
	}
}

// GetLocationByOrg is a handler that serves /org/location endpoint
// swagger:route GET /org/location Location orgLocation
//
// Returns a location for an organization
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Parameters:
//       + name: org
//         in: query
//         description: Organization to find location for
//         required: true
//         type: string
//
//     Responses:
//       200: location
//       400: badParameter
//		 404: notFound
func (controller *LocationController) GetLocationByOrg(ctx *gin.Context) {
	org := ctx.Query("org")
	if org == "" {
		ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{Message: "Query parameter org is not found"})
		return
	}
	location, err := controller.LocationService.GetLocationByOrg(org)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NotFoundResponse{Message: "Location not found"})
		return
	}

	ctx.JSON(200, location.ToResponse())
}

// GetLocationsByCity is a handler that serves /city/locations endpoint
// swagger:route GET /city/location Location cityLocations
//
// Returns all locations in a city
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Parameters:
//       + name: city
//         in: query
//         description: City to find locations in
//         required: true
//         type: string
//
//     Responses:
//       200: locations
//       400: badParameter
//		 404: notFound
func (controller *LocationController) GetLocationsByCity(ctx *gin.Context) {
	city := ctx.Query("city")
	if city == "" {
		ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{Message: "Query parameter city is not found"})
		return
	}
	var locations models.Locations
	locations, err := controller.LocationService.GetLocationsByCity(city)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NotFoundResponse{Message: "Location not found"})
		return
	}
	result := locations.ToResponse()
	ctx.JSON(200, result)
}

// GetLocationsByPostal is a handler that serves /postal/locations endpoint
// swagger:route GET /postal/location Location postalLocations
//
// Returns all locations with specified postal code
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Parameters:
//       + name: postal
//         in: query
//         description: Postal code to get locations with
//         required: true
//         type: string
//
//     Responses:
//       200: locations
//       400: badParameter
//		 404: notFound
func (controller *LocationController) GetLocationsByPostal(ctx *gin.Context) {
	postal := ctx.Query("postal")
	if postal == "" {
		ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{Message: "Query parameter postal is not found"})
		return
	}
	var locations models.Locations
	locations, err := controller.LocationService.GetLocationsByPostal(postal)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NotFoundResponse{Message: "Location not found"})

		return
	}

	result := locations.ToResponse()
	ctx.JSON(200, result)
}

// GetLocationsByRegion is a handler that serves /region/locations endpoint
// swagger:route GET /region/location Location regionLocations
//
// Returns all locations from specified region
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Parameters:
//       + name: region
//         in: query
//         description: Region to get locations from
//         required: true
//         type: string
//
//     Responses:
//       200: locations
//       400: badParameter
//		 404: notFound
func (controller *LocationController) GetLocationsByRegion(ctx *gin.Context) {
	region := ctx.Query("region")
	if region == "" {
		ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{Message: "Query parameter region is not found"})
		return
	}
	var locations models.Locations
	locations, err := controller.LocationService.GetLocationsByCity(region)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NotFoundResponse{Message: "Location not found"})
		return
	}

	result := locations.ToResponse()
	ctx.JSON(200, result)
}

// GetLocationsByCountry is a handler that serves /country/locations endpoint
// swagger:route GET /country/location Location countryLocations
//
// Returns all locations from specified country
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Parameters:
//       + name: country
//         in: query
//         description: Country to get locations from
//         required: true
//         type: string
//
//     Responses:
//       200: locations
//       400: badParameter
//		 404: notFound
func (controller *LocationController) GetLocationsByCountry(ctx *gin.Context) {
	country := ctx.Query("country")
	if country == "" {
		ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{Message: "Query parameter country is not found"})
		return
	}

	var locations models.Locations
	locations, err := controller.LocationService.GetLocationsByCity(country)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NotFoundResponse{Message: "Location not found"})
		return
	}

	result := locations.ToResponse()
	ctx.JSON(200, result)
}

func (controller *LocationController) RegisterLocationRouts(group *gin.RouterGroup) {
	group.GET("/org/location", controller.GetLocationByOrg)
	group.GET("/city/locations", controller.GetLocationsByCity)
	group.GET("/postal/locations", controller.GetLocationsByPostal)
	group.GET("/region/locations", controller.GetLocationsByRegion)
	group.GET("/country/locations", controller.GetLocationsByCountry)
}
