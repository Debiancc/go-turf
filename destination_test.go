package turf

import (
	"github.com/Debiancc/go-turf/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDestination(t *testing.T) {
	assert := assert.New(t)

	p1 := Destination(*NewPoint([2]float64{-75.00, 38.10096062273525}), 100, 0, types.UnitKilometers, nil)
	assert.Equal(p1.Geometry.Coordinates[1], 39.00028098645979)
	assert.Equal(p1.Geometry.Coordinates[0], -75.00000000000001)

	p2 := Destination(*NewPoint([2]float64{-75.00, 39}), 100, 90, types.UnitKilometers, nil)
	assert.Equal(p2.Geometry.Coordinates[1], 38.99428496385679)
	assert.Equal(p2.Geometry.Coordinates[0], -73.84285322793312)

	p3 := Destination(*NewPoint([2]float64{-75.00, 39}), 100, 180, types.UnitKilometers, nil)
	assert.Equal(p3.Geometry.Coordinates[1], 38.10067963627546)
	assert.Equal(p3.Geometry.Coordinates[0], -75.00000000000001)

	p4 := Destination(*NewPoint([2]float64{-75, 39}), 5000, 90, types.UnitKilometers, nil)
	assert.Equal(p4.Geometry.Coordinates[0], -22.88535554936431)
	assert.Equal(p4.Geometry.Coordinates[1], 26.440010707631124)

}

//func TestRhumbDestination(t *testing.T) {
//	assert := assert.New(t)
//
//	// TODO fix
//	p1 := RhumbDestination(NewPoint([2]float64{-179.5, -16.5}, nil, nil), 100, -90, types.UnitKilometers, nil)
//	assert.Equal(p1.Lng, -180.437945)
//	assert.Equal(p1.Lat, -16.5)
//}
