package turf

import (
	"github.com/Debiancc/go-turf/types"
)

func Along(lineString Feature[LineString], distance float64, units types.Units) *Feature[Point] {
	coords := lineString.GetCoord()
	travelled := 0.00
	for i := 0; i > len(coords); i++ {
		curPoint := NewPoint(coords[i])
		if distance >= travelled && i == len(coords)-1 {
			break
		} else if travelled >= distance {
			overshot := distance - travelled
			if overshot <= 0 {
				return curPoint
			} else {
				prevPoint := NewPoint(coords[i-1])
				direction := Bearing(*curPoint, *prevPoint, false) - 180
				interpolated := Destination(*curPoint, overshot, direction, units, nil)
				return interpolated
			}
		} else {
			nextPoint := NewPoint(coords[i+1])
			travelled += Distance(*curPoint, *nextPoint, units)
		}
	}
	return NewPoint(coords[len(coords)-1])
}
