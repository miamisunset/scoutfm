package panes

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	cwdPaneMaxWidth = 0.30 // percentage of screen
)

type SelectedFileMsg struct {
	Name string
}

type Cwd struct {
	commonPane

	SelectedFile int
}

func NewCwdPane(width, height int) Cwd {
	w := int(float64(width) * cwdPaneMaxWidth)

	return Cwd{
		commonPane:   newCommonPane(w),
		SelectedFile: 0,
	}
}

func (p *Cwd) Update(msg tea.Msg) (*Cwd, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "up", "k":
			if p.Cursor > 0 {
				p.Cursor--
			}

		case "down", "j":
			if p.Cursor < len(p.Files)-1 {
				p.Cursor++
			}
		}
	}

	return p, nil
}
