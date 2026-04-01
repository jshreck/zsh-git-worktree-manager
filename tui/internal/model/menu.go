package model

import (
	"fmt"
	"strings"

	"github.com/tmbtech/zsh-git-worktree-manager/tui/internal/data"
	"github.com/tmbtech/zsh-git-worktree-manager/tui/internal/style"
)

// ActionItem represents a selectable action in the menu.
type ActionItem struct {
	Label  string // Display label.
	Action string // Machine-readable action string returned on selection.
}

// Actions is the fixed list of top-level commands.
var Actions = []ActionItem{
	{Label: "Setup new worktree", Action: "setup"},
	{Label: "Remove worktree", Action: "remove"},
	{Label: "List worktrees", Action: "list"},
	{Label: "Pull latest", Action: "pull"},
	{Label: "Review PR", Action: "review"},
	{Label: "Init directory", Action: "dir"},
}

// MenuModel holds state for the left-column menu.
type MenuModel struct {
	cursor    int
	items     []menuEntry
	totalLen  int
}

type menuEntry struct {
	label     string
	action    string
	isCurrent bool
	isSection bool // section header (not selectable)
}

// NewMenuModel builds a MenuModel from TUI data.
func NewMenuModel(d data.TUIData) MenuModel {
	var items []menuEntry

	// Section: Actions
	items = append(items, menuEntry{label: "ACTIONS", isSection: true})
	for _, a := range Actions {
		items = append(items, menuEntry{label: a.Label, action: a.Action})
	}

	// Section: Worktrees
	if len(d.Worktrees) > 0 {
		items = append(items, menuEntry{label: "WORKTREES", isSection: true})
		for _, wt := range d.Worktrees {
			items = append(items, menuEntry{
				label:     wt.Name,
				action:    "navigate:" + wt.Name,
				isCurrent: wt.IsCurrent,
			})
		}
	}

	m := MenuModel{items: items, totalLen: len(items)}
	// Position cursor on the first selectable item.
	m.cursor = m.nextSelectable(0, 1)
	return m
}

// CursorUp moves the cursor to the previous selectable item.
func (m MenuModel) CursorUp() MenuModel {
	next := m.nextSelectable(m.cursor, -1)
	return MenuModel{cursor: next, items: m.items, totalLen: m.totalLen}
}

// CursorDown moves the cursor to the next selectable item.
func (m MenuModel) CursorDown() MenuModel {
	next := m.nextSelectable(m.cursor, 1)
	return MenuModel{cursor: next, items: m.items, totalLen: m.totalLen}
}

// SelectedAction returns the action string for the item under the cursor.
func (m MenuModel) SelectedAction() string {
	if m.cursor >= 0 && m.cursor < m.totalLen {
		return m.items[m.cursor].action
	}
	return ""
}

// HighlightedIndex returns the current cursor position.
func (m MenuModel) HighlightedIndex() int {
	return m.cursor
}

// HighlightedAction returns the action of the currently highlighted item.
func (m MenuModel) HighlightedAction() string {
	if m.cursor >= 0 && m.cursor < m.totalLen {
		return m.items[m.cursor].action
	}
	return ""
}

// View renders the menu column.
func (m MenuModel) View(width int) string {
	var b strings.Builder
	for i, item := range m.items {
		if item.isSection {
			if i > 0 {
				b.WriteString("\n")
			}
			b.WriteString(style.MenuSectionHeader.Width(width).Render(item.label))
			b.WriteString("\n")
			continue
		}

		label := item.label
		if item.isCurrent {
			label = fmt.Sprintf("%s  (current)", label)
		}

		if i == m.cursor {
			b.WriteString(style.MenuItemHighlighted.Width(width).Render("> " + label))
		} else {
			b.WriteString(style.MenuItemNormal.Width(width).Render(label))
		}
		b.WriteString("\n")
	}
	return b.String()
}

// nextSelectable finds the next selectable index starting from pos moving in
// direction dir (+1 or -1). Wraps around the list.
func (m MenuModel) nextSelectable(pos, dir int) int {
	if m.totalLen == 0 {
		return 0
	}
	candidate := pos + dir
	for range m.totalLen {
		if candidate < 0 {
			candidate = m.totalLen - 1
		} else if candidate >= m.totalLen {
			candidate = 0
		}
		if !m.items[candidate].isSection {
			return candidate
		}
		candidate += dir
	}
	return pos // fallback — everything is a section header (shouldn't happen)
}
