package turf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFeatureCollection(t *testing.T) {
	assert := assert.New(t)

	pointFeature := Feature[Point]{
		Type: "Point",
		Geometry: &Geometry[Point]{
			Type:        "Point",
			Coordinates: Point{1, 2},
		},
	}

	lineFeature := Feature[LineString]{
		Type: "LineString",
		Geometry: &Geometry[LineString]{
			Type:        "LineString",
			Coordinates: [][2]float64{{1, 2}, {3, 4}},
		},
	}

	list := NewFeatureCollection(FeatureList{
		pointFeature,
		lineFeature,
	})

	assert.Equal(len(list.Features), 2)
	assert.Equal(list.Features[0].GetGeometryType(), "Point")
	assert.Equal(list.Features[1].GetGeometryType(), "LineString")

}
