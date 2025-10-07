# Git Worktree Manager

An Oh-My-Zsh plugin for managing git worktrees with ease. Streamline your workflow when working with multiple branches simultaneously.

## 🌟 Features

- **worktree_setup** - Create a new worktree with automatic branch creation, config file copying, and dependency installation
- **worktree_list** - List all worktrees with pretty formatting
- **worktree_remove** - Remove a worktree with confirmation and optional branch deletion
- **worktree_pull** - Pull latest changes with automatic stash/unstash
- **wt** - Quick navigation between worktrees
- **pr_review** - Create a worktree from a GitHub PR for easy code review
- **wtp** - Alias for `worktree_pull`

## 📋 Prerequisites

- [Oh-My-Zsh](https://ohmyz.sh/) installed
- Git 2.15+ (for worktree support)
- [GitHub CLI (gh)](https://cli.github.com/) (required for `pr_review` function)
- Yarn (for dependency installation)

## 🚀 Installation

### Option 1: Clone and Symlink (Recommended)

```bash
# Clone the repository
cd ~/GitHub
git clone git@github.com:YOUR-USERNAME/zsh-git-worktree-manager.git

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

### worktree_setup

Create a new worktree with automatic configuration:

```bash
# Basic usage - creates a new worktree from main branch
worktree_setup feature/my-feature

# Skip yarn install
worktree_setup --skip-yarn bugfix/critical-fix

# Create from a different base branch
worktree_setup --base=develop feature/new-thing

# Show help
worktree_setup --help
```

**What it does:**
1. Updates the base branch (default: main)
2. Creates a new worktree and branch
3. Copies environment files (*.env, key.pem, cert.pem)
4. Installs dependencies with yarn
5. Navigates you to the new worktree

### worktree_list

Display all worktrees in a formatted list:

```bash
worktree_list
```

### worktree_remove

Remove a worktree with confirmation:

```bash
worktree_remove feature/my-feature
```

**What it does:**
1. Confirms worktree removal
2. Removes the worktree
3. Optionally deletes the associated branch
4. Handles force deletion if branch has unmerged changes

### worktree_pull

Pull latest changes in the current worktree:

```bash
worktree_pull

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
```

### pr_review

Create a worktree from a GitHub PR for code review:

```bash
pr_review https://github.com/owner/repo/pull/123
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

For the `pr_review` function, you need the GitHub CLI:

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
worktree_setup --skip-yarn my-branch
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

We use Commitizen format for commits:

```
<type>(<scope>): <description> <ticket-reference>

Example:
feat(worktree): add support for custom config files PLTW-123
fix(pr-review): handle edge case with branch names PLTW-456
docs: update installation instructions PLTW-000
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