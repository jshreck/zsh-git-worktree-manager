package model

import (
	"fmt"
	"strings"

	"github.com/tmbtech/zsh-git-worktree-manager/tui/internal/data"
	"github.com/tmbtech/zsh-git-worktree-manager/tui/internal/style"
)

// ContentModel renders the right-column context panel.
type ContentModel struct {
	tuiData data.TUIData
}

// NewContentModel creates a new ContentModel.
func NewContentModel(d data.TUIData) ContentModel {
	return ContentModel{tuiData: d}
}

// View renders content appropriate for the highlighted menu action.
func (c ContentModel) View(action string, width int) string {
	if strings.HasPrefix(action, "navigate:") {
		name := strings.TrimPrefix(action, "navigate:")
		return c.renderWorktreeDetail(name, width)
	}

	switch action {
	case "setup":
		return c.renderSetup(width)
	case "remove":
		return c.renderRemove(width)
	case "list":
		return c.renderList(width)
	case "pull":
		return c.renderPull(width)
	case "review":
		return c.renderReview(width)
	case "dir":
		return c.renderDir(width)
	default:
		return style.ContentMuted.Width(width).Render("Select an action from the menu.")
	}
}

func (c ContentModel) renderSetup(width int) string {
	var b strings.Builder
	b.WriteString(style.ContentTitle.Width(width).Render("Setup New Worktree"))
	b.WriteString("\n\n")
	b.WriteString(style.ContentBody.Width(width).Render(
		"Creates a new git worktree from a base branch (default: main)."))
	b.WriteString("\n\n")
	b.WriteString(style.ContentMuted.Width(width).Render("Usage:"))
	b.WriteString("\n")
	b.WriteString(style.ContentBody.Width(width).Render(
		"  worktree setup <branch-name>\n  worktree setup --base=dev <name>"))
	b.WriteString("\n\n")
	b.WriteString(style.ContentMuted.Width(width).Render("Options:"))
	b.WriteString("\n")
	b.WriteString(style.ContentBody.Width(width).Render(
		"  --skip-yarn   Skip yarn install\n  --base=NAME   Use alternate base branch"))
	return b.String()
}

func (c ContentModel) renderRemove(width int) string {
	var b strings.Builder
	b.WriteString(style.ContentTitle.Width(width).Render("Remove Worktree"))
	b.WriteString("\n\n")
	b.WriteString(style.ContentBody.Width(width).Render(
		"Remove a worktree and optionally delete its branch."))
	b.WriteString("\n\n")

	removable := c.removableWorktrees()
	if len(removable) == 0 {
		b.WriteString(style.ContentMuted.Width(width).Render("No removable worktrees (only main exists)."))
	} else {
		b.WriteString(style.ContentMuted.Width(width).Render("Removable worktrees:"))
		b.WriteString("\n")
		for _, wt := range removable {
			line := fmt.Sprintf("  %s  %s", wt.Name, wt.ShortHead())
			b.WriteString(style.ContentBody.Width(width).Render(line))
			b.WriteString("\n")
		}
	}
	return b.String()
}

func (c ContentModel) renderList(width int) string {
	var b strings.Builder
	b.WriteString(style.ContentTitle.Width(width).Render("Worktree List"))
	b.WriteString("\n\n")

	if len(c.tuiData.Worktrees) == 0 {
		b.WriteString(style.ContentMuted.Width(width).Render("No worktrees found."))
		return b.String()
	}

	for _, wt := range c.tuiData.Worktrees {
		marker := "  "
		if wt.IsCurrent {
			marker = "* "
		}
		line := fmt.Sprintf("%s%-20s %s  %s", marker, wt.Name, wt.ShortHead(), wt.Branch)
		b.WriteString(style.ContentBody.Width(width).Render(line))
		b.WriteString("\n")
	}
	return b.String()
}

func (c ContentModel) renderPull(width int) string {
	var b strings.Builder
	b.WriteString(style.ContentTitle.Width(width).Render("Pull Latest"))
	b.WriteString("\n\n")
	b.WriteString(style.ContentBody.Width(width).Render(
		"Pull latest changes in the current worktree with auto-stash."))
	b.WriteString("\n\n")
	b.WriteString(style.ContentMuted.Width(width).Render(
		"Automatically stashes uncommitted changes before pulling\nand restores them afterward."))

	if c.tuiData.InWorktree {
		b.WriteString("\n\n")
		current := c.currentWorktree()
		if current != nil {
			b.WriteString(style.ContentBody.Width(width).Render(
				fmt.Sprintf("Current: %s (%s)", current.Branch, current.ShortHead())))
		}
	}
	return b.String()
}

func (c ContentModel) renderReview(width int) string {
	var b strings.Builder
	b.WriteString(style.ContentTitle.Width(width).Render("Review PR"))
	b.WriteString("\n\n")
	b.WriteString(style.ContentBody.Width(width).Render(
		"Create a worktree from a GitHub PR for code review."))
	b.WriteString("\n\n")
	b.WriteString(style.ContentMuted.Width(width).Render("Usage:"))
	b.WriteString("\n")
	b.WriteString(style.ContentBody.Width(width).Render(
		"  worktree review <github-pr-url>"))

	if !c.tuiData.HasGH {
		b.WriteString("\n\n")
		b.WriteString(style.ContentMuted.Width(width).Render(
			"Warning: GitHub CLI (gh) not found.\nInstall it with: brew install gh"))
	}
	return b.String()
}

func (c ContentModel) renderDir(width int) string {
	var b strings.Builder
	b.WriteString(style.ContentTitle.Width(width).Render("Init Directory"))
	b.WriteString("\n\n")
	b.WriteString(style.ContentBody.Width(width).Render(
		"Initialize a new bare repository with worktree structure."))
	b.WriteString("\n\n")
	b.WriteString(style.ContentMuted.Width(width).Render("Usage:"))
	b.WriteString("\n")
	b.WriteString(style.ContentBody.Width(width).Render(
		"  worktree dir <git-url> [dir-name]"))
	b.WriteString("\n\n")
	b.WriteString(style.ContentMuted.Width(width).Render(
		"Creates a .bare/ clone, sets up fetch refs,\nand creates an initial main worktree."))
	return b.String()
}

func (c ContentModel) renderWorktreeDetail(name string, width int) string {
	var b strings.Builder

	var wt *data.Worktree
	for i := range c.tuiData.Worktrees {
		if c.tuiData.Worktrees[i].Name == name {
			wt = &c.tuiData.Worktrees[i]
			break
		}
	}

	if wt == nil {
		b.WriteString(style.ContentMuted.Width(width).Render("Worktree not found."))
		return b.String()
	}

	b.WriteString(style.ContentTitle.Width(width).Render(wt.Name))
	b.WriteString("\n\n")

	b.WriteString(style.ContentMuted.Width(width).Render("Branch:"))
	b.WriteString("\n")
	b.WriteString(style.ContentBody.Width(width).Render("  " + wt.Branch))
	b.WriteString("\n\n")

	b.WriteString(style.ContentMuted.Width(width).Render("HEAD:"))
	b.WriteString("\n")
	b.WriteString(style.ContentBody.Width(width).Render("  " + wt.ShortHead()))
	b.WriteString("\n\n")

	b.WriteString(style.ContentMuted.Width(width).Render("Path:"))
	b.WriteString("\n")
	b.WriteString(style.ContentBody.Width(width).Render("  " + wt.Path))

	if wt.IsCurrent {
		b.WriteString("\n\n")
		b.WriteString(style.ContentBody.Width(width).Render("(current worktree)"))
	}

	b.WriteString("\n\n")
	b.WriteString(style.ContentMuted.Width(width).Render("Press Enter to navigate to this worktree."))
	return b.String()
}

func (c ContentModel) removableWorktrees() []data.Worktree {
	var result []data.Worktree
	rootPath := c.tuiData.Root
	for _, wt := range c.tuiData.Worktrees {
		if wt.IsBare || wt.Path == rootPath {
			continue
		}
		result = append(result, wt)
	}
	return result
}

func (c ContentModel) currentWorktree() *data.Worktree {
	for i := range c.tuiData.Worktrees {
		if c.tuiData.Worktrees[i].IsCurrent {
			return &c.tuiData.Worktrees[i]
		}
	}
	return nil
}
