package main

import (
	"github.com/IMQS/options"
)

func main() {
	con := options.NewConsole()
	defer con.Close()

	boxes := []options.Checkbox{
		{false, "one"},
		{true, "two"},
	}

	con.CheckBoxes("Some checkboxes", "", boxes)
	// Use boxes[i].Checked to get status of the boxes

	con.Radio("Some radio buttons", "", -1, []string{"thing one", "thing two"})
	// Use return value from con.Radio to find out which items was selected

	con.Close()
}
