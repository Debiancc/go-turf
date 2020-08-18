package point

type Coord [2]float64

type Point struct {
	Coord      Coord
	Lng        float64
	Lat        float64
	Properties map[string]interface{}
	Options    Options
}

type Options struct {
	Bbox float64
	Id   interface{} // int64 or string
}

func NewPoint(coordinates Coord, properties *map[string]interface{}, options *Options) Point {
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
