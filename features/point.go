package features

type PointGeometry struct {
	Geometry
	Coordinates Position
}

type Point struct {
	Feature
	Geometry PointGeometry
}

type Options struct {
	Bbox float64
	Id   interface{} // int64 or string
}

func NewPoint(coordinates Position, properties *Properties, options *Options) *Point {
	p := new(Point)
	p.Type = "Feature"
	p.Geometry.Type = "Point"
	p.Geometry.Coordinates = coordinates
	if properties != nil {
		p.Properties = *properties
	}
	return p
}

func (p *Point) GetCoord() Position {
	return p.Geometry.Coordinates
}

func (p *Point) GetLng() float64 {
	return p.Geometry.Coordinates[0]
}

func (p *Point) GetLat() float64 {
	return p.Geometry.Coordinates[1]
}
