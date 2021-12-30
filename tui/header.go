package tui

import (
	"os"
	"os/user"
	"strings"
	"time"

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

	title  *title
	whoAmI *whoAmI
	cwd    *cwd
	clock  *clock
}

func newHeader(width int, styles styles.Style) *header {
	return &header{
		style:  styles.GetHeader(),
		width:  width,
		height: height,
		title:  newTitle(styles),
		whoAmI: newWhoAmI(styles),
		cwd:    newCwd(styles),
		clock:  newClock(styles),
	}
}

// func (h header) update(msg tea.Msg) (header, tea.Cmd) {
// 	return h, nil
// }

func (h header) view() string {
	width := lipgloss.Width

	title := h.title.view()
	whoAmI := h.whoAmI.view()
	clock := h.clock.view()
	cwd := h.cwd.view(h.width - width(title) - width(whoAmI) - width(clock))

	header := lipgloss.JoinHorizontal(
		lipgloss.Top,
		title,
		whoAmI,
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

// func (t title) update(msg tea.Msg) (title, tea.Cmd) {
// 	return t, nil
// }

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

// func (c cwd) update(msg tea.Msg) (cwd, tea.Cmd) {
// 	return c, nil
// }

func (c cwd) view(width int) string {
	return c.styles.Width(width).Render(c.dir)
}

// Clock
type clock struct {
	styles *lipgloss.Style
	format string
}

func newClock(styles styles.Style) *clock {
	return &clock{
		styles: styles.GetClock(),
		format: "⏰ 3:04:05 pm",
	}
}

// func (c clock) update(msg tea.Msg) (clock, tea.Cmd) {
// 	return c, nil
// }

func (c clock) view() string {
	return c.styles.Render(time.Now().Format(c.format))
}

type whoAmI struct {
	styles   *lipgloss.Style
	username string
	hostname string
}

func newWhoAmI(styles styles.Style) *whoAmI {
	return &whoAmI{
		styles:   styles.GetWhoAmI(),
		username: getUsername(),
		hostname: getHostname(),
	}
}

// func (w whoAmI) update(msg tea.Msg) (whoAmI, tea.Cmd) {
// 	return w, nil
// }

func (w whoAmI) view() string {
	b := strings.Builder{}

	b.WriteString(w.username)
	b.WriteRune('@')
	b.WriteString(w.hostname)

	return w.styles.Render(b.String())
}

func getUsername() string {
	user, err := user.Current()
	if err != nil {
		return "unknown"
	}

	return user.Username
}

func getHostname() string {
	hostname, err := os.Hostname()

	if err != nil {
		return "unknown"
	}

	return hostname
}
