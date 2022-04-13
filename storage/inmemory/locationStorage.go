package inmemory

import (
	"bytes"
	"errors"
	"sort"

	"github.com/sirupsen/logrus"

	"github.com/DORE145/geobase/models"
	"github.com/DORE145/geobase/utils/sorters"
)

// LocationStorage is an inmemory storage for Locations with several indexes
type LocationStorage struct {
	locationOffsets map[int]*models.Location
	cityIndex       []*models.Location
	countryIndex    []*models.Location
	orgIndex        []*models.Location
	postalIndex     []*models.Location
	regionIndex     []*models.Location
}

// NewLocationStorage constructs new inmemory LocationStorage and indexes for different fields
func NewLocationStorage(locationOffsets map[int]*models.Location, cityIndex []*models.Location) (*LocationStorage, error) {
	// copying references from already sorted cityIndex and creating indexes for all other fields in Location
	// by sorting them
	countryIndex := make([]*models.Location, len(cityIndex))
	count := copy(countryIndex, cityIndex)
	if count != len(cityIndex) {
		logrus.Debugf("failed to create countries index")
		return nil, errors.New("failed to build location storage")
	}
	sort.Sort(sorters.ByCountry(countryIndex))

	orgIndex := make([]*models.Location, len(cityIndex))
	count = copy(orgIndex, cityIndex)
	if count != len(cityIndex) {
		logrus.Debugf("failed to create organizations index")
		return nil, errors.New("failed to build location storage")
	}
	sort.Sort(sorters.ByOrg(orgIndex))

	postalIndex := make([]*models.Location, len(cityIndex))
	count = copy(postalIndex, cityIndex)
	if count != len(cityIndex) {
		logrus.Debugf("failed to create postal codes index")
		return nil, errors.New("failed to build location storage")
	}
	sort.Sort(sorters.ByPostal(postalIndex))

	regionIndex := make([]*models.Location, len(cityIndex))
	count = copy(regionIndex, cityIndex)
	if count != len(cityIndex) {
		logrus.Debugf("failed to create regions index")
		return nil, errors.New("failed to build location storage")
	}
	sort.Sort(sorters.ByRegion(regionIndex))
	return &LocationStorage{
		locationOffsets: locationOffsets,
		cityIndex:       cityIndex,
		countryIndex:    countryIndex,
		orgIndex:        orgIndex,
		postalIndex:     postalIndex,
		regionIndex:     regionIndex,
	}, nil
}

// GetByOrg return a Location of requested Organization
func (storage *LocationStorage) GetByOrg(org string) (*models.Location, error) {
	// converting string to bytes array by copying string to temp wrapper slice created in orgBytes[:] command
	var orgBytes [32]byte
	copy(orgBytes[:], org)

	// Binary search in already sorted by name organizations index
	low := 0
	high := len(storage.orgIndex) - 1
	for low <= high {
		mid := (low + high) / 2
		record := storage.orgIndex[mid]
		comparison := bytes.Compare(orgBytes[:], record.Organization[:])
		if comparison > 0 {
			low = mid + 1
		} else if comparison < 0 {
			high = mid - 1
		} else {
			return record, nil
		}
	}
	logrus.Debugf("Location for organiazation: %s not found", org)
	return nil, errors.New("location not found")
}

// GetByCity returns all locations in the same city
func (storage *LocationStorage) GetByCity(city string) ([]*models.Location, error) {
	// in theory there could be multiple locations in one city
	result := make([]*models.Location, 0)

	// converting string to bytes array by copying string to temp wrapper slice created in cityBytes[:] command
	var cityBytes [24]byte
	copy(cityBytes[:], city)

	// Binary search in already sorted by name organizations index
	low := 0
	high := len(storage.cityIndex) - 1
	mid := 0
	for low <= high {
		mid = (low + high) / 2
		record := storage.cityIndex[mid]
		comparison := bytes.Compare(cityBytes[:], record.City[:])
		if comparison > 0 {
			low = mid + 1
		} else if comparison < 0 {
			high = mid - 1
		} else {
			result = append(result, record)
			break
		}
	}

	if len(result) == 0 {
		logrus.Debugf("Locations in the city %s not found", city)
		return nil, errors.New("locations not found")
	}

	// Checking neighbours of found record for matching city
	for i := mid - 1; i >= 0; i-- {
		record := storage.cityIndex[i]
		if bytes.Equal(record.City[:], cityBytes[:]) {
			result = append(result, record)
		} else {
			break
		}
	}
	for i := mid + 1; i < len(storage.cityIndex); i++ {
		record := storage.cityIndex[i]
		if bytes.Equal(record.City[:], cityBytes[:]) {
			result = append(result, record)
		} else {
			break
		}
	}

	logrus.Debugf("For city %s found %d location(s)", city, len(result))
	return result, nil
}

// GetByPostal returns all locations with the same postal code
func (storage *LocationStorage) GetByPostal(postal string) ([]*models.Location, error) {
	// in theory there could be multiple locations with same postal code
	result := make([]*models.Location, 0)

	// converting string to bytes array by copying string to temp wrapper slice created in postalBytes[:] command
	var postalBytes [12]byte
	copy(postalBytes[:], postal)

	// Binary search in already sorted by name organizations index
	low := 0
	high := len(storage.postalIndex) - 1
	mid := 0
	for low <= high {
		mid = (low + high) / 2
		record := storage.postalIndex[mid]
		comparison := bytes.Compare(postalBytes[:], record.Postal[:])
		if comparison > 0 {
			low = mid + 1
		} else if comparison < 0 {
			high = mid - 1
		} else {
			result = append(result, record)
			break
		}
	}

	if len(result) == 0 {
		logrus.Debugf("Locations in the post code %s not found", postal)
		return nil, errors.New("locations not found")
	}

	// Checking neighbours of found record for matching city
	for i := mid - 1; i >= 0; i-- {
		record := storage.cityIndex[i]
		if bytes.Equal(record.Postal[:], postalBytes[:]) {
			result = append(result, record)
		} else {
			break
		}
	}
	for i := mid + 1; i < len(storage.cityIndex); i++ {
		record := storage.cityIndex[i]
		if bytes.Equal(record.Postal[:], postalBytes[:]) {
			result = append(result, record)
		} else {
			break
		}
	}

	logrus.Debugf("For postal code %s found %d location(s)", postal, len(result))
	return result, nil
}

// GetByRegion returns all locations in the same region
func (storage *LocationStorage) GetByRegion(region string) ([]*models.Location, error) {
	// in theory there could be multiple locations in one region
	result := make([]*models.Location, 0)

	// converting string to bytes array by copying string to temp wrapper slice created in regionBytes[:] command
	var regionBytes [12]byte
	copy(regionBytes[:], region)

	// Binary search in already sorted by name region index
	low := 0
	high := len(storage.regionIndex) - 1
	mid := 0
	for low <= high {
		mid = (low + high) / 2
		record := storage.regionIndex[mid]
		comparison := bytes.Compare(regionBytes[:], record.Region[:])
		if comparison > 0 {
			low = mid + 1
		} else if comparison < 0 {
			high = mid - 1
		} else {
			result = append(result, record)
			break
		}
	}

	if len(result) == 0 {
		logrus.Debugf("Locations in the region %s not found", region)
		return nil, errors.New("locations not found")
	}

	// Checking neighbours of found record for matching city
	for i := mid - 1; i >= 0; i-- {
		record := storage.regionIndex[i]
		if bytes.Equal(record.Region[:], regionBytes[:]) {
			result = append(result, record)
		} else {
			break
		}
	}
	for i := mid + 1; i < len(storage.regionIndex); i++ {
		record := storage.regionIndex[i]
		if bytes.Equal(record.Region[:], regionBytes[:]) {
			result = append(result, record)
		} else {
			break
		}
	}

	logrus.Debugf("In the region %s found %d location(s)", region, len(result))
	return result, nil
}

// GetByCountry returns all locations in the same country
func (storage *LocationStorage) GetByCountry(country string) ([]*models.Location, error) {
	// in theory there could be multiple locations in one city
	result := make([]*models.Location, 0)

	// converting string to bytes array by copying string to temp wrapper slice created in countryBytes[:] command
	var countryBytes [8]byte
	copy(countryBytes[:], country)

	// Binary search in already sorted by name organizations index
	low := 0
	high := len(storage.countryIndex) - 1
	mid := 0
	for low <= high {
		mid = (low + high) / 2
		record := storage.countryIndex[mid]
		comparison := bytes.Compare(countryBytes[:], record.Country[:])
		if comparison > 0 {
			low = mid + 1
		} else if comparison < 0 {
			high = mid - 1
		} else {
			result = append(result, record)
			break
		}
	}

	if len(result) == 0 {
		logrus.Debugf("Locations in the country %s not found", country)
		return nil, errors.New("locations not found")
	}

	// Checking neighbours of found record for matching city
	for i := mid - 1; i >= 0; i-- {
		record := storage.cityIndex[i]
		if bytes.Equal(record.Postal[:], countryBytes[:]) {
			result = append(result, record)
		} else {
			break
		}
	}
	for i := mid + 1; i < len(storage.cityIndex); i++ {
		record := storage.cityIndex[i]
		if bytes.Equal(record.Postal[:], countryBytes[:]) {
			result = append(result, record)
		} else {
			break
		}
	}

	logrus.Debugf("In the country %s found %d location(s)", country, len(result))
	return result, nil
}

func (storage *LocationStorage) GetByIndex(index int) (*models.Location, error) {
	// Multiplying index by models.LOCATION_SIZE (96) because it is the size of location record, and they stored by the offset
	location, ok := storage.locationOffsets[index*models.LOCATION_SIZE]
	if !ok {
		logrus.Debugf("Location with offset %d not found", index*models.LOCATION_SIZE)
		return nil, errors.New("location not found")
	}
	return location, nil
}
