package main

import (
	"github.com/IMQS/options"
)

func main() {
	con := options.NewConsole()
	defer con.Close()

	// Because our initial choice is -1, the user must make a choice before continuing
	con.Radio("Your next move", "", -1, []string{"down", "left", "back"})
	con.Close()
}
