package options

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell"
)

// Console represents a terminal console that you can control for UI
type Console struct {
	Screen         tcell.Screen
	NoExitOnEscape bool // If true, then commands will not issue an os.Exit() if the user presses Escape. The user can still press Ctrl-C to perform an os.Exit()
}

// DrawString is a helper function for drawing a string at a given location
func (c *Console) DrawString(x, y int, str string, style tcell.Style) {
	for i, ch := range str {
		c.Screen.SetContent(x+i, y, ch, nil, style)
	}
}

// Exit closes the console neatly, then does an os.Exit(1)
func (c *Console) Exit() {
	c.Close()
	os.Exit(1)
}

// Close the console neatly
func (c *Console) Close() {
	if c.Screen != nil {
		c.Screen.Fini()
	}
}

// EventPoll is a helper function that runs a Goroutine to perform event handling,
// and handles common things such as ENTER press, ESCAPE, Ctrl-C, Ctrl-L.
// eventHandler returns true to indicate that it has handled the event, or
// false to pass control onto this function.
func (c *Console) EventPoll(quit chan bool, eventHandler func(ev tcell.Event) bool) {
	go func() {
		for {
			ev := c.Screen.PollEvent()
			if eventHandler(ev) {
				// custom handler has already taken care of it
				continue
			}
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEnter:
					quit <- true
					return
				case tcell.KeyEscape:
					quit <- false
					return
				case tcell.KeyCtrlL:
					c.Screen.Sync()
				case tcell.KeyCtrlC:
					c.Exit()
				}
			case *tcell.EventResize:
				c.Screen.Sync()
			}
		}
	}()
}

// UILoop runs the standard event loop
func (c *Console) UILoop(quit chan bool, render func()) bool {
	result := false
loop:
	for {
		select {
		case result = <-quit:
			break loop
		case <-time.After(time.Millisecond * 50):
		}
		render()
	}
	if !result && !c.NoExitOnEscape {
		c.Exit()
	}
	return result
}

// NewConsole creates a new console option, and initializes it.
// If anything fails, the function calls os.Exit(1)
// You must call Close() when you are finished using a console object,
// otherwise you will leave your user's terminal in an unusable state.
func NewConsole() Console {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e = s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	return Console{Screen: s}
}
