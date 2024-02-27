package models

type StarWarsCharacter struct {
	first_name  string
	middle_name string
	last_name   string
	parents     []string
	siblings    []string
	children    []string
	birth_year  string
	cybernetics string
	gender      string
	height      int // in centimeters
	mass        int // in kilograms
	hair_color  string
	skin_color  string
	eye_color   string
	masters     []string
	apprentices []string
	homeworld   string
	main_weapon string
	species     []string
	vehicles    []string
	starships   []string
	died        []string
	tags        []string // Affiliated tags, were they a jedi, sith, clone, bountry hunter, etc.
}
