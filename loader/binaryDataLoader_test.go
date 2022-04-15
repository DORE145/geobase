package loader

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewBinaryDataLoader(t *testing.T) {
	filePath := "testdata/test.dat"
	loader, err := NewBinaryDataLoader(filePath)
	require.NoError(t, err)
	require.NotNil(t, loader)

	assert.Equal(t, int32(5), loader.header.Records)
	assert.Equal(t, int32(1), loader.header.Version)
	assert.Equal(t, uint32(600), loader.header.OffsetCities)
}

func TestBinaryDataLoader_LoadIPRanges(t *testing.T) {
	filePath := "testdata/test.dat"
	loader, err := NewBinaryDataLoader(filePath)
	require.NoError(t, err)
	require.NotNil(t, loader)

	ranges, err := loader.LoadIPRanges()
	require.NoError(t, err)
	require.NotNil(t, ranges)

	assert.Len(t, ranges, 5)
	assert.Equal(t, uint32(191323), ranges[2].IpTo)
	assert.Equal(t, uint32(4), ranges[4].LocationIndex)
	assert.Equal(t, uint32(55473), ranges[1].IpFrom)
}

func TestBinaryDataLoader_LoadLocations(t *testing.T) {
	filePath := "testdata/test.dat"
	loader, err := NewBinaryDataLoader(filePath)
	require.NoError(t, err)
	require.NotNil(t, loader)

	locations, err := loader.LoadLocations()
	require.NoError(t, err)
	require.NotNil(t, locations)

	assert.Len(t, locations, 5)
	assert.Equal(t, float32(-96.2552), locations[0].Latitude)
	assert.Equal(t, "reg_Alimos\u0000\u0000", string(locations[3].Region[:]))
	assert.Equal(t, "cou_AK\u0000\u0000", string(locations[2].Country[:]))
	assert.Equal(t, float32(-45.284), locations[4].Longitude)
	assert.Equal(t, "pos_96188\u0000\u0000\u0000", string(locations[1].Postal[:]))
}

func TestBinaryDataLoader_LoadLocationsCityIndex(t *testing.T) {
	filePath := "testdata/test.dat"
	loader, err := NewBinaryDataLoader(filePath)
	require.NoError(t, err)
	require.NotNil(t, loader)

	locations, err := loader.LoadLocations()
	require.NoError(t, err)
	require.NotNil(t, locations)

	locationIndex, err := loader.LoadLocationsCityIndex(locations)
	require.NoError(t, err)
	require.NotNil(t, locationIndex)

	assert.Equal(t, float32(-96.2552), locationIndex[4].Latitude)
	assert.Equal(t, "reg_Alimos\u0000\u0000", string(locationIndex[1].Region[:]))
	assert.Equal(t, "cou_AK\u0000\u0000", string(locationIndex[2].Country[:]))
	assert.Equal(t, float32(-45.284), locationIndex[0].Longitude)
	assert.Equal(t, "pos_96188\u0000\u0000\u0000", string(locationIndex[3].Postal[:]))
}
