package point

import (
	"github.com/Debiancc/go-turf/helper"
	"github.com/Debiancc/go-turf/types"
	"math"
)

func Distance(from Point, to Point, units types.Units) float64 {
	dLng := helper.DegreesToRadians(to.GetLng() - from.GetLng())
	dLat := helper.DegreesToRadians(to.GetLat() - from.GetLat())
	lat1 := helper.DegreesToRadians(from.GetLat())
	lat2 := helper.DegreesToRadians(to.GetLat())

	a := math.Pow(math.Sin(dLat/2), 2) + math.Pow(math.Sin(dLng/2), 2)*math.Cos(lat1)*math.Cos(lat2)

	return helper.RadiansToLength(2*math.Atan2(math.Sqrt(a), math.Sqrt(1-a)), units)
}
