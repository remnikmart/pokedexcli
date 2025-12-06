package textui

import "fmt"

func ShowInvitation() {
	fmt.Print("Pokedex > ")
}

func ShowMessage(caption string, msg string) {
	fmt.Printf("%s %s\n", caption, msg)
}

func ShowError(caption string, err error) {
	fmt.Printf("%s %v\n", caption, err)
}
