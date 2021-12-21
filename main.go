package main

import (
	"golang.org/x/term"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"gitlab.com/Synthwave/scoutfm/internal/tui/colors"
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
	s.BorderColor = colors.Green

	return s
}

type scout struct {
	styles     *Styles
	termWidth  int
	termHeight int
	cwd        string // current working directory
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

	fb := strings.Builder{}

	files := readDir(s.cwd)
	l := len(files) - 1

	for i, f := range files {
		fb.WriteString(f.Name())
		if i < l {
			fb.WriteRune('\n')
		}
	}

	return s.styles.App.
		BorderStyle(s.styles.FileBrowserBorder).
		BorderForeground(s.styles.BorderColor).
		Width(s.termWidth).
		Height(s.termHeight).
		Render(fb.String())
}

func (s scout) View() string {
	b := strings.Builder{}
	b.WriteString(s.headerView())
	b.WriteRune('\n')
	b.WriteString(s.fileBrowser())
	return s.styles.App.Render(b.String())
}

func NewScout(cwd string, termWidth int, termHeight int) *scout {
	return &scout{
		styles:     DefaultStyles(),
		cwd:        cwd,
		termWidth:  termWidth - 2,
		termHeight: termHeight - 3,
	}
}

func readDir(cwd string) []fs.FileInfo {
	files, err := ioutil.ReadDir(cwd)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func main() {
	twidth, theight, _ := term.GetSize(int(os.Stdout.Fd()))

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(NewScout(cwd, twidth, theight), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
