package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/remnikmart/pokedexcli/poke"
	ui "github.com/remnikmart/pokedexcli/textui"
	"github.com/remnikmart/pokedexcli/web"
)

func (c *command) commandExit() error {
	ui.ShowMessage("Closing the Pokedex... Goodbye!", "")
	os.Exit(0)
	return nil
}

func (c *command) commandHelp() error {
	content := []string{
		"Welcome to the Pokedex!",
		"Usage:",
	}
	for _, command := range c.commandMapping {
		content = append(content, fmt.Sprintf("- %s: %s", command.name, command.description))
	}
	for _, v := range content {
		ui.ShowMessage(v, "")
	}
	return nil
}

// Area
func (c *command) commandMap() error {
	return c.doAreaStep(c.areaPrevNext.nextUrl, "forward")
}

func (c *command) commandMapBack() error {
	return c.doAreaStep(c.areaPrevNext.prevUrl, "back")
}

func (c *command) doAreaStep(url string, direction string) error {
	if url == "" {
		return fmt.Errorf("no way %s", direction)
	}
	content, errContent := web.GetPage[poke.AreaResults](c.areaPrevNext.nextUrl, c.cache)
	if errContent != nil {
		return errContent
	}
	c.processArea(content)
	return nil
}

func (c *command) processArea(content poke.AreaResults) {
	c.areaPrevNext.nextUrl = content.NextUrl
	c.areaPrevNext.prevUrl = content.PrevUrl
	for _, v := range content.Results {
		ui.ShowMessage(v.Name, "")
	}
}

// Area encounters
func (c *command) commandExplore() error {
	if len(c.commandParams) == 0 {
		ui.ShowMessage("Insufficient parameters:", "Enter aria name")
		return nil
	}
	url, err := web.FormUrl(c.areaPrevNext.nextUrl, c.commandParams[0], nil, true)
	if err != nil {
		ui.ShowError("failed to form URL:", err)
	}
	content, errContent := web.GetPage[poke.PokemonEncounters](url, c.cache)
	if errContent != nil {
		ui.ShowError("failed to get area encounters data:", errContent)
	}
	c.registerSeen(content)
	for _, v := range content.Encounters {
		ui.ShowMessage(v.Pokemon.Name, "")
	}
	return nil
}

func (c *command) registerSeen(seen poke.PokemonEncounters) {
	for _, v := range seen.Encounters {
		_, exists := c.dealtPokemons.seen[v.Pokemon.Name]
		if !exists {
			c.dealtPokemons.seen[v.Pokemon.Name] = poke.PokemonEncounter{Name: v.Pokemon.Name, Url: v.Pokemon.Url}
		}
	}
}

// Pokemons
func (c *command) commandCatch() error {
	if len(c.commandParams) == 0 {
		ui.ShowMessage("Insufficient parameters:", "Enter the pokemon name to catch")
		return nil
	}
	pokeName := c.commandParams[0]
	seenPoke, seen := c.dealtPokemons.seen[pokeName]
	if !seen {
		ui.ShowMessage(fmt.Sprintf("You haven't encountered %s so far", pokeName), "")
		return nil
	}

	if _, caught := c.dealtPokemons.caught[pokeName]; caught {
		ui.ShowMessage(fmt.Sprintf("%s has been already caught!", pokeName), "")
		return nil
	} else {
		url := seenPoke.Url
		// url := "https://pokeapi.co/api/v2/pokemon/" + pokeName
		content, errContent := web.GetPage[poke.PokemonStats](url, c.cache)
		if errContent != nil {
			ui.ShowError("failed to get pokemon stats:", errContent)
			return nil
		}
		ui.ShowMessage(fmt.Sprintf("Throwing a Pokeball at %s...", pokeName), "")
		ui.ShowMessage(fmt.Sprintf("%s was caught!", pokeName), "")
		c.registerCaught(content)
	}
	return nil
}

func (c *command) commandInspect() error {
	if len(c.commandParams) == 0 {
		ui.ShowMessage("Insufficient parameters:", "Enter the pokemon name to view")
		return nil
	}
	pokeName := c.commandParams[0]
	caughtPoke, caught := c.dealtPokemons.caught[pokeName]
	var content poke.PokemonStats
	if caught {
		ui.ShowMessage(fmt.Sprintf("%s has been already caught!", pokeName), "")
		content = caughtPoke
	} else {
		ui.ShowMessage(fmt.Sprintf("%s hasn't been caught yet", pokeName), "")
		return nil
	}
	showStats(content)
	return nil
}

func showStats(stats poke.PokemonStats) {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", stats.Name, stats.Height, stats.Weight))
	for _, v := range stats.Stats {
		b.WriteString(fmt.Sprintf("  - %s: %d\n", v.PokemonStat.Name, v.BaseStat))
	}
	b.WriteString("Types:\n")
	for _, v := range stats.Types {
		b.WriteString(fmt.Sprintf("  - %s\n", v.PokemonType.Name))
	}
	ui.ShowMessage(b.String(), "")
}

func (c *command) registerCaught(stats poke.PokemonStats) {
	if _, caught := c.dealtPokemons.caught[stats.Name]; caught {
		return
	}
	c.dealtPokemons.caught[stats.Name] = stats
}

func (c *command) commandPokedex() error {
	ui.ShowMessage("Your Pokedex:", "")
	for k, _ := range c.dealtPokemons.caught {
		ui.ShowMessage(fmt.Sprintf("  - %s", k), "")
	}
	return nil
}
