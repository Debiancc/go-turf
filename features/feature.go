package features

/*
* RFC ref https://tools.ietf.org/html/rfc7946#section-3.1.1
* [lng, lat] pair
 */
type Position [2]float64
type Properties map[string]interface{}

type Geometry struct {
	Type        string
	Coordinates interface{} // point = Position, linestring = []Position, polygon = []Position
}

type Feature struct {
	Type       string
	Geometry   Geometry
	Properties Properties
}
