package inmemory

import (
	"bytes"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/DORE145/geobase/models"
	"github.com/DORE145/geobase/utils/sorters"
)

func TestNewLocationStorage(t *testing.T) {
	data := getTestData()
	locations := make([]*models.Location, 0)
	locations = append(locations, data...)

	sort.Sort(sorters.ByCity(data))
	storage, err := NewLocationStorage(locations, data)
	assert.NoError(t, err)
	assert.Equal(t, len(data), len(storage.cityIndex))

	assert.True(t, sort.SliceIsSorted(storage.postalIndex, func(i, j int) bool {
		return bytes.Compare(storage.postalIndex[i].Postal[:], storage.postalIndex[j].Postal[:]) < 0
	}))
	assert.True(t, sort.SliceIsSorted(storage.regionIndex, func(i, j int) bool {
		return bytes.Compare(storage.regionIndex[i].Region[:], storage.regionIndex[j].Region[:]) < 0
	}))
	assert.True(t, sort.SliceIsSorted(storage.countryIndex, func(i, j int) bool {
		return bytes.Compare(storage.countryIndex[i].Country[:], storage.countryIndex[j].Country[:]) < 0
	}))
	assert.True(t, sort.SliceIsSorted(storage.orgIndex, func(i, j int) bool {
		return bytes.Compare(storage.orgIndex[i].Organization[:], storage.orgIndex[j].Organization[:]) < 0
	}))
	assert.True(t, sort.SliceIsSorted(storage.cityIndex, func(i, j int) bool {
		return bytes.Compare(storage.cityIndex[i].City[:], storage.cityIndex[j].City[:]) < 0
	}))

}

func TestLocationStorage_GetByCity(t *testing.T) {
	data := getTestData()
	locations := make([]*models.Location, 0)
	locations = append(locations, data...)

	sort.Sort(sorters.ByCity(data))
	storage, _ := NewLocationStorage(locations, data)

	item, err := storage.GetByCity("cit_Elu")
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Len(t, item, 1)

	var expected [24]byte
	copy(expected[:], "cit_Elu")
	assert.Equal(t, expected, item[0].City)

	item, err = storage.GetByCity("something")
	assert.Error(t, err)
	assert.Nil(t, item)
}

func TestLocationStorage_GetByCountry(t *testing.T) {
	data := getTestData()
	locations := make([]*models.Location, 0)
	locations = append(locations, data...)

	sort.Sort(sorters.ByCity(data))
	storage, err := NewLocationStorage(locations, data)
	require.NoError(t, err)

	item, err := storage.GetByCountry("cou_AK")
	assert.NoError(t, err)
	assert.NotNil(t, item)
	require.Len(t, item, 1)

	var expected [8]byte
	copy(expected[:], "cou_AK")
	assert.Equal(t, expected, item[0].Country)

	item, err = storage.GetByCountry("something")
	assert.Error(t, err)
	assert.Nil(t, item)
}

func TestLocationStorage_GetByOrg(t *testing.T) {
	data := getTestData()
	locations := make([]*models.Location, 0)
	locations = append(locations, data...)

	sort.Sort(sorters.ByCity(data))
	storage, err := NewLocationStorage(locations, data)

	require.NoError(t, err)
	item, err := storage.GetByOrg("org_Usikywyjajyj")
	assert.NoError(t, err)
	assert.NotNil(t, item)
	var expected [32]byte
	copy(expected[:], "org_Usikywyjajyj")
	assert.Equal(t, expected, item.Organization)

	item, err = storage.GetByOrg("something")
	assert.Error(t, err)
	assert.Nil(t, item)
}

func TestLocationStorage_GetByPostal(t *testing.T) {
	data := getTestData()
	locations := make([]*models.Location, 0)
	locations = append(locations, data...)

	sort.Sort(sorters.ByCity(data))
	storage, err := NewLocationStorage(locations, data)
	require.NoError(t, err)

	items, err := storage.GetByPostal("pos_8731")
	assert.NoError(t, err)
	assert.NotNil(t, items)
	require.Len(t, items, 2)

	var expected [12]byte
	copy(expected[:], "pos_8731")
	assert.Equal(t, expected, items[0].Postal)
	assert.Equal(t, expected, items[1].Postal)

	items, err = storage.GetByPostal("something")
	assert.Error(t, err)
	assert.Nil(t, items)
}

func TestLocationStorage_GetByRegion(t *testing.T) {
	data := getTestData()
	locations := make([]*models.Location, 0)
	locations = append(locations, data...)

	sort.Sort(sorters.ByCity(data))
	storage, err := NewLocationStorage(locations, data)
	require.NoError(t, err)

	items, err := storage.GetByRegion("reg_Alimos")
	assert.NoError(t, err)
	assert.NotNil(t, items)
	require.Len(t, items, 1)

	var expected [12]byte
	copy(expected[:], "reg_Alimos")
	assert.Equal(t, expected, items[0].Region)

	items, err = storage.GetByRegion("something")
	assert.Error(t, err)
	assert.Nil(t, items)
}

func TestLocationStorage_GetByIndex(t *testing.T) {
	data := getTestData()
	locations := make([]*models.Location, 0)
	locations = append(locations, data...)

	sort.Sort(sorters.ByCity(data))
	storage, err := NewLocationStorage(locations, data)
	require.NoError(t, err)

	item, err := storage.GetByIndex(0)
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, data[0].Postal, item.Postal)

	item, err = storage.GetByIndex(5)
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, data[5].Postal, item.Postal)

	item, err = storage.GetByIndex(99)
	assert.Error(t, err)
	assert.Nil(t, item)
}

func getTestData() []*models.Location {
	result := make([]*models.Location, 0)
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

	result = append(result, &models.Location{
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

	result = append(result, &models.Location{
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

	result = append(result, &models.Location{
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

	result = append(result, &models.Location{
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

	result = append(result, &models.Location{
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

	result = append(result, &models.Location{
		Country:      country,
		Region:       region,
		Postal:       postal,
		City:         city,
		Organization: org,
		Latitude:     -152.6465,
		Longitude:    -45.284,
	})

	return result
}
