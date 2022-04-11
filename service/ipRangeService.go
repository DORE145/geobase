package service

import (
	"github.com/DORE145/geobase/models"
	"github.com/DORE145/geobase/utils/conversion"
	"github.com/sirupsen/logrus"

	"github.com/DORE145/geobase/storage"
)

type IPRangeService struct {
	ipStorage       storage.IpRangeStorage
	locationStorage storage.LocationStorage
}

func NewIpRangeService(ipStorage storage.IpRangeStorage, locationStorage storage.LocationStorage) *IPRangeService {
	return &IPRangeService{
		ipStorage:       ipStorage,
		locationStorage: locationStorage,
	}
}

func (service *IPRangeService) GetLocationByIP(ip uint32) (*models.Location, error) {
	ipRange, err := service.ipStorage.GetIPRange(ip)
	if err != nil {
		logrus.Debugf("Failed to get IP range")
		return nil, err
	} else {
		logrus.Debugf("Found ip range for ip %s", conversion.Uint32toIPString(ip))
	}

	result, err := service.locationStorage.GetByIndex(int(ipRange.LocationIndex))
	if err != nil {
		logrus.Debugf("Failed match ip range and location")
		return nil, err
	} else {
		logrus.Debugf("Found location for ip %s", conversion.Uint32toIPString(ip))
	}
	return result, err
}
