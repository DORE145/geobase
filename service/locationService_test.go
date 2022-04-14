package service

import (
	"sort"
	"testing"

	"github.com/DORE145/geobase/models"
	"github.com/DORE145/geobase/storage"
	"github.com/DORE145/geobase/storage/inmemory"
	"github.com/DORE145/geobase/utils/sorters"
)

func TestNewLocationService(t *testing.T) {

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
	for _, item := range locations {
		locationsIndex = append(locationsIndex, item)
	}
	sort.Sort(sorters.ByCity(locationsIndex))

	return inmemory.NewLocationStorage(locations, locationsIndex)

}
