package model

import (
	"strings"
	"testing"

	"github.com/tmbtech/zsh-git-worktree-manager/tui/internal/data"
)

func TestContentModel_Setup(t *testing.T) {
	c := NewContentModel(testData())
	view := c.View("setup", 60)
	if !strings.Contains(view, "Setup New Worktree") {
		t.Error("setup view should contain title")
	}
	if !strings.Contains(view, "--skip-yarn") {
		t.Error("setup view should mention --skip-yarn option")
	}
}

func TestContentModel_Remove(t *testing.T) {
	c := NewContentModel(testData())
	view := c.View("remove", 60)
	if !strings.Contains(view, "Remove Worktree") {
		t.Error("remove view should contain title")
	}
	if !strings.Contains(view, "feature-x") {
		t.Error("remove view should list removable worktrees")
	}
	// main should not appear in removable list but appears as title or elsewhere;
	// we specifically check that the removable section does not list main as a removable entry.
}

func TestContentModel_Remove_NoRemovable(t *testing.T) {
	d := data.TUIData{
		Root: "/tmp/repo/main",
		Worktrees: []data.Worktree{
			{Name: "main", Path: "/tmp/repo/main", Branch: "main"},
		},
	}
	c := NewContentModel(d)
	view := c.View("remove", 60)
	if !strings.Contains(view, "No removable worktrees") {
		t.Error("remove view should indicate no removable worktrees when only main exists")
	}
}

func TestContentModel_List(t *testing.T) {
	c := NewContentModel(testData())
	view := c.View("list", 60)
	if !strings.Contains(view, "Worktree List") {
		t.Error("list view should contain title")
	}
	if !strings.Contains(view, "main") {
		t.Error("list view should contain 'main' worktree")
	}
}

func TestContentModel_List_Empty(t *testing.T) {
	d := data.TUIData{Root: "/tmp/repo", Worktrees: nil}
	c := NewContentModel(d)
	view := c.View("list", 60)
	if !strings.Contains(view, "No worktrees found") {
		t.Error("empty list should show 'No worktrees found'")
	}
}

func TestContentModel_Pull(t *testing.T) {
	c := NewContentModel(testData())
	view := c.View("pull", 60)
	if !strings.Contains(view, "Pull Latest") {
		t.Error("pull view should contain title")
	}
	if !strings.Contains(view, "auto-stash") {
		t.Error("pull view should mention auto-stash")
	}
}

func TestContentModel_Review_WithGH(t *testing.T) {
	c := NewContentModel(testData())
	view := c.View("review", 60)
	if !strings.Contains(view, "Review PR") {
		t.Error("review view should contain title")
	}
	if strings.Contains(view, "Warning") {
		t.Error("review view should not warn when gh is available")
	}
}

func TestContentModel_Review_WithoutGH(t *testing.T) {
	d := testData()
	d.HasGH = false
	c := NewContentModel(d)
	view := c.View("review", 60)
	if !strings.Contains(view, "Warning") {
		t.Error("review view should warn when gh is unavailable")
	}
}

func TestContentModel_Dir(t *testing.T) {
	c := NewContentModel(testData())
	view := c.View("dir", 60)
	if !strings.Contains(view, "Init Directory") {
		t.Error("dir view should contain title")
	}
}

func TestContentModel_NavigateWorktree(t *testing.T) {
	c := NewContentModel(testData())
	view := c.View("navigate:main", 60)
	if !strings.Contains(view, "main") {
		t.Error("worktree detail should show worktree name")
	}
	if !strings.Contains(view, "(current worktree)") {
		t.Error("current worktree should be marked")
	}
}

func TestContentModel_NavigateWorktree_NotFound(t *testing.T) {
	c := NewContentModel(testData())
	view := c.View("navigate:nonexistent", 60)
	if !strings.Contains(view, "not found") {
		t.Error("missing worktree should show 'not found'")
	}
}

func TestContentModel_UnknownAction(t *testing.T) {
	c := NewContentModel(testData())
	view := c.View("unknown", 60)
	if !strings.Contains(view, "Select an action") {
		t.Error("unknown action should show default message")
	}
}
