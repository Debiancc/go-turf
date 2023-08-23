package turf

type FeatureList = []IFeature

type FeatureCollection struct {
	Features FeatureList // todo fix use ptr
}

func NewFeatureCollection(list FeatureList) *FeatureCollection {
	return &FeatureCollection{Features: list}
}
