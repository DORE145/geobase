package storage

import "github.com/DORE145/geobase/models"

type IpRangeStorage interface {
	GetIPRange(uint32) (*models.IpRange, error)
}
