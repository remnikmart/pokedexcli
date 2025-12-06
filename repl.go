package main

import (
	"bufio"
	"os"

	cmd "github.com/remnikmart/pokedexcli/command"
	ui "github.com/remnikmart/pokedexcli/textui"
)

func goRepl() {
	reader := bufio.NewReader(os.Stdin)
	cmd := cmd.NewCommand()
	for {
		ui.ShowInvitation()
		input, errInput := reader.ReadString('\n')
		if errInput != nil {
			ui.ShowError("Input error:", errInput)
			return
		}
		cmd.ProcessCommand(input)
	}
}
