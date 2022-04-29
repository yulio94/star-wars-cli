package internal

type Planet struct {
	Name           string
	RotationPeriod int
	OrbitalPeriod  int
	Diameter       int
	Climate        string
	Gravity        string
	Terrain        string
	SurfaceWater   int
	Population     int
	Films          []string
}

func NewPlanet(name string, rotationPeriod int, orbitalPeriod int, diameter int, climate string, gravity string, terrain string, surfaceWater int, population int, films []string) (p Planet) {
	p = Planet{
		Name:           name,
		RotationPeriod: rotationPeriod,
		OrbitalPeriod:  orbitalPeriod,
		Diameter:       diameter,
		Climate:        climate,
		Gravity:        gravity,
		Terrain:        terrain,
		SurfaceWater:   surfaceWater,
		Population:     population,
		Films:          films,
	}
	return
}
