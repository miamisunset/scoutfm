package tui

import (
	"fmt"
	"io/fs"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/miamisunset/scoutfm/tui/styles"
)

type footer struct {
	width  int
	height int

	fileInfo *fileInfo

	style *lipgloss.Style
}

func newFooter(width int, style styles.Style) *footer {
	return &footer{
		width:    width,
		style:    style.GetFooter(),
		fileInfo: newFileInfo(style),
	}
}

func (f *footer) update(msg tea.Msg) (*footer, tea.Cmd) {
	f.fileInfo.update(msg)
	return f, nil
}

func (f footer) view() string {
	fi := f.fileInfo.view()

	footer := lipgloss.JoinHorizontal(
		lipgloss.Top,
		fi,
	)

	return f.style.Width(f.width).Render(footer)
}

func (f *footer) setWidth(w int) {
	f.width = w
}

type fileInfo struct {
	mode  fs.FileMode
	style *lipgloss.Style
}

func newFileInfo(style styles.Style) *fileInfo {
	return &fileInfo{
		style: style.GetFileInfo(),
	}
}

func (f *fileInfo) update(msg tea.Msg) (*fileInfo, tea.Cmd) {
	switch msg := msg.(type) {
	case selectedFileMsg:
		f.mode = msg.file.Mode()
	}

	return f, nil
}

func (f fileInfo) view() string {
	return f.style.Render(fmt.Sprintf("%d", f.mode))
}
