package cmd

import "github.com/remnikmart/pokedexcli/poke"

type cliCommand struct {
	name        CommandName
	description string
	callback    func() error
}

type pageUrls struct {
	nextUrl string
	prevUrl string
}

// Seen and caught pokemons
type seenPokemons map[string]poke.PokemonEncounter
type caughtPokemons map[string]poke.PokemonStats

type dealtPokemons struct {
	seen   seenPokemons
	caught caughtPokemons
}
