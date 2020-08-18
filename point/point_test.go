package point

import "testing"
import "github.com/stretchr/testify/assert"

func TestNewPoint(t *testing.T) {
	assert := assert.New(t)
	p := NewPoint(
		[2]float64{-75.343, 39.984},
		&Properties{"city": "ShangHai", "osm_id": 123},
		nil,
	)
	assert.Equal(p.Lng, -75.343, "New point lng are incorrect")
	assert.Equal(p.Lat, 39.984, "New point lat are incorrect")
	assert.Equal(p.Properties["city"], "ShangHai", "New point prop are incorrect")
	assert.Equal(p.Properties["osm_id"], 123, "New point prop are incorrect")
}
