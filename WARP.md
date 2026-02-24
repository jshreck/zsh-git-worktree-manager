# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

## Project Overview

This is an Oh-My-Zsh plugin that provides comprehensive git worktree management functions. The plugin is written in Zsh and provides commands for creating, removing, listing, and navigating between git worktrees, plus a special command for reviewing GitHub PRs in isolated worktrees.

## Architecture

### Plugin Structure
- **git-worktree-manager.plugin.zsh** - Main plugin entry point that sources all function files and creates aliases
- **functions/** - Directory containing individual function files (one per command)
- **completions/** - Directory containing zsh completion definitions for autocomplete functionality
- **install.sh** - Installation script that creates symlink to Oh-My-Zsh custom plugins directory

### Completion System
The plugin includes a zsh completion system for enhanced user experience:

**Location:** `completions/_git_worktree_manager`

**Features:**
- Tab completion for `worktree` - lists all six subcommands with descriptions
- Tab completion for `worktree remove` - dynamically lists all removable worktrees (excludes "main")
- Tab completion for `wt` - lists all available worktrees for navigation
- Integration with Oh-My-Zsh's completion system via fpath

**Implementation Details:**
- Uses `#compdef` directive to register completion functions
- `_get_worktrees()` helper function parses `git worktree list --porcelain` output
- Filters worktree names using zsh array manipulation (e.g., `${worktrees:#main}` to exclude "main")
- Provides contextual completion based on the command being used

**Interactive Selection Menu:**
The `worktree remove` subcommand includes an interactive mode when called without arguments:
- Uses zsh's built-in `select` command to display a numbered menu
- Filters out "main" worktree (protected from deletion)
- Handles user cancellation gracefully (Ctrl+C)
- Seamlessly transitions to existing confirmation workflow after selection

### Core Functions
The plugin exposes a single `worktree` entry-point with subcommands. Each internal implementation is in its own file in the `functions/` directory:

1. **_get_worktree_root** - Helper that auto-detects the worktree root directory by examining git metadata
2. **worktree** - Dispatcher that routes subcommands to internal `_wt_*` functions
3. **_wt_setup** - Creates new worktrees with automatic branch creation, config copying, and yarn install
4. **_wt_remove** - Removes worktrees with confirmation prompts and optional branch deletion
5. **_wt_list** - Pretty-prints all worktrees
6. **_wt_pull** - Pulls changes with automatic stash/unstash handling
7. **_wt_dir** - Initializes bare repo directory structure
8. **wt** - Quick navigation utility (no args = root, or specify worktree name)
9. **_wt_pr** - Creates worktrees from GitHub PRs using gh CLI

### Key Design Patterns
- All functions use `_get_worktree_root` to dynamically detect the worktree parent directory
- Functions expect a "main" worktree to exist as the base reference point
- `HUSKY_SKIP_HOOKS=1` is used during worktree creation to avoid errors before dependencies are installed
- Configuration files (*.env, key.pem, cert.pem) are automatically copied using tar to preserve directory structure
- Interactive confirmations use Zsh's `read -r` for safety operations

## Development Commands

### Testing the Plugin
```bash
# Source the plugin manually for testing
source git-worktree-manager.plugin.zsh

# Test individual subcommands
worktree setup --help
worktree list
```

### Installing/Updating the Plugin
```bash
# Run the installation script
./install.sh

# Or manually create symlink
ln -sf ~/GitHub/zsh-git-worktree-manager ~/.oh-my-zsh/custom/plugins/git-worktree-manager

# Reload shell to test changes
source ~/.zshrc
```

### Testing Individual Functions
To test a single function in isolation:
```bash
# Source the helper first
source functions/_get_worktree_root

# Then source the function you want to test
source functions/_wt_setup

# Test it
worktree setup --help
```

### Testing Completions
After modifying completion files:
```bash
# Reload your shell configuration
source ~/.zshrc

# Or manually reload completions
autoload -U compinit && compinit

# Test tab completion
worktree remove <TAB>
wt <TAB>

# Test interactive menu
worktree remove
```

## Important Considerations

### Worktree Directory Structure
The plugin expects a specific directory structure:
```
parent-directory/
├── main/           # Required: The "main" worktree
├── feature-1/      # Additional worktrees created by plugin
├── pr-feature-2/   # PR review worktrees prefixed with "pr-"
└── ...
```

### Configuration File Handling
The plugin automatically copies these files from the main worktree to new worktrees:
- `*.env` files (excluding node_modules)
- `key.pem`
- `cert.pem`

Uses `find` with `-prune` to exclude node_modules, then `tar` for atomic copying with directory structure preservation.

### Dependencies
- Git 2.15+ (for worktree support)
- Yarn (for dependency installation)
- GitHub CLI (`gh`) - required only for `worktree review` command
- Oh-My-Zsh framework

### Commit Format
When contributing, use Commitizen format as specified in the README:
```
<type>(<scope>): <description> <ticket-reference>

Example:
feat(worktree): add support for custom config files PLTW-123
```

### Protected Operations
- The "main" worktree cannot be removed (hardcoded protection in `worktree remove`)
- Branch deletion requires confirmation and offers force-delete for unmerged changes
- Worktree removal always prompts for confirmation

## Common Workflows

### Adding a New Function
1. Create new file in `functions/` directory named `_wt_<name>`
2. Define the function using `function _wt_<name>() { ... }` syntax
3. Add a case entry in `functions/worktree` dispatcher
4. Add emoji-based output for consistency (🌳, ✅, ❌, ⚠️, 📂, 🌿, etc.)
5. Test by sourcing the plugin file
6. Update README.md with usage documentation
7. Update CHANGELOG.md following Keep a Changelog format

### Modifying Existing Functions
All functions follow similar patterns for error handling and user feedback. Maintain consistency:
- Use `echo "❌ Error: ..."` for errors with `return 1`
- Use `echo "✅ ..."` for success messages
- Use `echo "⚠️ Warning: ..."` for warnings
- Use decorative separator lines: `echo "━━━━━━...━━━━━━"`
- Check prerequisites at function start
- Store and restore `original_dir` when changing directories
