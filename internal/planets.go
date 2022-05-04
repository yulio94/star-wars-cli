package internal

type PlanetRepo interface {
	GetPlanet(id int) (Planet, error)
	GetPlanets() ([]Planet, error)
}

type Planet struct {
	Name           string   `json:"name"`
	RotationPeriod string   `json:"rotation_period"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Diameter       string   `json:"diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   string   `json:"surface_water"`
	Population     string   `json:"population"`
	Films          []string `json:"films"`
}

func NewPlanet(name string, rotationPeriod string, orbitalPeriod string, diameter string, climate string, gravity string, terrain string, surfaceWater string, population string, films []string) (p Planet) {
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
