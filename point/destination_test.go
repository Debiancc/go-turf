package point

import (
	"github.com/Debiancc/go-turf/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDestination(t *testing.T) {
	assert := assert.New(t)

	p1 := Destination(NewPoint([2]float64{-75.00, 38.10096062273525}, nil, nil), 100, 0, types.UnitKilometers, nil)
	assert.Equal(p1.Lat, 39.00028098645979)
	assert.Equal(p1.Lng, -75.00000000000001)

	p2 := Destination(NewPoint([2]float64{-75.00, 39}, nil, nil), 100, 90, types.UnitKilometers, nil)
	assert.Equal(p2.Lat, 38.99428496385679)
	assert.Equal(p2.Lng, -73.84285322793312)

	p3 := Destination(NewPoint([2]float64{-75.00, 39}, nil, nil), 100, 180, types.UnitKilometers, nil)
	assert.Equal(p3.Lat, 38.10067963627546)
	assert.Equal(p3.Lng, -75.00000000000001)

	p4 := Destination(NewPoint([2]float64{-75, 39}, nil, nil), 5000, 90, types.UnitKilometers, nil)
	assert.Equal(p4.Lng, -22.88535554936431)
	assert.Equal(p4.Lat, 26.440010707631124)

}
