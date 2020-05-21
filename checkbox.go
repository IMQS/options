package options

import (
	"github.com/gdamore/tcell"
)

// Checkbox represents an item (typically in a group of two or more), that the user can tick on/off
type Checkbox struct {
	Checked bool
	Title   string
}

// Returns true if the user pressed enter to continue, or false in any other case.
func (c *Console) CheckBoxes(title string, subTitle string, els []Checkbox) bool {
	s := c.Screen
	s.Clear()

	xCheck := 2
	yTop := 0
	item := 0

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
		s.ShowCursor(xCheck, cy+item)

		for _, el := range els {
			st := tcell.StyleDefault.Foreground(tcell.ColorWhite)
			s.SetContent(xCheck-1, cy, '[', nil, st)
			if el.Checked {
				s.SetContent(xCheck, cy, 'x', nil, st)
			} else {
				s.SetContent(xCheck, cy, ' ', nil, st)
			}
			s.SetContent(xCheck+1, cy, ']', nil, st)
			st = tcell.StyleDefault.Foreground(tcell.ColorWhite)
			c.DrawString(xCheck+3, cy, el.Title, st)
			cy++
		}

		st := tcell.StyleDefault.Foreground(tcell.ColorDarkGray)
		c.DrawString(0, cy+1, "SPACE: toggle on/off, ENTER: continue", st)

		s.Show()
	}

	quit := make(chan bool)

	c.EventPoll(quit, func(ev tcell.Event) bool {
		handled := false
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyRune:
				if ev.Rune() == ' ' {
					els[item].Checked = !els[item].Checked
					handled = true
				}
			case tcell.KeyUp:
				if item > 0 {
					item--
					handled = true
				}
			case tcell.KeyDown:
				if item < len(els)-1 {
					item++
					handled = true
				}
			}
		}
		if handled {
			render()
		}
		return handled
	})

	return c.UILoop(quit, render)
}
