package model

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestModel_Init(t *testing.T) {
	m := New(testData())
	cmd := m.Init()
	if cmd != nil {
		t.Error("Init should return nil cmd")
	}
}

func TestModel_QuitOnQ(t *testing.T) {
	m := New(testData())
	updated, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	if cmd == nil {
		t.Fatal("expected tea.Quit cmd on 'q' press")
	}
	result := updated.(Model)
	if result.Selected() != "" {
		t.Errorf("quit should produce empty selection, got %q", result.Selected())
	}
}

func TestModel_QuitOnCtrlC(t *testing.T) {
	m := New(testData())
	updated, cmd := m.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	if cmd == nil {
		t.Fatal("expected tea.Quit cmd on Ctrl+C")
	}
	result := updated.(Model)
	if result.Selected() != "" {
		t.Errorf("quit should produce empty selection, got %q", result.Selected())
	}
}

func TestModel_EnterSelectsAction(t *testing.T) {
	m := New(testData())
	// Initial selection is "setup" (first action).
	updated, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("expected tea.Quit cmd on Enter")
	}
	result := updated.(Model)
	if result.Selected() != "setup" {
		t.Errorf("Selected() = %q, want %q", result.Selected(), "setup")
	}
}

func TestModel_NavigateAndSelect(t *testing.T) {
	m := New(testData())

	// Move down once to "remove".
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})
	m = updated.(Model)

	// Press enter.
	updated, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("expected tea.Quit cmd on Enter")
	}
	result := updated.(Model)
	if result.Selected() != "remove" {
		t.Errorf("Selected() = %q, want %q", result.Selected(), "remove")
	}
}

func TestModel_WindowResize(t *testing.T) {
	m := New(testData())
	updated, _ := m.Update(tea.WindowSizeMsg{Width: 120, Height: 40})
	result := updated.(Model)

	// Should not crash and should produce a non-empty view.
	view := result.View()
	if view == "" {
		t.Error("View should not be empty after resize")
	}
}

func TestModel_VimKeys(t *testing.T) {
	m := New(testData())

	// 'j' should move down like down arrow.
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}})
	m = updated.(Model)
	updated, _ = m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	result := updated.(Model)
	if result.Selected() != "remove" {
		t.Errorf("after 'j', Selected() = %q, want %q", result.Selected(), "remove")
	}
}

func TestModel_ViewHeader(t *testing.T) {
	m := New(testData())
	// Set dimensions so view renders.
	updated, _ := m.Update(tea.WindowSizeMsg{Width: 100, Height: 30})
	result := updated.(Model)
	view := result.View()
	if view == "" {
		t.Error("View should not be empty")
	}
}

func TestModel_ViewQuitClearsScreen(t *testing.T) {
	m := New(testData())
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	result := updated.(Model)
	view := result.View()
	if view != "" {
		t.Errorf("View after quit should be empty, got %q", view)
	}
}
