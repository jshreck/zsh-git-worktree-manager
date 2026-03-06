# Git Worktree Manager

An Oh-My-Zsh plugin for managing git worktrees with ease. Streamline your workflow when working with multiple branches simultaneously.

## 🌟 Features

- **worktree dir** - Initialize a new bare repository with worktree structure from any git URL
- **worktree setup** - Create a new worktree with automatic branch creation, config file copying, and dependency installation
- **worktree list** - List all worktrees with pretty formatting
- **worktree remove** - Remove a worktree with confirmation and optional branch deletion
  - ✨ **Interactive menu** when called without arguments
  - ⌨️ **Tab completion** for worktree names
- **worktree pull** - Pull latest changes with automatic stash/unstash
- **wt** - Quick navigation between worktrees
  - ⌨️ **Tab completion** for worktree names
- **worktree review** - Create a worktree from a GitHub PR for easy code review
- **wtp** - Alias for `worktree pull`

## 📋 Prerequisites

- [Oh-My-Zsh](https://ohmyz.sh/) installed
- Git 2.15+ (for worktree support)
- [GitHub CLI (gh)](https://cli.github.com/) (required for `worktree review` command)
- Yarn (for dependency installation)

## 🚀 Installation

### Option 1: Clone and Symlink (Recommended)

```bash
# Clone the repository
cd ~/GitHub
git clone git@github.com:tmbtech/zsh-git-worktree-manager.git

# Create symlink to Oh-My-Zsh custom plugins directory
ln -s ~/GitHub/zsh-git-worktree-manager ~/.oh-my-zsh/custom/plugins/git-worktree-manager

# Add to your ~/.zshrc plugins array
# Open ~/.zshrc and find the plugins array, then add:
plugins=(git zsh-autosuggestions brew git-worktree-manager)

# Reload your shell
source ~/.zshrc
```

### Option 2: Use the Installation Script

```bash
cd ~/GitHub/zsh-git-worktree-manager
./install.sh
```

Then follow the instructions to add `git-worktree-manager` to your plugins array in `~/.zshrc`.

## 📖 Usage

### worktree dir

Initialize a new repository with bare repository worktree structure:

```bash
# Auto-generate directory name (creates {repo-name}-worktrees)
worktree dir git@github.com:user/my-repo.git
# Creates: ./my-repo-worktrees/

# Use custom directory name
worktree dir git@github.com:user/my-repo.git custom-project
# Creates: ./custom-project/

# Works with HTTPS URLs too
worktree dir https://github.com/user/another-repo.git

# Show help
worktree dir --help
```

**What it does:**
1. Creates parent directory in your current location
2. Clones repository as a bare repo into `.bare/`
3. Configures `.git` pointer file
4. Sets up fetch/pull configuration
5. Creates initial `main` worktree

**After creation:**
- Navigate to the main worktree: `cd {dir-name}/main`
- Copy your config files (*.env, certs) into the main directory
- Run `yarn install` or `npm install`
- Create additional worktrees with `worktree setup`

### worktree setup

Create a new worktree with automatic configuration:

```bash
# Basic usage - creates a new worktree from main branch
worktree setup feature/my-feature

# Skip yarn install
worktree setup --skip-yarn bugfix/critical-fix

# Create from a different base branch
worktree setup --base=develop feature/new-thing

# Show help
worktree setup --help
```

**What it does:**
1. Updates the base branch (default: main)
2. Creates a new worktree and branch
3. Copies environment files (*.env, key.pem, cert.pem)
4. Installs dependencies with yarn
5. Navigates you to the new worktree

### worktree list

Display all worktrees in a formatted list:

```bash
worktree list
```

### worktree remove

Remove a worktree with confirmation:

```bash
# Remove a specific worktree
worktree remove feature/my-feature

# Interactive mode - select from a menu (no arguments)
worktree remove

# Tab completion - press TAB to see available worktrees
worktree remove <TAB>
```

**What it does:**
1. Shows interactive selection menu if no worktree name provided
2. Confirms worktree removal
3. Removes the worktree
4. Optionally deletes the associated branch
5. Handles force deletion if branch has unmerged changes

**Interactive Mode:**
When called without arguments, displays a numbered menu of all removable worktrees (excluding the protected "main" worktree). Use arrow keys or type the number to select, then press Enter. Press Ctrl+C to cancel.

### worktree pull

Pull latest changes in the current worktree:

```bash
worktree pull

# Or use the alias
wtp
```

**What it does:**
1. Checks for uncommitted changes and stashes them
2. Fetches from remote
3. Pulls changes
4. Restores stashed changes

### wt

Quick navigation to worktree root or specific worktree:

```bash
# Navigate to worktree root
wt

# Navigate to specific worktree
wt feature/my-feature

# Tab completion - press TAB to see all available worktrees
wt <TAB>
```

### worktree review

Create a worktree from a GitHub PR for code review:

```bash
worktree review https://github.com/owner/repo/pull/123
```

**What it does:**
1. Fetches PR information using gh CLI
2. Creates a worktree with name `pr-<branch-name>`
3. Sets up upstream tracking
4. Copies environment files
5. Installs dependencies
6. Navigates you to the PR worktree

## 🔧 Configuration

This plugin automatically:
- Detects worktree root directory
- Uses `yarn` for dependency installation (configurable)
- Skips Husky hooks during worktree creation
- Copies `.env` files, `key.pem`, and `cert.pem` to new worktrees

## 🐛 Troubleshooting

### "Could not auto-detect worktree root"

This warning appears when the plugin cannot find your worktree structure. Make sure:
- You're running commands from within a worktree
- You have a `main` worktree directory

### "GitHub CLI (gh) is not installed"

For the `worktree review` command, you need the GitHub CLI:

```bash
brew install gh
gh auth login
```

### Functions not found after installation

Make sure you:
1. Added `git-worktree-manager` to the plugins array in `~/.zshrc`
2. Reloaded your shell with `source ~/.zshrc`

### yarn install fails

You can skip yarn installation with:

```bash
worktree setup --skip-yarn my-branch
```

Then manually run `yarn install` when ready.

## 🤝 Contributing

Contributions are welcome! Here's how you can help:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes using [Commitizen format](https://www.conventionalcommits.org/)
4. Push to your branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Commit Format

We use [Conventional Commits](https://www.conventionalcommits.org/) format:

```
<type>(<scope>): <description>

Example:
feat(worktree): add support for custom config files
fix(pr-review): handle edge case with branch names
docs: update installation instructions
```

## 📝 License

MIT License - see [LICENSE](LICENSE) file for details

## 🙏 Acknowledgments

Created to streamline worktree workflows for development teams working with monorepos and feature branches.

## 📬 Support

If you encounter issues or have questions:
1. Check the [Troubleshooting](#-troubleshooting) section
2. Review existing GitHub Issues
3. Open a new issue with detailed information about your environment and the problem

---

**Made with ❤️ for efficient git workflows**
