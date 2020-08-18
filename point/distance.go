package point

import (
	"github.com/Debiancc/go-turf/types"
	"math"
)

func Distance(form Point, to Point, units types.Units) float64 {
	dLng := DegreesToRadians(to.Lng - form.Lng)
	dLat := DegreesToRadians(to.Lat - form.Lat)
	lat1 := DegreesToRadians(form.Lat)
	lat2 := DegreesToRadians(to.Lat)

	a := math.Pow(math.Sin(dLat/2), 2) + math.Pow(math.Sin(dLng/2), 2)*math.Cos(lat1)*math.Cos(lat2)

	return RadiansToLength(2*math.Atan2(math.Sqrt(a), math.Sqrt(1-a)), units)
}

func DegreesToRadians(degrees float64) float64 {
	radians := math.Mod(degrees, 360)
	return radians * math.Pi / 180
}

func RadiansToLength(radians float64, units types.Units) float64 {
	switch units {
	case types.UnitDegrees:
		return radians * types.FactorDegrees
	case types.UnitKilometers:
		return radians * types.FactorKilometers
	case types.UnitMiles:
		return radians * types.FactorMiles
	case types.UnitNauticalmiles:
		return radians * types.FactorNauticalmiles
	case types.UnitRadians:
		return radians * types.FactorRadians
	default:
		return 0
	}
}
