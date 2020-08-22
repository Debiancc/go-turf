package point

import (
	"github.com/Debiancc/go-turf/types"
)

func MidPoint(start Point, end Point) *Point {
	dist := Distance(start, end, types.UnitKilometers)
	heading := Bearing(start, end, false)
	return Destination(start, dist/2, heading, types.UnitKilometers, nil)
}
