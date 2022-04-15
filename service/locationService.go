package service

import (
	"github.com/DORE145/geobase/models"
	"github.com/DORE145/geobase/storage"
	"github.com/sirupsen/logrus"
)

// LocationService is a service that retrieves locations data and processes it if needed
type LocationService struct {
	storage storage.LocationStorage
}

// NewLocationService constructs and returns new LocationService
func NewLocationService(storage storage.LocationStorage) *LocationService {
	return &LocationService{
		storage: storage,
	}
}

// GetLocationByOrg retrieves location of organization
func (service *LocationService) GetLocationByOrg(org string) (*models.Location, error) {
	result, err := service.storage.GetByOrg(org)
	if err != nil {
		logrus.Debugf("Failed to found location for orginaization %s", org)
	} else {
		logrus.Debugf("Found a location for organization %s", org)
	}
	return result, err
}

// GetLocationsByCity retrieves all locations from a city
func (service *LocationService) GetLocationsByCity(city string) ([]*models.Location, error) {
	result, err := service.storage.GetByCity(city)
	if err != nil {
		logrus.Debugf("Failed to found locations in city %s", city)
	} else {
		logrus.Debugf("Found %d location(s) in city %s", len(result), city)
	}
	return result, err
}

// GetLocationsByPostal retrieves all locations with a specific postal code
func (service *LocationService) GetLocationsByPostal(postal string) ([]*models.Location, error) {
	result, err := service.storage.GetByPostal(postal)
	if err != nil {
		logrus.Debugf("Failed to found locations with postal code %s", postal)
	} else {
		logrus.Debugf("Found %d locations for organization %s", len(result), postal)
	}
	return result, err
}

// GetLocationsByRegion retrieves all locations in a specific region
func (service *LocationService) GetLocationsByRegion(region string) ([]*models.Location, error) {
	result, err := service.storage.GetByRegion(region)
	if err != nil {
		logrus.Debugf("Failed to found locations in %s region", region)
	} else {
		logrus.Debugf("Found %d locations in %s region", len(result), region)
	}
	return result, err
}

// GetLocationsByCountry  retrieves all locations in a specific country
func (service *LocationService) GetLocationsByCountry(country string) ([]*models.Location, error) {
	result, err := service.storage.GetByCountry(country)
	if err != nil {
		logrus.Debugf("Failed to found locations in %s country", country)
	} else {
		logrus.Debugf("Found %d locations in %s country", len(result), country)
	}
	return result, err
}
