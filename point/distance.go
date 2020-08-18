package point

import (
	"github.com/Debiancc/go-turf/helper"
	"github.com/Debiancc/go-turf/types"
	"math"
)

func Distance(from Point, to Point, units types.Units) float64 {
	dLng := helper.DegreesToRadians(to.Lng - from.Lng)
	dLat := helper.DegreesToRadians(to.Lat - from.Lat)
	lat1 := helper.DegreesToRadians(from.Lat)
	lat2 := helper.DegreesToRadians(to.Lat)

	a := math.Pow(math.Sin(dLat/2), 2) + math.Pow(math.Sin(dLng/2), 2)*math.Cos(lat1)*math.Cos(lat2)

	return helper.RadiansToLength(2*math.Atan2(math.Sqrt(a), math.Sqrt(1-a)), units)
}
