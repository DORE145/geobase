package conversion

import (
	"encoding/binary"
	"errors"
	"net"
)

// IPStringToUint32 is an util method to convert string representation of an Ip address to uint32
func IPStringToUint32(address string) (uint32, error) {
	ip := net.ParseIP(address)
	if ip == nil {
		return 0, errors.New("wrong ipAddr format")
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip), nil
}

// Uint32toIPString s an util method to convert uint32 representation of an Ip address to string
func Uint32toIPString(address uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, address)
	ip := net.IP(ipByte)
	return ip.String()
}
