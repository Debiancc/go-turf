# go-turf
[![Build Status](https://travis-ci.com/Debiancc/go-turf.svg?branch=master)](https://travis-ci.com/Debiancc/go-turf)
### Still in WIP

turfjs lib implement by Golang.

### Usage

```golang
import (
	"github.com/Debiancc/go-turf"
)
func main() {
    form := point.NewPoint([2]float64{-75.343, 39.984}, nil, nil)
    to := point.NewPoint([2]float64{-75.534, 39.123}, nil, nil)

    distance := Distance(form, to, types.UnitDegrees)
    // 0.8724834600465156
}
```
