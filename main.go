package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
)

type Styles struct {
	FileBrowserBorder lipgloss.Border
}

func DefaultStyles() *Styles {
	s := new(Styles)

	s.FileBrowserBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "┬",
		BottomLeft:  "├",
		BottomRight: "┴",
	}

	return s
}

type scout struct {
	styles *Styles
}

func (s scout) Init() tea.Cmd {
	return nil
}

func (s scout) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return s, tea.Quit
		}
	}

	return s, nil
}

func (s scout) View() string {
	return "ok"
}

func NewScout() *scout {
	return &scout{
		styles: DefaultStyles(),
	}
}

func main() {
	p := tea.NewProgram(NewScout(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
