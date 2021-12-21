package main

import (
	"golang.org/x/term"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	App               lipgloss.Style
	FileBrowserBorder lipgloss.Border
	BorderColor       lipgloss.Color
}

func DefaultStyles() *Styles {
	s := new(Styles)

	s.App = lipgloss.NewStyle()

	s.FileBrowserBorder = lipgloss.RoundedBorder()
	s.BorderColor = "#874BFD"

	return s
}

type scout struct {
	styles    *Styles
	termWidth int
	cwd       string // current working directory
	browser   tea.Model
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

func (s scout) headerView() string {
	return s.cwd
}

func (s scout) fileBrowser() string {
	return s.styles.App.
		BorderStyle(s.styles.FileBrowserBorder).
		BorderForeground(s.styles.BorderColor).
		Width(s.termWidth).
		Render("File1\nFile2\nFile3")
}

func (s scout) View() string {
	b := strings.Builder{}
	b.WriteString(s.headerView())
	b.WriteRune('\n')
	b.WriteString(s.fileBrowser())
	return s.styles.App.Render(b.String())
}

func NewScout(cwd string, termWidth int) *scout {
	return &scout{
		styles:    DefaultStyles(),
		cwd:       cwd,
		termWidth: termWidth - 2,
	}
}

func main() {
	tw, _, _ := term.GetSize(int(os.Stdout.Fd()))

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(NewScout(cwd, tw), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
