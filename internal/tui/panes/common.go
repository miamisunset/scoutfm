package panes

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"

	fz "github.com/miamisunset/scoutfm/internal/fs"
	"github.com/miamisunset/scoutfm/internal/tui/styles"
)

// Pane messages

type fileMsg struct {
	cwd  string
	file fs.FileInfo
}

type commonPane struct {
	styles       *styles.Styles
	Width        int
	Height       int
	Cursor       int
	Browsable    bool
	Files        []fs.FileInfo
	selectedFile fs.FileInfo
	currentDir   string
	contentType  string
}

func newCommonPane(width int, browse bool) commonPane {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal()
	}

	files := fz.ReadDir(cwd)
	fz.SortByName(files, true)
	fz.SortDirFirst(files)

	return commonPane{
		styles:       styles.DefaultStyles(),
		Width:        width,
		Height:       0,
		Cursor:       0,
		Files:        files,
		Browsable:    browse,
		currentDir:   cwd,
		selectedFile: nil,
	}
}

func (c *commonPane) View() string {
	b := strings.Builder{}
	b.WriteString(c.fileBrowser())
	return c.styles.App.Render(b.String())
}

func (c commonPane) fileBrowser() string {

	var fileList string

	for i, file := range c.Files {
		cursor := " "

		filename := file.Name()

		if c.Cursor == i && c.Browsable {
			cursor = "|"

			sb := strings.Builder{}
			sb.WriteString(cursor)
			sb.WriteRune(' ')
			sb.WriteString(filename)

			selected := c.styles.App.
				Foreground(lipgloss.Color("#14F9D5")).
				Render(sb.String())

			fileList += fmt.Sprintf("%s\n", selected)
		} else {
			if file.IsDir() {
				filename = c.styles.App.
					Foreground(lipgloss.Color("#F25D94")).
					Render(file.Name())
			}

			fileList += fmt.Sprintf("%s %s\n", cursor, filename)
		}
	}

	if c.Files == nil {
		return c.contentType
	}

	return c.styles.App.
		BorderStyle(c.styles.FileBrowserBorder).
		BorderForeground(c.styles.BorderColor).
		Width(c.Width).
		Height(c.Height).
		Render(fileList)
}
