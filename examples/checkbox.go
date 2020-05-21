package main

import (
	"github.com/IMQS/options"
)

func main() {
	con := options.NewConsole()
	defer con.Close()

	boxes := []options.Checkbox{
		{false, "one"},
		{true, "number two"},
		{false, "three"},
	}

	con.CheckBoxes("Enabled Services", "Untick a service to disable it", boxes)
	con.Close()
}
