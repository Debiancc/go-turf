package turf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFeature(t *testing.T) {
	assert := assert.New(t)
	p := NewFeature[Point]([2]float64{1.2222, 1.1111})
	assert.Equal(p.Geometry.Type, "Point", "geometry type is correct")
	assert.Equal(p.Type, "Feature", "geometry type is correct")
}
