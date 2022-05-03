package internal

type StarshipsRepo interface {
	GetStarship(id int) (Starship, error)
	GetStarships() ([]Starship, error)
}

type Starship struct {
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	Manufacturer         string   `json:"manufacturer"`
	CostInCredits        int      `json:"cost_in_credits"`
	Length               int      `json:"length"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed"`
	Crew                 int      `json:"crew"`
	Passengers           string   `json:"passengers"`
	CargoCapacity        int      `json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	HyperdriveRating     float32  `json:"hyperdrive_rating"`
	MGLT                 int      `json:"mglt"`
	StarshipClass        string   `json:"starship_class"`
	Films                []string `json:"films"`
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
