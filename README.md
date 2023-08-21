# go-turf
[![Go](https://github.com/Debiancc/go-turf/actions/workflows/go.yml/badge.svg)](https://github.com/Debiancc/go-turf/actions/workflows/go.yml)
### Still in WIP

turfjs lib implement by Golang.

### Usage

```golang
import (
	"github.com/Debiancc/go-turf/point"
	"github.com/Debiancc/go-turf/types"
)

func main() {
    from := point.NewPoint([2]float64{-75.343, 39.984}, nil, nil)
    to := point.NewPoint([2]float64{-75.534, 39.123}, nil, nil)

    distance := point.Distance(from, to, types.UnitDegrees)
    // 0.8724834600465156
}
```

### Completion
- MEASUREMENT **WIP**
- UNIT CONVERSION **WIP**
- COORDINATE MUTATION **TODO**
- TRANSFORMATION **TODO**
- FEATURE CONVERSION **TODO**
- MISC **TODO**
- ...
 