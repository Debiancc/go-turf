package features

import (
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestNewPoint(t *testing.T) {
	assert := assert.New(t)
	p := NewPoint(
		[2]float64{-75.343, 39.984},
		&Properties{"city": "ShangHai", "osm_id": 123},
		nil,
	)
	lng := p.GetLng()
	lat := p.GetLat()
	assert.Equal(lng, -75.343, "New point lng are incorrect")
	assert.Equal(lat, 39.984, "New point lat are incorrect")
	assert.Equal(p.Properties["city"], "ShangHai", "New point prop are incorrect")
	assert.Equal(p.Properties["osm_id"], 123, "New point prop are incorrect")
}
