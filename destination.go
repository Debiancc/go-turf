package turf

import (
	"github.com/Debiancc/go-turf/helpers"
	"github.com/Debiancc/go-turf/types"
	"math"
)

func Destination(origin Feature[Point], distance float64, bearing float64, units types.Units, properties *Properties) *Feature[Point] {
	lng1 := helpers.DegreesToRadians(origin.Geometry.Coordinates[0])
	lat1 := helpers.DegreesToRadians(origin.Geometry.Coordinates[1])
	bearingRad := helpers.DegreesToRadians(bearing)
	radians := helpers.LengthToRadians(distance, units)

	lat2 := math.Asin(math.Sin(lat1)*math.Cos(radians) + math.Cos(lat1)*math.Sin(radians)*math.Cos(bearingRad))
	lng2 := lng1 + math.Atan2(
		math.Sin(bearingRad)*math.Sin(radians)*math.Cos(lat1),
		math.Cos(radians)-math.Sin(lat1)*math.Sin(lat2),
	)
	lng := helpers.RadiansToDegrees(lng2)
	lat := helpers.RadiansToDegrees(lat2)

	//return NewPoint([2]float64{lng, lat}, properties, nil)
	return NewPoint([2]float64{lng, lat})
}

func RhumbDestination(origin Feature[Point], distance float64, bearing float64, units types.Units, properties *Properties) *Feature[Point] {
	wasNegativeDostamce := distance < 0
	distanceInMeters := helpers.ConvertLength(math.Abs(distance), units, types.UnitMeters)
	if wasNegativeDostamce {
		distanceInMeters = -math.Abs(distanceInMeters)
	}
	destination := calcRhumbDestination(origin, distanceInMeters, bearing, nil)
	if destination.Geometry.Coordinates[0]-origin.Geometry.Coordinates[0] > 180 {
		destination.Geometry.Coordinates[0] += -360
	} else {
		if origin.Geometry.Coordinates[0]-destination.Geometry.Coordinates[0] > 180 {
			destination.Geometry.Coordinates[0] += 360
		}
	}

	//return NewPoint(destination.Geometry.Coordinates, properties, nil)
	return NewPoint(destination.Geometry.Coordinates)
}

func calcRhumbDestination(origin Feature[Point], distance float64, bearing float64, radius *float64) *Feature[Point] {
	var delta float64
	if radius == nil {
		delta = distance / types.FactorEarthRadius
	} else {
		delta = distance / *radius
	}

	lambda1 := origin.Geometry.Coordinates[0] * math.Pi / 180
	phi1 := helpers.DegreesToRadians(origin.Geometry.Coordinates[1])
	theta := helpers.DegreesToRadians(bearing)

	deltaPhi := delta * math.Cos(theta)
	phi2 := phi1 + deltaPhi
	if math.Abs(phi2) > math.Pi/2 {
		if phi2 > 0 {
			phi2 = math.Pi - phi2
		} else {
			phi2 = -math.Pi - phi2
		}
	}
	deltaPsi := math.Log(math.Tan(phi2/2+math.Pi/4) / math.Tan(phi1/2+math.Pi/4))
	var q float64
	if math.Abs(deltaPsi) > 10e-12 {
		q = deltaPhi / deltaPsi
	} else {
		q = math.Cos(phi1)
	}

	deltaLambda := delta * math.Sin(theta) / q
	lambda2 := lambda1 + deltaLambda

	lng := math.Mod((lambda2*180/math.Pi)+540, 360) - 180
	lat := phi2 * 180 / math.Pi
	//return NewPoint([2]float64{lng, lat}, nil, nil)
	return NewPoint([2]float64{lng, lat})
}
