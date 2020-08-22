package measurements

import (
	"github.com/Debiancc/go-turf/features"
	"github.com/Debiancc/go-turf/types"
)

func MidPoint(start features.Point, end features.Point) *features.Point {
	dist := Distance(start, end, types.UnitKilometers)
	heading := Bearing(start, end, false)
	return Destination(start, dist/2, heading, types.UnitKilometers, nil)
}
