package helpers

import (
	"github.com/Debiancc/go-turf/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBearing(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(LengthToRadians(1, types.UnitRadians), 1.00)
	assert.Equal(LengthToRadians(types.FactorEarthRadius/1000, types.UnitKilometers), 1.00)
	assert.Equal(LengthToRadians(types.FactorEarthRadius/1609.344, types.UnitMiles), 1.00)
}
