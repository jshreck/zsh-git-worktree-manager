package data

import (
	"strings"
	"testing"
)

func TestParseFromReader_ValidJSON(t *testing.T) {
	input := `{
		"root": "/tmp/repo",
		"current_dir": "/tmp/repo/main",
		"repo_name": "repo",
		"in_worktree": true,
		"has_gh": true,
		"worktrees": [
			{
				"name": "main",
				"path": "/tmp/repo/main",
				"branch": "main",
				"head": "abc1234567890",
				"is_bare": false,
				"is_current": true
			},
			{
				"name": "feature-foo",
				"path": "/tmp/repo/feature-foo",
				"branch": "feature/foo",
				"head": "def5678901234",
				"is_bare": false,
				"is_current": false
			}
		]
	}`

	d, err := ParseFromReader(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if d.Root != "/tmp/repo" {
		t.Errorf("Root = %q, want %q", d.Root, "/tmp/repo")
	}
	if d.RepoName != "repo" {
		t.Errorf("RepoName = %q, want %q", d.RepoName, "repo")
	}
	if !d.InWorktree {
		t.Error("InWorktree = false, want true")
	}
	if !d.HasGH {
		t.Error("HasGH = false, want true")
	}
	if len(d.Worktrees) != 2 {
		t.Fatalf("Worktrees len = %d, want 2", len(d.Worktrees))
	}
	if d.Worktrees[0].Name != "main" {
		t.Errorf("Worktrees[0].Name = %q, want %q", d.Worktrees[0].Name, "main")
	}
	if !d.Worktrees[0].IsCurrent {
		t.Error("Worktrees[0].IsCurrent = false, want true")
	}
	if d.Worktrees[1].Branch != "feature/foo" {
		t.Errorf("Worktrees[1].Branch = %q, want %q", d.Worktrees[1].Branch, "feature/foo")
	}
}

func TestParseFromReader_EmptyWorktrees(t *testing.T) {
	input := `{"root":"/tmp/repo","current_dir":"","repo_name":"repo","in_worktree":false,"has_gh":false,"worktrees":[]}`
	d, err := ParseFromReader(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(d.Worktrees) != 0 {
		t.Errorf("Worktrees len = %d, want 0", len(d.Worktrees))
	}
}

func TestParseFromReader_MissingRoot(t *testing.T) {
	input := `{"current_dir":"","worktrees":[]}`
	_, err := ParseFromReader(strings.NewReader(input))
	if err == nil {
		t.Fatal("expected error for missing root, got nil")
	}
}

func TestParseFromReader_MalformedJSON(t *testing.T) {
	input := `{not valid json`
	_, err := ParseFromReader(strings.NewReader(input))
	if err == nil {
		t.Fatal("expected error for malformed JSON, got nil")
	}
}

func TestParseFromReader_UnknownField(t *testing.T) {
	input := `{"root":"/tmp","unknown_field":true,"worktrees":[]}`
	_, err := ParseFromReader(strings.NewReader(input))
	if err == nil {
		t.Fatal("expected error for unknown field, got nil")
	}
}

func TestShortHead(t *testing.T) {
	tests := []struct {
		head string
		want string
	}{
		{"abc1234567890", "abc1234"},
		{"short", "short"},
		{"ab", "ab"},
		{"", ""},
		{"1234567", "1234567"},
	}
	for _, tt := range tests {
		wt := Worktree{Head: tt.head}
		got := wt.ShortHead()
		if got != tt.want {
			t.Errorf("ShortHead(%q) = %q, want %q", tt.head, got, tt.want)
		}
	}
}
