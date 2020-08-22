package measurements

import (
	"github.com/Debiancc/go-turf/features"
	"github.com/Debiancc/go-turf/helpers"
	"math"
)

func Bearing(start features.Point, end features.Point, final bool) float64 {
	if final {
		return calcFinal(start, end)
	}
	lat1 := helpers.DegreesToRadians(start.GetLat())
	lat2 := helpers.DegreesToRadians(end.GetLat())
	lng1 := helpers.DegreesToRadians(start.GetLng())
	lng2 := helpers.DegreesToRadians(end.GetLng())
	a := math.Sin(lng2-lng1) * math.Cos(lat2)
	b := math.Cos(lat1)*math.Sin(lat2) - math.Sin(lat1)*math.Cos(lat2)*math.Cos(lng2-lng1)

	return helpers.RadiansToDegrees(math.Atan2(a, b))
}

func calcFinal(start features.Point, end features.Point) float64 {
	bear := Bearing(end, start, false)
	return math.Mod(bear+180, 360)
}

func RhumbBearing(start features.Point, end features.Point, final bool) float64 {
	var bear360 float64
	if final {
		bear360 = calcRhumbBearing(end, start)
	} else {
		bear360 = calcRhumbBearing(start, end)
	}

	bear180 := bear360
	if bear360 > 180 {
		bear180 = -(360 - bear360)
	}

	return bear180
}

func calcRhumbBearing(start features.Point, end features.Point) float64 {
	phi1 := helpers.DegreesToRadians(start.GetLat())
	phi2 := helpers.DegreesToRadians(end.GetLat())
	deltaLambda := helpers.DegreesToRadians(end.GetLng() - start.GetLng())
	if deltaLambda > math.Pi {
		deltaLambda -= 2 * math.Pi
	}
	if deltaLambda < -math.Pi {
		deltaLambda += 2 * math.Pi
	}

	deltaPsi := math.Log(math.Tan(phi2/2+math.Pi/4) / math.Tan(phi1/2+math.Pi/4))
	theta := math.Atan2(deltaLambda, deltaPsi)

	return math.Mod(helpers.RadiansToDegrees(theta)+360, 360)
}
