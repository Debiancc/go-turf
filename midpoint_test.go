package turf

import (
	"fmt"
	"github.com/Debiancc/go-turf/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMidPoint(t *testing.T) {
	assert := assert.New(t)

	start1 := NewPoint([2]float64{0, 0})
	end1 := NewPoint([2]float64{10, 0})
	mid1 := MidPoint(*start1, *end1)
	assert.Equal(
		fmt.Sprintf("%.6f", Distance(*start1, *mid1, types.UnitKilometers)),
		fmt.Sprintf("%.6f", Distance(*end1, *mid1, types.UnitKilometers)),
	)

	start2 := NewPoint([2]float64{0, 0})
	end2 := NewPoint([2]float64{0, 10})
	mid2 := MidPoint(*start2, *end2)
	assert.Equal(
		fmt.Sprintf("%.6f", Distance(*start2, *mid2, types.UnitKilometers)),
		fmt.Sprintf("%.6f", Distance(*end2, *mid2, types.UnitKilometers)),
	)

	start3 := NewPoint([2]float64{0, 10})
	end3 := NewPoint([2]float64{0, 0})
	mid3 := MidPoint(*start3, *end3)
	assert.Equal(
		fmt.Sprintf("%.6f", Distance(*start3, *mid3, types.UnitKilometers)),
		fmt.Sprintf("%.6f", Distance(*end3, *mid3, types.UnitKilometers)),
	)

	start4 := NewPoint([2]float64{-1, 10})
	end4 := NewPoint([2]float64{1, -1})
	mid4 := MidPoint(*start4, *end4)
	assert.Equal(
		fmt.Sprintf("%.6f", Distance(*start4, *mid4, types.UnitKilometers)),
		fmt.Sprintf("%.6f", Distance(*end4, *mid4, types.UnitKilometers)),
	)

	start5 := NewPoint([2]float64{-5, -1})
	end5 := NewPoint([2]float64{5, 10})
	mid5 := MidPoint(*start5, *end5)
	assert.Equal(
		fmt.Sprintf("%.6f", Distance(*start5, *mid5, types.UnitKilometers)),
		fmt.Sprintf("%.6f", Distance(*end5, *mid5, types.UnitKilometers)),
	)

	start6 := NewPoint([2]float64{22.5, 21.94304553343818})
	end6 := NewPoint([2]float64{92.10937499999999, 46.800059446787316})
	mid6 := MidPoint(*start6, *end6)
	assert.Equal(
		fmt.Sprintf("%.6f", Distance(*start6, *mid6, types.UnitKilometers)),
		fmt.Sprintf("%.6f", Distance(*end6, *mid6, types.UnitKilometers)),
	)
}
