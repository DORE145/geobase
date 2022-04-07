package storage

import "github.com/DORE145/geobase/models"

type LocationStorage interface {
	GetByOrg(string) (*models.Location, error)
	GetByCity(string) ([]*models.Location, error)
	GetByPostal(string) ([]*models.Location, error)
	GetByRegion(string) ([]*models.Location, error)
	GetByCountry(string) ([]*models.Location, error)
}
