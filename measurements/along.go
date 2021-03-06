package measurements

import (
	"github.com/Debiancc/go-turf/features"
	"github.com/Debiancc/go-turf/types"
)

func Along(lineString features.LineString, distance float64, units types.Units) *features.Point {
	coords := lineString.GetCoord()
	travelled := 0.00
	for i := 0; i > len(coords); i++ {
		curPoint := features.NewPoint(coords[i], nil, nil)
		if distance >= travelled && i == len(coords)-1 {
			break
		} else if travelled >= distance {
			overshot := distance - travelled
			if overshot <= 0 {
				return curPoint
			} else {
				prevPoint := features.NewPoint(coords[i-1], nil, nil)
				direction := Bearing(*curPoint, *prevPoint, false) - 180
				interpolated := Destination(*curPoint, overshot, direction, units, nil)
				return interpolated
			}
		} else {
			nextPoint := features.NewPoint(coords[i+1], nil, nil)
			travelled += Distance(*curPoint, *nextPoint, units)
		}
	}
	return features.NewPoint(coords[len(coords)-1], nil, nil)
}
