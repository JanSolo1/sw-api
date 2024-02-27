package models

type StarWarsCharacter struct {
	FirstName   string   `json:"first_name"`
	MiddleName  string   `json:"middle_name"`
	LastName    string   `json:"last_name"`
	Parents     []string `json:"parents"`
	Siblings    []string `json:"siblings"`
	Children    []string `json:"children"`
	BirthYear   string   `json:"birth_year"`
	Cybernetics string   `json:"cybernetics"`
	Gender      string   `json:"gender"`
	Height      int      `json:"height"` // in centimeters
	Mass        int      `json:"mass"`   // in kilograms
	HairColor   string   `json:"hair_color"`
	SkinColor   string   `json:"skin_color"`
	EyeColor    string   `json:"eye_color"`
	Masters     []string `json:"masters"`
	Apprentices []string `json:"apprentices"`
	Homeworld   string   `json:"homeworld"`
	MainWeapon  string   `json:"main_weapon"`
	Species     []string `json:"species"`
	Vehicles    []string `json:"vehicles"`
	Starships   []string `json:"starships"`
	Died        []string `json:"died"`
	Tags        []string `json:"tags"` // Affiliated tags, were they a jedi, sith, clone, bountry hunter, etc.
}
