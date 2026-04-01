package model

import (
	"strings"
	"testing"

	"github.com/tmbtech/zsh-git-worktree-manager/tui/internal/data"
)

func testData() data.TUIData {
	return data.TUIData{
		Root:       "/tmp/repo",
		CurrentDir: "/tmp/repo/main",
		RepoName:   "repo",
		InWorktree: true,
		HasGH:      true,
		Worktrees: []data.Worktree{
			{Name: "main", Path: "/tmp/repo/main", Branch: "main", Head: "abc1234", IsCurrent: true},
			{Name: "feature-x", Path: "/tmp/repo/feature-x", Branch: "feature/x", Head: "def5678"},
		},
	}
}

func TestNewMenuModel_InitialCursorOnFirstAction(t *testing.T) {
	m := NewMenuModel(testData())
	// Cursor should skip the ACTIONS section header (index 0) and land on index 1.
	action := m.SelectedAction()
	if action != "setup" {
		t.Errorf("initial action = %q, want %q", action, "setup")
	}
}

func TestMenuModel_CursorDown(t *testing.T) {
	m := NewMenuModel(testData())

	m = m.CursorDown()
	if m.SelectedAction() != "remove" {
		t.Errorf("after one down, action = %q, want %q", m.SelectedAction(), "remove")
	}

	m = m.CursorDown()
	if m.SelectedAction() != "list" {
		t.Errorf("after two down, action = %q, want %q", m.SelectedAction(), "list")
	}
}

func TestMenuModel_CursorUp(t *testing.T) {
	m := NewMenuModel(testData())

	m = m.CursorDown() // remove
	m = m.CursorDown() // list
	m = m.CursorUp()   // back to remove
	if m.SelectedAction() != "remove" {
		t.Errorf("after up, action = %q, want %q", m.SelectedAction(), "remove")
	}
}

func TestMenuModel_CursorSkipsSectionHeaders(t *testing.T) {
	m := NewMenuModel(testData())

	// Move down through all 6 actions. Next should be the first worktree
	// (skipping the WORKTREES section header).
	for range 5 {
		m = m.CursorDown()
	}
	// Now on "dir" (last action)
	if m.SelectedAction() != "dir" {
		t.Errorf("on last action, got %q, want %q", m.SelectedAction(), "dir")
	}

	m = m.CursorDown() // Should skip WORKTREES header, land on "main"
	if m.SelectedAction() != "navigate:main" {
		t.Errorf("after crossing section, action = %q, want %q", m.SelectedAction(), "navigate:main")
	}
}

func TestMenuModel_CursorWraps(t *testing.T) {
	m := NewMenuModel(testData())

	// Move up from the first selectable item — should wrap to the last.
	m = m.CursorUp()
	if m.SelectedAction() != "navigate:feature-x" {
		t.Errorf("wrap up, action = %q, want %q", m.SelectedAction(), "navigate:feature-x")
	}

	// Move down from the last should wrap to the first.
	m = m.CursorDown()
	if m.SelectedAction() != "setup" {
		t.Errorf("wrap down, action = %q, want %q", m.SelectedAction(), "setup")
	}
}

func TestMenuModel_NoWorktrees(t *testing.T) {
	d := data.TUIData{Root: "/tmp/repo", Worktrees: nil}
	m := NewMenuModel(d)

	if m.SelectedAction() != "setup" {
		t.Errorf("no worktrees, initial action = %q, want %q", m.SelectedAction(), "setup")
	}
}

func TestMenuModel_ViewContainsCurrentMarker(t *testing.T) {
	m := NewMenuModel(testData())
	view := m.View(40)
	if !strings.Contains(view, "(current)") {
		t.Error("View should contain '(current)' marker for current worktree")
	}
}

func TestMenuModel_ViewContainsHighlightPrefix(t *testing.T) {
	m := NewMenuModel(testData())
	view := m.View(40)
	if !strings.Contains(view, "> ") {
		t.Error("View should contain '> ' cursor prefix")
	}
}
