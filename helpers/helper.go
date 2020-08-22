package helpers

import (
	"github.com/Debiancc/go-turf/types"
	"math"
)

func DegreesToRadians(degrees float64) float64 {
	radians := math.Mod(degrees, 360)
	return radians * math.Pi / 180
}

func RadiansToDegrees(r float64) float64 {
	degrees := math.Mod(r, 2*math.Pi)
	return degrees * 180 / math.Pi
}

func LengthToRadians(distance float64, units types.Units) float64 {
	switch units {
	case types.UnitDegrees:
		return distance / types.FactorDegrees
	case types.UnitKilometers:
		return distance / types.FactorKilometers
	case types.UnitMiles:
		return distance / types.FactorMiles
	case types.UnitNauticalmiles:
		return distance / types.FactorNauticalmiles
	case types.UnitRadians:
		return distance / types.FactorRadians
	default:
		return 0
	}
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

func ConvertLength(length float64, originUnit types.Units, finalUnit types.Units) float64 {
	if length >= 0 == false {
		return 0
	}
	return RadiansToLength(LengthToRadians(length, originUnit), finalUnit)
}
