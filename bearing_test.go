package turf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBearing(t *testing.T) {
	assert := assert.New(t)
	form := NewPoint([2]float64{-75, 45})
	to := NewPoint([2]float64{20, 60})

	r1 := Bearing(*form, *to, false)
	assert.Equal(r1, 37.75495852601733)

	r2 := Bearing(*form, *to, true)
	assert.Equal(r2, 120.01405215181421)
}

func TestRhumbBearing(t *testing.T) {
	assert := assert.New(t)
	form := NewPoint([2]float64{-75, 45})
	to := NewPoint([2]float64{20, 60})

	r1 := RhumbBearing(*form, *to, false)
	assert.Equal(r1, 75.28061364784332)

	r2 := RhumbBearing(*form, *to, true)
	assert.Equal(r2, -104.7193863521567)
}
