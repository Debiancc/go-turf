package features

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
