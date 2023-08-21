package measurements

import (
	"github.com/Debiancc/go-turf/features"
	"github.com/Debiancc/go-turf/helpers"
	"github.com/Debiancc/go-turf/types"
	"math"
)

func Distance(from features.Feature[features.Point], to features.Feature[features.Point], units types.Units) float64 {
	dLng := helpers.DegreesToRadians(to.Geometry.Coordinates[0] - from.Geometry.Coordinates[0])
	dLat := helpers.DegreesToRadians(to.Geometry.Coordinates[1] - from.Geometry.Coordinates[1])
	lat1 := helpers.DegreesToRadians(from.Geometry.Coordinates[1])
	lat2 := helpers.DegreesToRadians(to.Geometry.Coordinates[1])

	a := math.Pow(math.Sin(dLat/2), 2) + math.Pow(math.Sin(dLng/2), 2)*math.Cos(lat1)*math.Cos(lat2)

	return helpers.RadiansToLength(2*math.Atan2(math.Sqrt(a), math.Sqrt(1-a)), units)
}
