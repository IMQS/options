# options
`options` is a minimal Go library for console apps that want a little more than scanf.

You get the ability to easily construct terminal UI such as this:



### Controls
* Check Box
* Radio Box

### Conventions
* Space toggles a choice
* Enter accepts input
* Escape exits the program (you can turn this off with `Console.NoExitOnEscape`)

### Example
```go
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
```
For more examples, see the 'examples' directory in the source tree.

This library is built on top of https://github.com/gdamore/tcell.
