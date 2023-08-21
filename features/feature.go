package features

/*
* RFC ref https://tools.ietf.org/html/rfc7946#section-3.1.1
* [lng, lat] pair
 */
type Point [2]float64

func (p Point) GetType() string {
	return "Point"
}

type LineString [][2]float64

func (l LineString) GetType() string {
	return "LineString"
}

type Coords interface {
	Point | LineString
	GetType() string
}

type Properties map[string]interface{}

type Geometry[T Coords] struct {
	Type        string
	Coordinates T
}

type Feature[T Coords] struct {
	Type       string
	Geometry   *Geometry[T]
	Properties *Properties
}

func (f Feature[T]) GetCoord() T {
	return f.Geometry.Coordinates
}

func NewFeature[T Coords](coord T) *Feature[T] {
	return &Feature[T]{
		Type: "Feature",
		Geometry: &Geometry[T]{
			Type:        coord.GetType(),
			Coordinates: coord,
		},
	}
}

type Options struct {
	Bbox float64
	Id   string
}

func NewPoint(coord Point) *Feature[Point] {
	return NewFeature(coord)
}

func NewLineString(coord LineString, properties *Properties, options *Options) *Feature[LineString] {
	return NewFeature[LineString](coord)
}
