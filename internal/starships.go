package internal

type StarshipsRepo interface {
	GetStarship(id int) (Starship, error)
	GetStarships() ([]Starship, error)
}

type Starship struct {
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	Manufacturer         string   `json:"manufacturer"`
	CostInCredits        string   `json:"cost_in_credits"`
	Length               string   `json:"length"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed"`
	Crew                 string   `json:"crew"`
	Passengers           string   `json:"passengers"`
	CargoCapacity        string   `json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	HyperdriveRating     string   `json:"hyperdrive_rating"`
	MGLT                 string   `json:"mglt"`
	StarshipClass        string   `json:"starship_class"`
	Films                []string `json:"films"`
}

func NewStarship(name string, model string, manufacturer string, costInCredits string, length string, maxAtmospheringSpeed string, crew string, passengers string, cargoCapacity string, consumables string, hyperdriveRating string, mglt string, starshipClass string, films []string) (s Starship) {
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
