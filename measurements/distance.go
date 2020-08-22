package measurements

import (
	"github.com/Debiancc/go-turf/features"
	"github.com/Debiancc/go-turf/helpers"
	"github.com/Debiancc/go-turf/types"
	"math"
)

func Distance(from features.Point, to features.Point, units types.Units) float64 {
	dLng := helpers.DegreesToRadians(to.GetLng() - from.GetLng())
	dLat := helpers.DegreesToRadians(to.GetLat() - from.GetLat())
	lat1 := helpers.DegreesToRadians(from.GetLat())
	lat2 := helpers.DegreesToRadians(to.GetLat())

	a := math.Pow(math.Sin(dLat/2), 2) + math.Pow(math.Sin(dLng/2), 2)*math.Cos(lat1)*math.Cos(lat2)

	return helpers.RadiansToLength(2*math.Atan2(math.Sqrt(a), math.Sqrt(1-a)), units)
}
