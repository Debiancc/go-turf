package features

type LineStringGeometry struct {
	Geometry
	Coordinates []Position
}

type LineString struct {
	Feature
	Geometry LineStringGeometry
}

func NewLineString(coordinates []Position, properties *Properties, options *Options) *LineString {
	p := new(LineString)
	p.Type = "Feature"
	p.Geometry.Type = "LineString"
	p.Geometry.Coordinates = coordinates
	if properties != nil {
		p.Properties = *properties
	}
	return p
}

func (p *LineString) GetCoord() []Position {
	return p.Geometry.Coordinates
}
