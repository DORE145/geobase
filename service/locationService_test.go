package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sort"
	"testing"

	"github.com/DORE145/geobase/models"
	"github.com/DORE145/geobase/storage"
	"github.com/DORE145/geobase/storage/inmemory"
	"github.com/DORE145/geobase/utils/sorters"
)

func TestNewLocationService(t *testing.T) {
	locationStorage, err := getLocationStorage()
	require.NoError(t, err)
	require.NotNil(t, locationStorage)

	service := NewLocationService(locationStorage)
	assert.NotNil(t, service)
}

func TestLocationService_GetLocationByOrg(t *testing.T) {
	locationStorage, err := getLocationStorage()
	require.NoError(t, err)
	require.NotNil(t, locationStorage)

	service := NewLocationService(locationStorage)
	require.NotNil(t, service)

	location, err := service.GetLocationByOrg("org_Usikywyjajyj")
	assert.NoError(t, err)
	assert.NotNil(t, location)
	var expected [32]byte
	copy(expected[:], "org_Usikywyjajyj")
	assert.Equal(t, expected, location.Organization)

	location, err = service.GetLocationByOrg("something")
	assert.Error(t, err)
	assert.Nil(t, location)
}

func TestLocationService_GetLocationsByCity(t *testing.T) {
	locationStorage, err := getLocationStorage()
	require.NoError(t, err)
	require.NotNil(t, locationStorage)

	service := NewLocationService(locationStorage)
	require.NotNil(t, service)

	locations, err := service.GetLocationsByCity("cit_Elu")
	assert.NoError(t, err)
	assert.NotNil(t, locations)
	assert.Len(t, locations, 1)

	var expected [24]byte
	copy(expected[:], "cit_Elu")
	assert.Equal(t, expected, locations[0].City)

	locations, err = service.GetLocationsByCity("something")
	assert.Error(t, err)
	assert.Nil(t, locations)
}

func TestLocationService_GetLocationsByCountry(t *testing.T) {
	locationStorage, err := getLocationStorage()
	require.NoError(t, err)
	require.NotNil(t, locationStorage)

	service := NewLocationService(locationStorage)
	require.NotNil(t, service)

	locations, err := service.GetLocationsByCountry("cou_AK")
	assert.NoError(t, err)
	assert.NotNil(t, locations)
	require.Len(t, locations, 1)

	var expected [8]byte
	copy(expected[:], "cou_AK")
	assert.Equal(t, expected, locations[0].Country)

	locations, err = locationStorage.GetByCountry("something")
	assert.Error(t, err)
	assert.Nil(t, locations)
}

func TestLocationService_GetLocationsByPostal(t *testing.T) {
	locationStorage, err := getLocationStorage()
	require.NoError(t, err)
	require.NotNil(t, locationStorage)

	service := NewLocationService(locationStorage)
	require.NotNil(t, service)

	locations, err := service.GetLocationsByPostal("pos_8731")
	assert.NoError(t, err)
	assert.NotNil(t, locations)
	require.Len(t, locations, 2)

	var expected [12]byte
	copy(expected[:], "pos_8731")
	assert.Equal(t, expected, locations[0].Postal)
	assert.Equal(t, expected, locations[1].Postal)

	locations, err = service.GetLocationsByPostal("something")
	assert.Error(t, err)
	assert.Nil(t, locations)
}

func TestLocationService_GetLocationsByRegion(t *testing.T) {
	locationStorage, err := getLocationStorage()
	require.NoError(t, err)
	require.NotNil(t, locationStorage)

	service := NewLocationService(locationStorage)
	require.NotNil(t, service)

	locations, err := service.GetLocationsByRegion("reg_Alimos")
	assert.NoError(t, err)
	assert.NotNil(t, locations)
	require.Len(t, locations, 1)

	var expected [12]byte
	copy(expected[:], "reg_Alimos")
	assert.Equal(t, expected, locations[0].Region)

	locations, err = service.GetLocationsByRegion("something")
	assert.Error(t, err)
	assert.Nil(t, locations)
}

func getLocationStorage() (storage.LocationStorage, error) {
	locations := make([]*models.Location, 0)
	var country [8]byte
	var region [12]byte
	var postal [12]byte
	var city [24]byte
	var org [32]byte

	copy(country[:], "cou_UJO")
	copy(region[:], "reg_U")
	copy(postal[:], "pos_582423")
	copy(city[:], "cit_Elu")
	copy(org[:], "org_Eba Abacir L")

	locations = append(locations, &models.Location{
		Country:      country,
		Region:       region,
		Postal:       postal,
		City:         city,
		Organization: org,
		Latitude:     -96.2552,
		Longitude:    -51.5246,
	})

	country = [8]byte{}
	region = [12]byte{}
	postal = [12]byte{}
	city = [24]byte{}
	org = [32]byte{}

	copy(country[:], "cou_EDE")
	copy(region[:], "reg_Yjema ")
	copy(postal[:], "pos_96188")
	copy(city[:], "cit_Ujami ")
	copy(org[:], "org_Axaw ")

	locations = append(locations, &models.Location{
		Country:      country,
		Region:       region,
		Postal:       postal,
		City:         city,
		Organization: org,
		Latitude:     133.968,
		Longitude:    -111.9784,
	})

	country = [8]byte{}
	region = [12]byte{}
	postal = [12]byte{}
	city = [24]byte{}
	org = [32]byte{}
	copy(country[:], "cou_AK")
	copy(region[:], "reg_I")
	copy(postal[:], "pos_0679")
	copy(city[:], "cit_Oqyhys")
	copy(org[:], "org_Ociperi Asetyn")

	locations = append(locations, &models.Location{
		Country:      country,
		Region:       region,
		Postal:       postal,
		City:         city,
		Organization: org,
		Latitude:     -54.5288,
		Longitude:    -53.1396,
	})

	country = [8]byte{}
	region = [12]byte{}
	postal = [12]byte{}
	city = [24]byte{}
	org = [32]byte{}
	copy(country[:], "cou_OMO")
	copy(region[:], "reg_Alimos")
	copy(postal[:], "pos_0327")
	copy(city[:], "cit_I Yfuqahil Hehela")
	copy(org[:], "org_Ynyn Hafuk Bodode")

	locations = append(locations, &models.Location{
		Country:      country,
		Region:       region,
		Postal:       postal,
		City:         city,
		Organization: org,
		Latitude:     118.2995,
		Longitude:    34.4529,
	})

	country = [8]byte{}
	region = [12]byte{}
	postal = [12]byte{}
	city = [24]byte{}
	org = [32]byte{}
	copy(country[:], "cou_UCY")
	copy(region[:], "reg_O Y")
	copy(postal[:], "pos_8731")
	copy(city[:], "cit_Uwol Z Hyt Xavi")
	copy(org[:], "org_Usikywyjajyj")

	locations = append(locations, &models.Location{
		Country:      country,
		Region:       region,
		Postal:       postal,
		City:         city,
		Organization: org,
		Latitude:     -152.6465,
		Longitude:    -45.284,
	})

	country = [8]byte{}
	region = [12]byte{}
	postal = [12]byte{}
	city = [24]byte{}
	org = [32]byte{}
	copy(country[:], "cou_UCY2")
	copy(region[:], "reg_O Y2")
	copy(postal[:], "pos_8731")
	copy(city[:], "cit_Uwol Z Hyt Xavi2")
	copy(org[:], "org_Usikywyjajyj2")

	locations = append(locations, &models.Location{
		Country:      country,
		Region:       region,
		Postal:       postal,
		City:         city,
		Organization: org,
		Latitude:     -152.6465,
		Longitude:    -45.284,
	})

	locationsIndex := make([]*models.Location, 0)
	locationsIndex = append(locationsIndex, locations...)

	sort.Sort(sorters.ByCity(locationsIndex))

	return inmemory.NewLocationStorage(locations, locationsIndex)

}
