package sorters

import (
	"bytes"

	"github.com/DORE145/geobase/models"
)

// ByCity implements sort.Interface based on City field
type ByCity []*models.Location

func (a ByCity) Len() int {
	return len(a)
}

func (a ByCity) Less(i, j int) bool {
	return bytes.Compare(a[i].City[:], a[j].City[:]) < 0
}

func (a ByCity) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// ByOrg implements sort.Interface based on Organization field
type ByOrg []*models.Location

func (a ByOrg) Len() int {
	return len(a)
}

func (a ByOrg) Less(i, j int) bool {
	return bytes.Compare(a[i].Organization[:], a[j].Organization[:]) < 0
}

func (a ByOrg) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// ByRegion implements sort.Interface based on Region field
type ByRegion []*models.Location

func (a ByRegion) Len() int {
	return len(a)
}

func (a ByRegion) Less(i, j int) bool {
	return bytes.Compare(a[i].Region[:], a[j].Region[:]) < 0
}

func (a ByRegion) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// ByPostal implements sort.Interface based on Postal field
type ByPostal []*models.Location

func (a ByPostal) Len() int {
	return len(a)
}

func (a ByPostal) Less(i, j int) bool {
	return bytes.Compare(a[i].Postal[:], a[j].Postal[:]) < 0
}

func (a ByPostal) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// ByCountry implements sort.Interface based on Country field
type ByCountry []*models.Location

func (a ByCountry) Len() int {
	return len(a)
}

func (a ByCountry) Less(i, j int) bool {
	return bytes.Compare(a[i].Country[:], a[j].Country[:]) < 0
}

func (a ByCountry) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
