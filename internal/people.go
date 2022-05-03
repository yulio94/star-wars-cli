package internal

type PeopleRepo interface {
	GetPerson(id int) (Person, error)
	GetPeople() ([]Person, error)
}

type Person struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	Height    float32  `json:"height"`
	Mass      float32  `json:"mass"`
	HairColor string   `json:"hair_color"`
	EyeColor  string   `json:"eye_color"`
	BirthYear string   `json:"birth_year"`
	Gender    string   `json:"gender"`
	HomeWorld string   `json:"homeworld"`
	Films     []string `json:"films"`
}

func NewPerson(name string, age int, height float32, mass float32, hairColor string, eyeColor string, birthYear string, gender string, homeWorld string, films []string) (p Person) {
	p = Person{
		Name:      name,
		Age:       age,
		Height:    height,
		Mass:      mass,
		HairColor: hairColor,
		EyeColor:  eyeColor,
		BirthYear: birthYear,
		Gender:    gender,
		HomeWorld: homeWorld,
		Films:     films,
	}
	return
}
