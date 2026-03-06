package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tab int

const (
	aboutTab tab = iota
	linksTab
)

type theme struct {
	Accent string
	Dim    string
	Border string
}

var themes = []theme{
	{Accent: "#F5A623", Dim: "#555555", Border: "#333333"}, // Orange
	{Accent: "#00D9FF", Dim: "#555555", Border: "#333333"}, // Cyan
	{Accent: "#FF6B9D", Dim: "#555555", Border: "#333333"}, // Pink
	{Accent: "#50FA7B", Dim: "#555555", Border: "#333333"}, // Green
	{Accent: "#BD93F9", Dim: "#555555", Border: "#333333"}, // Purple
}

type Model struct {
	activeTab    tab
	width        int
	height       int
	currentTheme int
	renderer     *lipgloss.Renderer
}

func NewModel() Model {
	return Model{
		activeTab:    aboutTab,
		currentTheme: 0,
		renderer:     lipgloss.DefaultRenderer(),
	}
}

func NewModelWithRenderer(r *lipgloss.Renderer) Model {
	return Model{
		activeTab:    aboutTab,
		currentTheme: 0,
		renderer:     r,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}

		switch msg.Type {
		case tea.KeyLeft:
			if m.activeTab > aboutTab {
				m.activeTab--
			}
		case tea.KeyRight:
			if m.activeTab < linksTab {
				m.activeTab++
			}
		}

		switch msg.String() {
		case "h":
			if m.activeTab > aboutTab {
				m.activeTab--
			}
		case "l":
			if m.activeTab < linksTab {
				m.activeTab++
			}
		case "j":
			m.currentTheme = (m.currentTheme - 1 + len(themes)) % len(themes)
		case "k":
			m.currentTheme = (m.currentTheme + 1) % len(themes)
		}
	}
	return m, nil
}

func (m Model) View() string {
	if m.width == 0 {
		return ""
	}

	header := m.renderHeader()
	footer := m.renderFooter()

	// Fill remaining space
	headerHeight := 2
	footerHeight := 1
	contentHeight := m.height - headerHeight - footerHeight - 2 // 2 for newlines

	var content string
	switch m.activeTab {
	case aboutTab:
		content = renderAbout(m.width, contentHeight, themes[m.currentTheme], m.renderer)
	case linksTab:
		content = renderLinks(m.width, contentHeight, themes[m.currentTheme], m.renderer)
	}

	contentStyle := m.renderer.NewStyle().
		Height(contentHeight).
		Width(m.width)

	return header + "\n" + contentStyle.Render(content) + "\n" + footer
}

func (m Model) renderHeader() string {
	theme := themes[m.currentTheme]

	activeStyle := m.renderer.NewStyle().
		Foreground(lipgloss.Color(theme.Accent)).
		Bold(true)

	inactiveStyle := m.renderer.NewStyle().
		Foreground(lipgloss.Color(theme.Dim))

	var about, links string
	if m.activeTab == aboutTab {
		about = activeStyle.Render("About")
		links = inactiveStyle.Render("Links")
	} else {
		about = inactiveStyle.Render("About")
		links = activeStyle.Render("Links")
	}

	version := m.renderer.NewStyle().
		Foreground(lipgloss.Color(theme.Accent)).
		Bold(true).
		Render("v0.0.1")

	tabs := version + "  " + about + "  " + links

	spacerWidth := m.width - lipgloss.Width(tabs)
	if spacerWidth < 0 {
		spacerWidth = 0
	}
	spacer := m.renderer.NewStyle().Width(spacerWidth).Render("")

	line := m.renderer.NewStyle().
		Foreground(lipgloss.Color(theme.Accent)).
		Width(m.width).
		Render("─────────────────────────────────────────────────────────────────────────────────")

	return tabs + spacer + "\n" + line
}

func (m Model) renderFooter() string {
	theme := themes[m.currentTheme]

	style := m.renderer.NewStyle().
		Foreground(lipgloss.Color(theme.Dim))

	left := style.Render("← → / h l  navigate")
	mid := m.renderer.NewStyle().
		Foreground(lipgloss.Color(theme.Dim)).
		Render("j k  change color")
	right := style.Render("q quit")

	leftWidth := lipgloss.Width(left)
	rightWidth := lipgloss.Width(right)
	midWidth := lipgloss.Width(mid)
	totalPadding := m.width - leftWidth - rightWidth - midWidth
	if totalPadding < 0 {
		totalPadding = 0
	}
	leftPad := m.renderer.NewStyle().Width(totalPadding / 2).Render("")
	rightPad := m.renderer.NewStyle().Width(totalPadding - totalPadding/2).Render("")

	return left + leftPad + mid + rightPad + right
}
