package point

import (
	"github.com/Debiancc/go-turf/helper"
	"math"
)

func Bearing(start Point, end Point, final bool) float64 {
	if final {
		return calcFinal(start, end)
	}
	lat1 := helper.DegreesToRadians(start.Lat)
	lat2 := helper.DegreesToRadians(end.Lat)
	lng1 := helper.DegreesToRadians(start.Lng)
	lng2 := helper.DegreesToRadians(end.Lng)
	a := math.Sin(lng2-lng1) * math.Cos(lat2)
	b := math.Cos(lat1)*math.Sin(lat2) - math.Sin(lat1)*math.Cos(lat2)*math.Cos(lng2-lng1)

	return helper.RadiansToDegrees(math.Atan2(a, b))
}

func calcFinal(start Point, end Point) float64 {
	bear := Bearing(end, start, false)
	return math.Mod(bear+180, 360)
}
