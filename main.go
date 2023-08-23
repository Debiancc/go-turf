package turf

import "fmt"

//
//func main() {
//	line2 := NewLineString(LineString{
//		[2]float64{1.111, 2.2222},
//	}, nil, nil)
//	//line2.GetLat()
//	point := NewPoint([2]float64{1.111, 2.2222})
//
//	getLat[Feature[LineString]](line2)
//
//	NewFeatureCollection[Feature]()
//
//
//
//}

func acceptFeatureCollection[T Coords](features []Feature[T]) {
	for _, feature := range features {
		fmt.Println(feature.Type, feature.Geometry.Type, feature.Properties)
	}
}

func main() {

}
