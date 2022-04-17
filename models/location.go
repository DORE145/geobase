package models

import "bytes"

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

type Locations []*Location

type LocationResp struct {
	Country      string  `json:"country"`
	Region       string  `json:"region"`
	Postal       string  `json:"postal"`
	City         string  `json:"city"`
	Organization string  `json:"organization"`
	Latitude     float32 `json:"latitude"`
	Longitude    float32 `json:"longitude"`
}

type LocationsResp []*LocationResp

func (location Location) ToResponse() LocationResp {
	return LocationResp{
		Country:      string(bytes.Trim(location.Country[:], "\u0000")),
		Region:       string(bytes.Trim(location.Region[:], "\u0000")),
		Postal:       string(bytes.Trim(location.Postal[:], "\u0000")),
		City:         string(bytes.Trim(location.City[:], "\u0000")),
		Organization: string(bytes.Trim(location.Organization[:], "\u0000")),
		Latitude:     location.Latitude,
		Longitude:    location.Longitude,
	}
}

func (locations Locations) ToResponse() LocationsResp {
	locationsResp := make([]*LocationResp, 0, len(locations))
	for _, loc := range locations {
		locationResp := loc.ToResponse()
		locationsResp = append(locationsResp, &locationResp)
	}
	return locationsResp
}
