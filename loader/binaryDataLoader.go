package loader

import (
	"bytes"
	"encoding/binary"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/DORE145/geobase/models"
	"github.com/DORE145/geobase/utils/conversion"
)

// BinaryDataLoader is a struct that loads data from provided binary file
type BinaryDataLoader struct {
	filePath string
	header   *models.Header
}

// NewBinaryDataLoader creates new BinaryDataLoader with very simple initialization
func NewBinaryDataLoader(path string) (*BinaryDataLoader, error) {
	file, err := os.Open(path)
	if err != nil {
		logrus.Errorf("error while opening file: %s", err)
		return nil, err
	}

	defer file.Close()
	logrus.Debugf("Oppened file %s", path)

	// Reading binary file header to check that file exists, and it has right format
	header := models.Header{}
	data := readBytesAtOffset(file, 0, models.HEADER_SIZE)
	buffer := bytes.NewBuffer(data)
	err = binary.Read(buffer, binary.LittleEndian, &header)
	if err != nil {
		logrus.Errorf("failed to read binary file header with error: %s", err)
	}
	logrus.Info("Binary file header processed")
	logrus.Debugf("Binary file with database %s oppened", string(header.Name[:]))

	return &BinaryDataLoader{
		filePath: path,
		header:   &header,
	}, nil
}

// LoadIPRanges reads binary file and returns ip ranges stored in it
func (loader *BinaryDataLoader) LoadIPRanges() ([]*models.IpRange, error) {
	file, err := os.Open(loader.filePath)
	if err != nil {
		logrus.Errorf("error while opening file: %s", err)
		return nil, err
	}

	defer file.Close()
	logrus.Debugf("Oppened file %s", loader.filePath)

	ranges := make([]*models.IpRange, 0, loader.header.Records)
	for i := 0; i < int(loader.header.Records); i++ {
		// Calculating offset for every read
		offset := int64(loader.header.OffsetRanges) + int64(i*models.IP_RANGE_SIZE)
		IPRange := models.IpRange{}
		data := readBytesAtOffset(file, offset, models.IP_RANGE_SIZE)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.LittleEndian, &IPRange)
		if err != nil {
			logrus.Errorf("reading binary file failed with error: %s", err)
			return nil, err
		}
		logrus.Debugf("Info for IP range %s - %s processed", conversion.Uint32toIPString(IPRange.IpFrom),
			conversion.Uint32toIPString(IPRange.IpTo))
		ranges = append(ranges, &IPRange)
	}

	logrus.Infof("Read %d ip ranges from binary file", len(ranges))
	return ranges, nil
}

func (loader *BinaryDataLoader) LoadLocations() ([]*models.Location, error) {
	file, err := os.Open(loader.filePath)
	if err != nil {
		logrus.Errorf("error while opening file: %s", err)
		return nil, err
	}

	defer file.Close()
	logrus.Debugf("Oppened file %s", loader.filePath)

	locationsList := make([]*models.Location, 0, loader.header.Records)
	for i := 0; i < int(loader.header.Records); i++ {
		// Calculating offset for every read
		offset := int64(loader.header.OffsetLocations) + int64(i*models.LOCATION_SIZE)
		location := models.Location{}
		data := readBytesAtOffset(file, offset, models.LOCATION_SIZE)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.LittleEndian, &location)
		if err != nil {
			logrus.Errorf("Reading binary file failed with error %s", err)
		}
		locationsList = append(locationsList, &location)
		logrus.Debugf("Location for org %s is processed", string(location.Organization[:]))
	}
	logrus.Infof("Read %d locations from binary file", len(locationsList))
	return locationsList, err
}

// LoadLocationsCityIndex reads index of sorted locations from binary file and returns sorted locations list by city
func (loader *BinaryDataLoader) LoadLocationsCityIndex(locationsList []*models.Location) ([]*models.Location, error) {
	file, err := os.Open(loader.filePath)
	if err != nil {
		logrus.Errorf("error while opening file: %s", err)
		return nil, err
	}

	locationIndexes := make([]uint32, 0, loader.header.Records)
	for i := 0; i < int(loader.header.Records); i++ {
		var index uint32
		offset := int64(loader.header.OffsetCities) + int64(i*4)
		data := readBytesAtOffset(file, offset, 4)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.LittleEndian, &index)
		if err != nil {
			logrus.Errorf("reading binary file failed with error: %s", err)
			return nil, err
		}
		locationIndexes = append(locationIndexes, index)
	}
	logrus.Debugf("Read %d indexies records from binary files", len(locationIndexes))

	if len(locationsList) != len(locationIndexes) {
		logrus.Error("Amount of locations doesn't match with records number in locations index")
		return nil, err
	}

	// Creating list of locations sorted by city
	sortedLocations := make([]*models.Location, 0, len(locationsList))
	for i, location := range locationIndexes {
		sortedLocations[i] = locationsList[location]
	}

	return sortedLocations, nil
}

func readBytesAtOffset(file *os.File, offset int64, number int) []byte {
	b := make([]byte, number)

	_, err := file.ReadAt(b, offset)
	if err != nil {
		logrus.Fatal(err)
	}

	return b
}
