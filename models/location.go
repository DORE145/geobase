package models

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
