package poke

// Area
type Area struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type AreaResults struct {
	Count   int    `json:"count"`
	NextUrl string `json:"next"`
	PrevUrl string `json:"previous"`
	Results []Area `json:"results"`
}

// Encounters
type PokemonEncounter struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonEncounterWrapped struct {
	Pokemon PokemonEncounter `json:"pokemon"`
}

type PokemonEncounters struct {
	Encounters []PokemonEncounterWrapped `json:"pokemon_encounters"`
}

// Pokemon stats
type PokemonStat struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonStatWrapped struct {
	BaseStat    int         `json:"base_stat"`
	PokemonStat PokemonStat `json:"stat"`
}

type PokemonType struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonTypeWrapped struct {
	Slot        int         `json:"slot"`
	PokemonType PokemonType `json:"type"`
}

type PokemonStats struct {
	Name   string               `json:"name"`
	Height int                  `json:"height"`
	Weight int                  `json:"weight"`
	Stats  []PokemonStatWrapped `json:"stats"`
	Types  []PokemonTypeWrapped `json:"types"`
}
