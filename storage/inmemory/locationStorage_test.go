package inmemory

import (
	"bytes"
	"github.com/DORE145/geobase/models"
	"github.com/DORE145/geobase/utils/sorters"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestNewLocationStorage(t *testing.T) {
	data := getTestData()
	locationsMaps := make(map[int]*models.Location)
	for i, item := range data {
		locationsMaps[i*96] = item
	}
	sort.Sort(sorters.ByCity(data))
	storage, err := NewLocationStorage(locationsMaps, data)
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
}

func TestLocationStorage_GetByCountry(t *testing.T) {

}

func TestLocationStorage_GetByOrg(t *testing.T) {

}

func TestLocationStorage_GetByPostal(t *testing.T) {

}

func TestLocationStorage_GetByRegion(t *testing.T) {

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

	return result
}
