package options

import (
	"github.com/gdamore/tcell"
)

// Returns the index of the chosen item, or -1 if the user pressed escape
// choice is the index of the currently chosen item
func (c *Console) Radio(title string, subTitle string, choice int, els []string) int {
	if len(els) == 0 {
		return -1
	}
	if choice < 0 {
		choice = 0
	} else if choice >= len(els) {
		choice = len(els) - 1
	}
	s := c.Screen
	s.Clear()

	xCheck := 2
	yTop := 0

	if title != "" {
		c.DrawString(0, yTop, title, tcell.StyleDefault.Foreground(tcell.ColorGreenYellow))
		yTop++
	}
	if subTitle != "" {
		c.DrawString(0, yTop, subTitle, tcell.StyleDefault.Foreground(tcell.ColorLightGray))
		yTop++
	}

	if yTop != 0 {
		// separator line
		yTop++
	}

	render := func() {
		cy := yTop

		for i, el := range els {
			st := tcell.StyleDefault.Foreground(tcell.ColorWhite)
			s.SetContent(xCheck-1, cy, '(', nil, st)
			if i == choice {
				s.ShowCursor(xCheck, cy)
				s.SetContent(xCheck, cy, 'o', nil, st)
			} else {
				s.SetContent(xCheck, cy, ' ', nil, st)
			}
			s.SetContent(xCheck+1, cy, ')', nil, st)
			st = tcell.StyleDefault.Foreground(tcell.ColorWhite)
			c.DrawString(xCheck+3, cy, el, st)
			cy++
		}

		st := tcell.StyleDefault.Foreground(tcell.ColorDarkGray)
		c.DrawString(0, cy+1, "ENTER: continue", st)

		s.Show()
	}

	quit := make(chan bool)

	c.EventPoll(quit, func(ev tcell.Event) bool {
		handled := false
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEnter:
				if choice < 0 || choice >= len(els) {
					// The user has not yet chosen an item, so cannot proceed
					handled = true
				}
			case tcell.KeyRune:
				if ev.Rune() == ' ' {
					handled = true
				}
			case tcell.KeyUp:
				if choice > 0 {
					choice--
					handled = true
				}
			case tcell.KeyDown:
				if choice < len(els)-1 {
					choice++
					handled = true
				}
			}
		}
		if handled {
			render()
		}
		return handled
	})

	if !c.UILoop(quit, render) {
		return -1
	}
	return choice
}
