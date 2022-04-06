package models

// IpRange represents the mapping of Ip addresses to the respected location
// 12 bytes in binary file
// 4 bytes uint ip_from
// 4 bytes uint ip_to
// 4 byte uint location_index
type IpRange struct {
	IpFrom        uint32
	IpTo          uint32
	LocationIndex uint32
}
