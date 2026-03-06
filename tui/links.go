package tui

import (
	"github.com/charmbracelet/lipgloss"
)

type link struct {
	label string
	value string
}

func renderLinks(width, height int, t theme, r *lipgloss.Renderer) string {
	links := []link{
		{label: "GitHub", value: "github.com/aoyn1xw"},
		{label: "Instagram", value: "instagram.com/ayon1xw"},
		{label: "Discord", value: "ayon1xw"},
	}

	labelStyle := r.NewStyle().
		Foreground(lipgloss.Color(t.Accent)).
		Bold(true)

	valueStyle := r.NewStyle().
		Foreground(lipgloss.Color("#AAAAAA"))

	cardStyle := r.NewStyle().
		Width(30).
		PaddingLeft(2)

	var cards []string
	for _, l := range links {
		card := cardStyle.Render(
			labelStyle.Render(l.label) + "\n" + valueStyle.Render(l.value),
		)
		cards = append(cards, card)
	}

	// 2 per row
	row1 := lipgloss.JoinHorizontal(lipgloss.Top, cards[0], cards[1])
	row2 := lipgloss.JoinHorizontal(lipgloss.Top, cards[2])

	grid := lipgloss.JoinVertical(lipgloss.Left, row1, row2)

	artStyle := r.NewStyle().
		Foreground(lipgloss.Color(t.Accent))

	art := artStyle.Render(`
⠀⠀⠀⢸⣦⡀⠀⠀⠀⠀⢀⡄
⠀⠀⠀⢸⣏⠻⣶⣤⡶⢾⡿⠁⠀⢠⣄⡀⢀⣴
⠀⠀⣀⣼⠷⠀⠀⠁⢀⣿⠃⠀⠀⢀⣿⣿⣿⣇
⠴⣾⣯⣅⣀⠀⠀⠀⠈⢻⣦⡀⠒⠻⠿⣿⡿⠿⠓⠂⠀⠀⢀⡇
⠀⠀⠀⠉⢻⡇⣤⣾⣿⣷⣿⣿⣤⠀⠀⣿⠁⠀⠀⠀⢀⣴⣿⣿
⠀⠀⠀⠀⠸⣿⡿⠏⠀⢀⠀⠀⠿⣶⣤⣤⣤⣄⣀⣴⣿⡿⢻⣿⡆
⠀⠀⠀⠀⠀⠟⠁⠀⢀⣼⠀⠀⠀⠹⣿⣟⠿⠿⠿⡿⠋⠀⠘⣿⣇
⠀⠀⠀⠀⠀⢳⣶⣶⣿⣿⣇⣀⠀⠀⠙⣿⣆⠀⠀⠀⠀⠀⠀⠛⠿⣿⣦⣤⣀
⠀⠀⠀⠀⠀⠀⣹⣿⣿⣿⣿⠿⠋⠁⠀⣹⣿⠳⠀⠀⠀⠀⠀⠀⢀⣠⣽⣿⡿⠟⠃
⠀⠀⠀⠀⠀⢰⠿⠛⠻⢿⡇⠀⠀⠀⣰⣿⠏⠀⠀⢀⠀⠀⠀⣾⣿⠟⠋⠁
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠋⠀⠀⣰⣿⣿⣾⣿⠿⢿⣷⣀⢀⣿⡇⠁
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠋⠉⠁⠀⠀⠀⠀⠙⢿⣿⣿⠇
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⢿`)

	content := lipgloss.JoinVertical(lipgloss.Center, art, "", grid)

	return r.NewStyle().
		Width(width).
		Height(height).
		Align(lipgloss.Center, lipgloss.Center).
		Render(content)
}
