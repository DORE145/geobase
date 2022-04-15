package service

import (
	"github.com/DORE145/geobase/models"
	"github.com/DORE145/geobase/storage"
	"github.com/DORE145/geobase/storage/inmemory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewIpRangeService(t *testing.T) {
	ipStorage := getIPRangeStorage()
	locationStorage, err := getLocationStorage()
	require.NoError(t, err)

	service := NewIpRangeService(ipStorage, locationStorage)
	assert.NotNil(t, service)
}

func TestIPRangeService_GetLocationByIP(t *testing.T) {
	ipStorage := getIPRangeStorage()
	locationStorage, err := getLocationStorage()
	require.NoError(t, err)

	service := NewIpRangeService(ipStorage, locationStorage)
	require.NotNil(t, service)

	location, err := service.GetLocationByIP(55)
	require.NoError(t, err)
	require.NotNil(t, location)

	var expected [8]byte
	copy(expected[:], "cou_UJO")
	assert.Equal(t, expected, location.Country)

	location, err = service.GetLocationByIP(99999)
	require.NoError(t, err)
	require.NotNil(t, location)

	expected = [8]byte{}
	copy(expected[:], "cou_EDE")
	assert.Equal(t, expected, location.Country)

	location, err = service.GetLocationByIP(999999)
	assert.Error(t, err)
	assert.Nil(t, location)

	location, err = service.GetLocationByIP(0)
	assert.Error(t, err)
	assert.Nil(t, location)
}

func getIPRangeStorage() storage.IpRangeStorage {
	data := []*models.IpRange{
		{
			IpFrom:        8,
			IpTo:          55472,
			LocationIndex: 0,
		},
		{
			IpFrom:        55473,
			IpTo:          151737,
			LocationIndex: 1,
		},
		{
			IpFrom:        151738,
			IpTo:          191323,
			LocationIndex: 2,
		},
	}
	storage := inmemory.NewIpRangeStorage(data)

	return storage
}
