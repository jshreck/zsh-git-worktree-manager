package style

import "github.com/charmbracelet/lipgloss"

// Layout constants.
const (
	MinTermWidth    = 80
	LeftColumnRatio = 0.35
	MinLeftWidth    = 25
)

// Colors — adaptive (dark/light terminal detection handled by lipgloss).
var (
	AccentColor    = lipgloss.AdaptiveColor{Light: "33", Dark: "87"}  // Blue / Cyan
	SubtleColor    = lipgloss.AdaptiveColor{Light: "245", Dark: "241"}
	HighlightColor = lipgloss.AdaptiveColor{Light: "213", Dark: "213"} // Pink
	TextColor      = lipgloss.AdaptiveColor{Light: "0", Dark: "252"}
	MutedColor     = lipgloss.AdaptiveColor{Light: "250", Dark: "238"}
)

// Header styles.
var (
	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(AccentColor).
			Padding(0, 1)

	HeaderRepoStyle = lipgloss.NewStyle().
			Foreground(SubtleColor).
			Padding(0, 1)
)

// Menu styles.
var (
	MenuSectionHeader = lipgloss.NewStyle().
				Bold(true).
				Foreground(AccentColor).
				MarginBottom(1)

	MenuItemNormal = lipgloss.NewStyle().
			Foreground(TextColor).
			PaddingLeft(2)

	MenuItemHighlighted = lipgloss.NewStyle().
				Foreground(HighlightColor).
				Bold(true).
				PaddingLeft(1)

	MenuItemCurrent = lipgloss.NewStyle().
			Foreground(SubtleColor).
			Italic(true)
)

// Content panel styles.
var (
	ContentTitle = lipgloss.NewStyle().
			Bold(true).
			Foreground(AccentColor).
			MarginBottom(1)

	ContentBody = lipgloss.NewStyle().
			Foreground(TextColor)

	ContentMuted = lipgloss.NewStyle().
			Foreground(SubtleColor)
)

// Layout container styles.
var (
	LeftColumnStyle = lipgloss.NewStyle().
			BorderRight(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(MutedColor).
			Padding(1, 1)

	RightColumnStyle = lipgloss.NewStyle().
				Padding(1, 2)

	FooterStyle = lipgloss.NewStyle().
			Foreground(SubtleColor).
			Padding(0, 1)

	HeaderBarStyle = lipgloss.NewStyle().
			BorderBottom(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(MutedColor).
			Padding(0, 1).
			MarginBottom(0)
)

// CalculateColumnWidths returns left and right column widths for the given
// terminal width. If the terminal is narrower than MinTermWidth the right
// column width is 0, signalling single-column mode.
func CalculateColumnWidths(termWidth int) (left, right int) {
	if termWidth < MinTermWidth {
		return termWidth, 0
	}
	left = int(float64(termWidth) * LeftColumnRatio)
	if left < MinLeftWidth {
		left = MinLeftWidth
	}
	// Account for border character.
	right = termWidth - left - 1
	if right < 0 {
		right = 0
	}
	return left, right
}
