package models

const LOCATION_SIZE = 96

// Location represents the location of specific organizations and cities
// 96 bytes per record in binary file
// sbyte country[8] - country name (starts with "cou_")
// sbyte region[12] - region name (starts with "reg_")
// sbyte postal[12] - postal index (starts with "pos_")
// sbyte city[24] - city name (stars with "cit_")
// sbyte organization[32] - organization name (starts with "org_")
// float latitude - latitude, 4 bytes
// float longitude - longitude, 4 bytes
type Location struct {
	Country      [8]byte
	Region       [12]byte
	Postal       [12]byte
	City         [24]byte
	Organization [32]byte
	Latitude     float32
	Longitude    float32
}

type LocationResp struct {
	Country      string  `json:"country"`
	Region       string  `json:"region"`
	Postal       string  `json:"postal"`
	City         string  `json:"city"`
	Organization string  `json:"organization"`
	Latitude     float32 `json:"latitude"`
	Longitude    float32 `json:"longitude"`
}

func (location Location) ToResponse() LocationResp {
	return LocationResp{
		Country:      string(location.Country[:]),
		Region:       string(location.Region[:]),
		Postal:       string(location.Postal[:]),
		City:         string(location.City[:]),
		Organization: string(location.Organization[:]),
		Latitude:     location.Latitude,
		Longitude:    location.Longitude,
	}
}
