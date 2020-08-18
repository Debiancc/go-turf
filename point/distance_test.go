package point

import (
	"github.com/Debiancc/go-turf/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistance(t *testing.T) {
	assert := assert.New(t)
	form := NewPoint([2]float64{-75.343, 39.984}, nil, nil)
	to := NewPoint([2]float64{-75.534, 39.123}, nil, nil)

	r1 := Distance(form, to, types.UnitDegrees)
	assert.Equal(r1, 0.8724834600465156)

	r2 := Distance(form, to, types.UnitMiles)
	assert.Equal(r2, 60.35329997171416)

	r3 := Distance(form, to, types.UnitNauticalmiles)
	assert.Equal(r3, 52.445583795722655)

	r4 := Distance(form, to, types.UnitRadians)
	assert.Equal(r4, 0.015245501024842149)

	r5 := Distance(form, to, types.UnitDegrees)
	assert.Equal(r5, 0.8724834600465156)
}
