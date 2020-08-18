package point

import (
	"github.com/Debiancc/go-turf/helper"
	"github.com/Debiancc/go-turf/types"
	"math"
)

func Destination(origin Point, distance float64, bearing float64, units types.Units, properties *Properties) Point {
	lng1 := helper.DegreesToRadians(origin.Lng)
	lat1 := helper.DegreesToRadians(origin.Lat)
	bearingRad := helper.DegreesToRadians(bearing)
	radians := helper.LengthToRadians(distance, units)

	lat2 := math.Asin(math.Sin(lat1)*math.Cos(radians) + math.Cos(lat1)*math.Sin(radians)*math.Cos(bearingRad))
	lng2 := lng1 + math.Atan2(
		math.Sin(bearingRad)*math.Sin(radians)*math.Cos(lat1),
		math.Cos(radians)-math.Sin(lat1)*math.Sin(lat2),
	)
	lng := helper.RadiansToDegrees(lng2)
	lat := helper.RadiansToDegrees(lat2)

	return NewPoint([2]float64{lng, lat}, properties, nil)
}
