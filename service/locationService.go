package service

import (
	"context"
	"github.com/DORE145/geobase/models"
	"github.com/DORE145/geobase/storage"
	"github.com/sirupsen/logrus"
)

type LocationService struct {
	storage storage.LocationStorage
	ctx     context.Context
}

func NewLocationService(storage storage.LocationStorage, ctx context.Context) *LocationService {
	return &LocationService{
		storage: storage,
		ctx:     ctx,
	}
}

func (service *LocationService) GetLocationByOrg(org string) (*models.Location, error) {
	result, err := service.storage.GetByOrg(org)
	if err != nil {
		logrus.Debugf("Failed to found location for orginaization %s", org)
	} else {
		logrus.Debugf("Found a location for organization %s", org)
	}
	return result, err
}

func (service *LocationService) GetLocationsByCity(city string) ([]*models.Location, error) {
	result, err := service.storage.GetByCity(city)
	if err != nil {
		logrus.Debugf("Failed to found locations in city %s", city)
	} else {
		logrus.Debugf("Found %d location(s) in city %s", len(result), city)
	}
	return result, err
}

func (service *LocationService) GetLocationsByPostal(postal string) ([]*models.Location, error) {
	result, err := service.storage.GetByPostal(postal)
	if err != nil {
		logrus.Debugf("Failed to found locations with postal code %s", postal)
	} else {
		logrus.Debugf("Found %d locations for organization %s", len(result), postal)
	}
	return result, err
}
func (service *LocationService) GetLocationsByRegion(region string) ([]*models.Location, error) {
	result, err := service.storage.GetByRegion(region)
	if err != nil {
		logrus.Debugf("Failed to found locations in %s region", region)
	} else {
		logrus.Debugf("Found %d locations in %s region", len(result), region)
	}
	return result, err
}
func (service *LocationService) GetLocationsByCountry(country string) ([]*models.Location, error) {
	result, err := service.storage.GetByCountry(country)
	if err != nil {
		logrus.Debugf("Failed to found locations in %s country", country)
	} else {
		logrus.Debugf("Found %d locations in %s country", len(result), country)
	}
	return result, err
}
