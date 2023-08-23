package turf

type IFeature interface {
	isFeature() bool
	GetGeometryType() string
}

type Options struct {
	Bbox float64
	Id   string
}

type Feature[T Coords] struct {
	Type       string
	Geometry   *Geometry[T]
	Coords     T
	Properties *Properties
}

func (f Feature[T]) isFeature() bool {
	return true
}

func (f Feature[T]) GetGeometryType() string {
	return f.Coords.GetType()
}

func (f Feature[T]) GetCoord() T {
	return f.Geometry.Coordinates
}

//func (f Feature[Point]) GetLat[T Feature[Point]]() T {
//	return f.Geometry.Coordinates
//}

func NewFeature[T Coords](coord T) *Feature[T] {
	return &Feature[T]{
		Type:   "Feature",
		Coords: coord,
		Geometry: &Geometry[T]{
			Type:        coord.GetType(),
			Coordinates: coord,
		},
	}
}
