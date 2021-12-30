package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/miamisunset/scoutfm/tui/styles"
)

const (
	appName = "SCOUT"
	height  = 1
)

// Header
type header struct {
	style *lipgloss.Style

	width  int
	height int

	title *title
	cwd   *cwd
	clock *clock
}

func newHeader(width int, styles styles.Style) *header {
	return &header{
		style:  styles.GetHeader(),
		width:  width,
		height: height,
		title:  newTitle(styles),
		cwd:    newCwd(styles),
		clock:  newClock(styles),
	}
}

func (h header) update(msg tea.Msg) (header, tea.Cmd) {
	return h, nil
}

func (h header) view() string {
	width := lipgloss.Width

	title := h.title.view()
	clock := h.clock.view()
	cwd := h.cwd.view(h.width - width(title) - width(clock))

	header := lipgloss.JoinHorizontal(
		lipgloss.Top,
		title,
		cwd,
		clock,
	)

	return h.style.Width(h.width).Render(header)
}

func (h *header) setWidth(w int) {
	h.width = w
}

// Title
type title struct {
	styles *lipgloss.Style

	name string
}

func newTitle(styles styles.Style) *title {
	return &title{
		styles: styles.GetTitle(),
		name:   appName,
	}
}

func (t title) update(msg tea.Msg) (title, tea.Cmd) {
	return t, nil
}

func (t title) view() string {
	return t.styles.Render(t.name)
}

// Current working directory
type cwd struct {
	styles *lipgloss.Style
	dir    string
}

func newCwd(styles styles.Style) *cwd {
	return &cwd{
		styles: styles.GetCwd(),
		dir:    "/usr/local/bin",
	}
}

func (c cwd) update(msg tea.Msg) (cwd, tea.Cmd) {
	return c, nil
}

func (c cwd) view(width int) string {
	return c.styles.Width(width).Render(c.dir)
}

// Clock
type clock struct {
	styles *lipgloss.Style
	name   string
}

func newClock(styles styles.Style) *clock {
	return &clock{
		styles: styles.GetClock(),
		name:   "TIME",
	}
}

func (c clock) update(msg tea.Msg) (clock, tea.Cmd) {
	return c, nil
}

func (c clock) view() string {
	return c.styles.Render(c.name)
}
