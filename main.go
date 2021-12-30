package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/miamisunset/scoutfm/tui"
)

func main() {
	p := tea.NewProgram(tui.NewScout("."), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
