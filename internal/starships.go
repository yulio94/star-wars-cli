package internal

type Starship struct {
	Name                 string
	Model                string
	Manufacturer         string
	CostInCredits        int
	Length               int
	MaxAtmospheringSpeed string
	Crew                 int
	Passengers           string
	CargoCapacity        int
	Consumables          string
	HyperdriveRating     float32
	MGLT                 int
	StarshipClass        string
	Films                []string
}

func NewStarship(name string, model string, manufacturer string, costInCredits int, length int, maxAtmospheringSpeed string, crew int, passengers string, cargoCapacity int, consumables string, hyperdriveRating float32, mglt int, starshipClass string, films []string) (s Starship) {
	s = Starship{
		Name:                 name,
		Model:                model,
		Manufacturer:         manufacturer,
		CostInCredits:        costInCredits,
		Length:               length,
		MaxAtmospheringSpeed: maxAtmospheringSpeed,
		Crew:                 crew,
		Passengers:           passengers,
		CargoCapacity:        cargoCapacity,
		Consumables:          consumables,
		HyperdriveRating:     hyperdriveRating,
		MGLT:                 mglt,
		StarshipClass:        starshipClass,
		Films:                films,
	}
	return
}
