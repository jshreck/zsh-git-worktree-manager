package model

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tmbtech/zsh-git-worktree-manager/tui/internal/data"
	"github.com/tmbtech/zsh-git-worktree-manager/tui/internal/style"
)

// Model is the root Bubble Tea model composing header, menu, and content.
type Model struct {
	tuiData  data.TUIData
	menu     MenuModel
	content  ContentModel
	width    int
	height   int
	selected string // action selected by the user (empty until Enter pressed)
	quitting bool
}

// New creates a root Model from parsed TUI data.
func New(d data.TUIData) Model {
	return Model{
		tuiData: d,
		menu:    NewMenuModel(d),
		content: NewContentModel(d),
	}
}

// Selected returns the action string chosen by the user, or "" if they quit.
func (m Model) Selected() string {
	return m.selected
}

// Init implements tea.Model.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return Model{
			tuiData:  m.tuiData,
			menu:     m.menu,
			content:  m.content,
			width:    msg.Width,
			height:   msg.Height,
			selected: m.selected,
			quitting: m.quitting,
		}, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return Model{
				tuiData:  m.tuiData,
				menu:     m.menu,
				content:  m.content,
				width:    m.width,
				height:   m.height,
				selected: m.selected,
				quitting: true,
			}, tea.Quit

		case "up", "k":
			return Model{
				tuiData:  m.tuiData,
				menu:     m.menu.CursorUp(),
				content:  m.content,
				width:    m.width,
				height:   m.height,
				selected: m.selected,
				quitting: m.quitting,
			}, nil

		case "down", "j":
			return Model{
				tuiData:  m.tuiData,
				menu:     m.menu.CursorDown(),
				content:  m.content,
				width:    m.width,
				height:   m.height,
				selected: m.selected,
				quitting: m.quitting,
			}, nil

		case "enter":
			action := m.menu.SelectedAction()
			return Model{
				tuiData:  m.tuiData,
				menu:     m.menu,
				content:  m.content,
				width:    m.width,
				height:   m.height,
				selected: action,
				quitting: true,
			}, tea.Quit
		}
	}

	return m, nil
}

// View implements tea.Model.
func (m Model) View() string {
	if m.quitting && m.selected == "" {
		return ""
	}

	leftW, rightW := style.CalculateColumnWidths(m.width)

	header := m.renderHeader()
	footer := m.renderFooter()

	// Reserve lines for header (2 lines + border) and footer (1 line).
	bodyHeight := m.height - 4
	if bodyHeight < 1 {
		bodyHeight = 1
	}

	menuView := style.LeftColumnStyle.
		Width(leftW).
		Height(bodyHeight).
		Render(m.menu.View(leftW - 4)) // subtract padding+border

	var body string
	if rightW > 0 {
		contentAction := m.menu.HighlightedAction()
		contentView := style.RightColumnStyle.
			Width(rightW).
			Height(bodyHeight).
			Render(m.content.View(contentAction, rightW-4))
		body = lipgloss.JoinHorizontal(lipgloss.Top, menuView, contentView)
	} else {
		body = menuView
	}

	return lipgloss.JoinVertical(lipgloss.Left, header, body, footer)
}

func (m Model) renderHeader() string {
	title := style.HeaderStyle.Render("Worktree Manager")
	repoInfo := style.HeaderRepoStyle.Render(
		fmt.Sprintf("repo: %s", m.tuiData.RepoName))

	gap := m.width - lipgloss.Width(title) - lipgloss.Width(repoInfo)
	if gap < 1 {
		gap = 1
	}

	headerLine := title + fmt.Sprintf("%*s", gap, "") + repoInfo
	return style.HeaderBarStyle.Width(m.width).Render(headerLine)
}

func (m Model) renderFooter() string {
	return style.FooterStyle.Width(m.width).Render(
		"[↑/k] up  [↓/j] down  [enter] select  [q] quit")
}
