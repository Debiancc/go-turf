package point

import "github.com/Debiancc/go-turf/features"

type Geometry struct {
	features.Geometry
	Coordinates features.Position
}

type Point struct {
	features.Feature
	Geometry Geometry
}

type Options struct {
	Bbox float64
	Id   interface{} // int64 or string
}

func NewPoint(coordinates features.Position, properties *features.Properties, options *Options) *Point {
	p := new(Point)
	p.Type = "Feature"
	p.Geometry.Type = "Point"
	p.Geometry.Coordinates = coordinates
	if properties != nil {
		p.Properties = *properties
	}
	return p
}
