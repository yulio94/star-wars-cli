package internal

type PlanetRepo interface {
	GetPlanet(id int) (Planet, error)
	GetPlanets() ([]Planet, error)
}

type Planet struct {
	Name           string   `json:"name"`
	RotationPeriod int      `json:"rotation_period"`
	OrbitalPeriod  int      `json:"orbital_period"`
	Diameter       int      `json:"diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   int      `json:"surface_water"`
	Population     int      `json:"population"`
	Films          []string `json:"films"`
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
