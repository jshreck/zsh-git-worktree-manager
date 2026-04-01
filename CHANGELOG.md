# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed
- **Migrated to `worktree <subcommand>` UX** - all standalone commands are now subcommands of `worktree`
  - `worktree setup` (was `worktree_setup`)
  - `worktree remove` (was `worktree_remove`)
  - `worktree list` (was `worktree_list`)
  - `worktree pull` (was `worktree_pull`)
  - `worktree dir` (was `worktree_dir`)
  - `worktree review` (was `pr_review`)
  - `wtp` alias now runs `worktree pull`
  - Old standalone names removed entirely (clean break, no aliases)
- Renamed `worktree pr` subcommand to `worktree review` for clarity
- **Updated tab completion** to support `worktree <TAB>` (lists all subcommands) and `worktree remove <TAB>` (lists worktree names)

### Added
- **Zsh completion system** for enhanced user experience
  - Tab completion for `worktree remove` command (excludes protected "main" worktree)
  - Tab completion for `wt` command (includes all worktrees)
  - Dynamic worktree name suggestions from current git repository
  - Integration with Oh-My-Zsh completion system
- **Interactive selection menu** for `worktree remove`
  - Shows numbered menu when called without arguments
  - Displays all removable worktrees (excluding "main")
  - Supports cancellation with Ctrl+C
  - Seamlessly integrates with existing confirmation workflow
- `worktree dir` - Initialize a new bare repository with worktree structure
  - Auto-generates directory name with `-worktrees` suffix from repository URL
  - Supports custom directory naming
  - Works with SSH and HTTPS git URLs
  - Creates bare repository structure (`.bare/` directory)
  - Configures `.git` pointer file
  - Sets up fetch/pull configuration
  - Creates initial `main` worktree
  - Comprehensive error handling with automatic cleanup on failure
  - Repository-agnostic design works with any git repository

## [1.0.0] - 2025-01-07

### Added
- Initial release with core worktree management functions
- `worktree_setup` - Create new worktrees with automatic configuration
  - Branch creation from base branch (default: main)
  - Automatic copying of environment files (*.env, key.pem, cert.pem)
  - Automatic dependency installation with yarn
  - Options for skipping yarn install and custom base branch
- `worktree_list` - List all git worktrees with pretty formatting
- `worktree_remove` - Remove worktrees with confirmation prompts
  - Optional branch deletion after worktree removal
  - Force deletion option for unmerged branches
  - Protection against deleting main worktree
- `worktree_pull` - Pull latest changes in current worktree
  - Automatic stash/unstash of uncommitted changes
  - Detached HEAD state detection
  - Remote tracking branch validation
- `wt` - Quick navigation between worktrees
  - Navigate to worktree root or specific worktree by name
- `pr_review` - Create worktrees from GitHub PRs for code review
  - GitHub CLI integration for PR information
  - Automatic worktree and branch setup
  - Upstream tracking configuration
  - Environment file copying and dependency installation
- `wtp` - Alias for `worktree_pull`
- Comprehensive README with installation instructions and usage examples
- MIT License
- Installation script with prerequisite checking
- Oh-My-Zsh plugin integration

### Features
- Automatic worktree root detection
- Husky hooks skipping during worktree creation
- Support for yarn dependency management
- Colorful emoji-based CLI output
- Error handling and user-friendly messages

[Unreleased]: https://github.com/tmbtech/zsh-git-worktree-manager/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/tmbtech/zsh-git-worktree-manager/releases/tag/v1.0.0
