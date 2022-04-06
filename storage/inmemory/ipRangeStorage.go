package inmemory

import (
	"errors"

	"github.com/sirupsen/logrus"

	"github.com/DORE145/geobase/models"
	"github.com/DORE145/geobase/utils/conversion"
)

type IpRangeStorage struct {
	records []*models.IpRange
}

func NewIpRangeStorage(records []*models.IpRange) *IpRangeStorage {
	return &IpRangeStorage{
		records: records,
	}
}

func (storage *IpRangeStorage) GetIPRange(address uint32) (*models.IpRange, error) {
	// Binary search implementation, not as graceful as standard library
	low := 0
	high := len(storage.records)
	for low <= high {
		mid := low/2 + high/2 //Avoiding potential overflow here
		record := storage.records[mid]
		if address > record.IpTo {
			low = mid + 1
		} else if address < record.IpFrom {
			high = mid - 1
		} else {
			return record, nil
		}
	}

	// Initial implementation based on internal sort.Search
	// This implementation is easier to write and use, but it is not obvious how it works
	// and what is going without a documentation dive
	// sort.Search searches for the lowest index where the passed function is true
	//index := sort.Search(len(storage.records), func(i int) bool {
	//	return address >= storage.records[i].IpFrom
	//})
	//if index < len(storage.records){
	//	return storage.records[index], nil
	//}

	logrus.Debugf("ip record for address %s not found", conversion.Uint32toIPString(address))
	return nil, errors.New("ip record not found")
}
