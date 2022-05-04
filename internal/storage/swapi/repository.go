package swapi

import (
	"encoding/json"
	"fmt"
	"github.com/yulio94/star-wars-cli/internal"
	"io"
	"net/http"
)

const (
	starshipsEndpoint = "/starships"
	planetsEndpoint   = "/planets"
	peopleEndpoint    = "/people"
	swapiURL          = "https://swapi.dev/api"
)

type swapiRepo struct {
	url string
}

type swapiPeopleRepo struct {
	swapiRepo
}

type swapiPlanetsRepo struct {
	swapiRepo
}
type swapiStarshipsRepo struct {
	swapiRepo
}

func NewSwapiPeopleRepo() internal.PeopleRepo {
	return &swapiPeopleRepo{swapiRepo{url: swapiURL}}
}

func NewSwapiPlanetsRepo() internal.PlanetRepo {
	return &swapiPlanetsRepo{swapiRepo{url: swapiURL}}
}

func NewSwapiStarshipsRepo() internal.StarshipsRepo {
	return &swapiStarshipsRepo{swapiRepo{url: swapiURL}}
}

func (r *swapiPeopleRepo) GetPerson(id int) (person internal.Person, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v/%v", r.url, peopleEndpoint, id))

	if err != nil {
		return person, err
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return person, err
	}

	err = json.Unmarshal(contents, &person)
	if err != nil {
		return person, err
	}

	return
}

func (r *swapiPeopleRepo) GetPeople() (people []internal.Person, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v", r.url, peopleEndpoint))

	if err != nil {
		return nil, err
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &people)
	if err != nil {
		return nil, err
	}

	return
}

func (r *swapiPlanetsRepo) GetPlanet(id int) (planet internal.Planet, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v/%v", r.url, planetsEndpoint, id))

	if err != nil {
		return planet, err
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return planet, err
	}

	err = json.Unmarshal(contents, &planet)
	if err != nil {
		return planet, err
	}

	return

}

func (r *swapiPlanetsRepo) GetPlanets() (planets []internal.Planet, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v", r.url, planetsEndpoint))

	if err != nil {
		return nil, err
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &planets)
	if err != nil {
		return nil, err
	}

	return
}

func (r *swapiStarshipsRepo) GetStarship(id int) (starship internal.Starship, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v/%v", r.url, starshipsEndpoint, id))

	if err != nil {
		return starship, err
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return starship, err
	}

	err = json.Unmarshal(contents, &starship)
	if err != nil {
		return starship, err
	}

	return
}

func (r *swapiStarshipsRepo) GetStarships() (starships []internal.Starship, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v", r.url, starshipsEndpoint))

	if err != nil {
		return nil, err
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &starships)
	if err != nil {
		return nil, err
	}

	return
}
