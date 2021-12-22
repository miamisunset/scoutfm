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
}

func NewCwdPane(width, height int) Cwd {
	w := int(float64(width) * cwdPaneMaxWidth)

	return Cwd{
		commonPane: newCommonPane(w, true),
	}
}

func (p *Cwd) Update(msg tea.Msg) (*Cwd, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "up", "k":
			if p.Cursor > 0 {
				p.Cursor--
				cmds = append(cmds, p.sendSelectedFileMessage)
			}

		case "down", "j":
			if p.Cursor < len(p.Files)-1 {
				p.Cursor++
				cmds = append(cmds, p.sendSelectedFileMessage)
			}
		}
	}

	return p, tea.Batch(cmds...)
}

func (c *Cwd) sendSelectedFileMessage() tea.Msg {
	return fileMsg{
		cwd:  c.currentDir,
		file: c.Files[c.Cursor],
	}
}
