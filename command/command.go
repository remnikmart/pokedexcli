package cmd

import (
	"fmt"

	"github.com/remnikmart/pokedexcli/funcs"
	"github.com/remnikmart/pokedexcli/internal/pokecache"
	ui "github.com/remnikmart/pokedexcli/textui"
)

type CommandName string

const (
	Map     CommandName = "map"
	Mapb    CommandName = "mapb"
	Explore CommandName = "explore"
	Catch   CommandName = "catch"
	Inspect CommandName = "inspect"
	Pokedex CommandName = "pokedex"
	Help    CommandName = "help"
	Exit    CommandName = "exit"
)

type command struct {
	commandMapping map[CommandName]cliCommand
	areaPrevNext   pageUrls
	cache          *pokecache.Cache
	commandParams  []string
	dealtPokemons  dealtPokemons
}

func NewCommand() *command {
	c := command{}
	c.commandMapping = map[CommandName]cliCommand{
		Exit: {
			name:        Exit,
			description: "Exit the Pokedex",
			callback:    c.commandExit,
		},
		Help: {
			name:        Help,
			description: "Displays a help message",
			callback:    c.commandHelp,
		},
		Map: {
			name:        Map,
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    c.commandMap,
		},
		Mapb: {
			name:        Mapb,
			description: "Maps back to a previous page",
			callback:    c.commandMapBack,
		},
		Explore: {
			name:        Explore,
			description: "Shows pokemons in the area (needs area name as a parameter)",
			callback:    c.commandExplore,
		},
		Catch: {
			name:        Catch,
			description: "Catches a pokemon by name (needs the name as a parameter)",
			callback:    c.commandCatch,
		},
		Inspect: {
			name:        Inspect,
			description: "Shows pokemon's stats (needs the name as a parameter)",
			callback:    c.commandInspect,
		},
		Pokedex: {
			name:        Pokedex,
			description: "Shows caught pokemons",
			callback:    c.commandPokedex,
		},
	}
	c.areaPrevNext.nextUrl = "https://pokeapi.co/api/v2/location-area/"
	c.cache = pokecache.NewCache(5)
	c.dealtPokemons.seen = make(seenPokemons)
	c.dealtPokemons.caught = make(caughtPokemons)
	return &c
}

func (c *command) ProcessCommand(input string) {
	res := funcs.CleanInput(input)
	var s string
	if len(res) > 0 {
		s = res[0]
	} else {
		s = ""
	}
	command, errCommand := c.commandMapping[CommandName(s)]
	if !errCommand {
		ui.ShowMessage("Unknown command", "")
		return
	}
	c.commandParams = res[1:]
	err := command.callback()
	if err != nil {
		ui.ShowError(fmt.Sprintf("Error in %s:", command.name), err)
	}
}
