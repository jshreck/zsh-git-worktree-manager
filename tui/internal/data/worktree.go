package data

import (
	"encoding/json"
	"fmt"
	"io"
)

// Worktree represents a single git worktree.
type Worktree struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	Branch    string `json:"branch"`
	Head      string `json:"head"`
	IsBare    bool   `json:"is_bare"`
	IsCurrent bool   `json:"is_current"`
}

// TUIData is the top-level JSON contract between the Zsh data collector and
// the Go TUI binary.
type TUIData struct {
	Root        string     `json:"root"`
	CurrentDir  string     `json:"current_dir"`
	Worktrees   []Worktree `json:"worktrees"`
	InWorktree  bool       `json:"in_worktree"`
	HasGH       bool       `json:"has_gh"`
	RepoName    string     `json:"repo_name"`
}

// ShortHead returns an abbreviated HEAD hash (first 7 chars).
func (w Worktree) ShortHead() string {
	if len(w.Head) >= 7 {
		return w.Head[:7]
	}
	return w.Head
}

// ParseFromReader decodes TUIData JSON from the supplied reader.
func ParseFromReader(r io.Reader) (TUIData, error) {
	var d TUIData
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&d); err != nil {
		return TUIData{}, fmt.Errorf("failed to parse TUI data: %w", err)
	}
	if d.Root == "" {
		return TUIData{}, fmt.Errorf("missing required field: root")
	}
	return d, nil
}
