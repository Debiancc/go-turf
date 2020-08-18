package point

type Coord [2]float64
type Properties map[string]interface{}

type Point struct {
	Coord Coord
	Lng   float64
	Lat   float64
	Properties
	Options Options
}

type Options struct {
	Bbox float64
	Id   interface{} // int64 or string
}

func NewPoint(coordinates Coord, properties *Properties, options *Options) Point {
	p := Point{
		Lng:   coordinates[0],
		Lat:   coordinates[1],
		Coord: coordinates,
	}
	if properties != nil {
		p.Properties = *properties
	}
	if options != nil {
		p.Options = *options
	}
	return p
}
