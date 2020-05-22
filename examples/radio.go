package main

import (
	"fmt"

	"github.com/IMQS/options"
)

func main() {
	con := options.NewConsole()
	defer con.Close()

	choice := con.Radio("Your next move", "", -1, []string{"down", "left", "back"})
	con.Close()
	fmt.Printf("Your choice was: %v\n", choice)
}
