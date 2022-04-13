package models

const HEADER_SIZE = 60

// Header describes the contents of the source data
// int   version;
// sbyte name[32];
// ulong timestamp;
// int   records;
// uint  offset_ranges;
// uint  offset_cities;
// uint  offset_locations;
type Header struct {
	Version         int32
	Name            [32]byte
	Timestamp       uint64
	Records         int32
	OffsetRanges    uint32
	OffsetCities    uint32
	OffsetLocations uint32
}
