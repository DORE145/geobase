package service

import (
	"github.com/DORE145/geobase/models"
	"github.com/DORE145/geobase/utils/conversion"
	"github.com/sirupsen/logrus"

	"github.com/DORE145/geobase/storage"
)

// IPRangeService is a service that retrieves ip ranges data and processes it if needed
type IPRangeService struct {
	ipStorage       storage.IpRangeStorage
	locationStorage storage.LocationStorage
}

// NewIpRangeService constructs and returns new IPRangeService
func NewIpRangeService(ipStorage storage.IpRangeStorage, locationStorage storage.LocationStorage) *IPRangeService {
	return &IPRangeService{
		ipStorage:       ipStorage,
		locationStorage: locationStorage,
	}
}

// GetLocationByIP returns a location associated with given IP address
func (service *IPRangeService) GetLocationByIP(ip uint32) (*models.Location, error) {
	// Retrieving IP range info with associated location index
	ipRange, err := service.ipStorage.GetIPRange(ip)
	if err != nil {
		logrus.Debugf("Failed to get IP range")
		return nil, err
	} else {
		logrus.Debugf("Found ip range for ip %s", conversion.Uint32toIPString(ip))
	}

	// Retrieving location by its index
	result, err := service.locationStorage.GetByIndex(int(ipRange.LocationIndex))
	if err != nil {
		logrus.Debugf("Failed match ip range and location")
		return nil, err
	} else {
		logrus.Debugf("Found location for ip %s", conversion.Uint32toIPString(ip))
	}

	return result, err
}
