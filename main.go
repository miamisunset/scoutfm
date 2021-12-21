package main

import (
	"golang.org/x/term"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"gitlab.com/Synthwave/scoutfm/internal/tui"
)

func main() {
	twidth, theight, _ := term.GetSize(int(os.Stdout.Fd()))

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(tui.NewScout(cwd, twidth, theight), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
