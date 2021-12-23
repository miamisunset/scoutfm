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

func (c Cwd) Init() tea.Cmd {
	return c.sendSelectedFileMessage
}

func (c *Cwd) Update(msg tea.Msg) (*Cwd, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "up", "k":
			if c.Cursor > 0 {
				c.Cursor--
				cmd = c.sendSelectedFileMessage
			}

		case "down", "j":
			if c.Cursor < len(c.Files)-1 {
				c.Cursor++
				cmd = c.sendSelectedFileMessage
			}
		}
	}

	return c, cmd
}

func (c *Cwd) sendSelectedFileMessage() tea.Msg {
	return fileMsg{
		cwd:  c.currentDir,
		file: c.Files[c.Cursor],
	}
}
