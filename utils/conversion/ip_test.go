package conversion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIPStringToUint32(t *testing.T) {
	tests := []struct {
		ip       string
		expected uint32
	}{
		{
			ip:       "41.68.19.196",
			expected: 692327364,
		},
		{
			ip:       "7.51.60.30",
			expected: 120798238,
		},
		{
			ip:       "246.214.101.153",
			expected: 4141245849,
		},
		{
			ip:       "237.247.40.142",
			expected: 3992397966,
		},
		{
			ip:       "48.194.212.84",
			expected: 818074708,
		},
	}

	for _, test := range tests {
		res, err := IPStringToUint32(test.ip)
		assert.NoError(t, err)
		assert.Equal(t, test.expected, res)
	}
}

func TestUint32toIPString(t *testing.T) {
	tests := []struct {
		ip       uint32
		expected string
	}{
		{
			ip:       2596996162,
			expected: "154.203.4.66",
		},
		{
			ip:       4039455774,
			expected: "240.197.52.30",
		},
		{
			ip:       2854263694,
			expected: "170.32.155.142",
		},
		{
			ip:       1879968118,
			expected: "112.14.9.118",
		},
		{
			ip:       1823804162,
			expected: "108.181.11.2",
		},
	}
	for _, test := range tests {
		res := Uint32toIPString(test.ip)
		assert.Equal(t, test.expected, res)
	}
}
