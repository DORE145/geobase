package inmemory

import (
	"github.com/DORE145/geobase/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIpRangeStorage_GetIPRange(t *testing.T) {
	ipRanges := []*models.IpRange{
		{
			IpFrom:        8,
			IpTo:          55472,
			LocationIndex: 0,
		},
		{
			IpFrom:        55473,
			IpTo:          151737,
			LocationIndex: 1,
		},
		{
			IpFrom:        151738,
			IpTo:          191323,
			LocationIndex: 2,
		},
	}

	storage := NewIpRangeStorage(ipRanges)

	IPRange, err := storage.GetIPRange(184908)
	assert.NoError(t, err)
	assert.Equal(t, ipRanges[2], IPRange)

	IPRange, err = storage.GetIPRange(2)
	assert.Error(t, err)
	assert.Nil(t, IPRange)
}
