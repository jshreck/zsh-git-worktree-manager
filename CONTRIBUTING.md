# Contributing to Git Worktree Manager

Thanks for your interest in contributing! This guide will help you get started.

## Development Setup

1. **Fork & clone** the repository:

   ```bash
   git clone https://github.com/<your-username>/zsh-git-worktree-manager.git
   cd zsh-git-worktree-manager
   ```

2. **Symlink** the plugin into Oh-My-Zsh for local testing:

   ```bash
   ln -sf "$(pwd)" ~/.oh-my-zsh/custom/plugins/git-worktree-manager
   source ~/.zshrc
   ```

3. **Prerequisites**:
   - Zsh + [Oh-My-Zsh](https://ohmyz.sh/)
   - Git 2.15+
   - [GitHub CLI (`gh`)](https://cli.github.com/) (for `worktree review`)
   - [Go 1.21+](https://go.dev/) (for building the TUI binary)

## Project Structure

```
functions/          # Zsh functions (core plugin logic)
completions/        # Zsh tab-completion definitions
tui/                # Go Bubble Tea interactive TUI
bin/                # Compiled TUI binaries (gitignored)
scripts/            # Helper scripts
```

## Making Changes

### Shell Functions (`functions/`)

- Each function lives in its own file
- Follow existing naming conventions (`_worktree_*` for internal helpers)
- Test manually by reloading: `source ~/.zshrc`

### TUI (`tui/`)

The interactive TUI is written in Go with [Bubble Tea](https://github.com/charmbracelet/bubbletea).

```bash
cd tui
go build -o ../bin/worktree-tui .
go test ./...
```

## Commit Messages

We use [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>(<scope>): <description>

Examples:
feat(worktree): add support for custom config files
fix(review): handle edge case with branch names
docs: update installation instructions
test(tui): add unit tests for selection model
```

**Types:** `feat`, `fix`, `docs`, `test`, `refactor`, `chore`, `perf`, `ci`

## Pull Request Process

1. Create a feature branch from `main`
2. Make your changes with clear, focused commits
3. Ensure any new shell functions include help text (`--help` flag)
4. Update `README.md` if you add or change commands
5. Update `CHANGELOG.md` under the `[Unreleased]` section
6. Open a PR against `main` with a clear description of what and why

## Reporting Bugs

Use the [Bug Report](https://github.com/tmbtech/zsh-git-worktree-manager/issues/new?template=bug_report.md) issue template. Include:

- Your OS and shell version (`zsh --version`)
- Git version (`git --version`)
- Steps to reproduce
- Expected vs. actual behavior

## Suggesting Features

Use the [Feature Request](https://github.com/tmbtech/zsh-git-worktree-manager/issues/new?template=feature_request.md) issue template.

## Code Style

- **Shell**: Follow existing patterns in `functions/`. Use `local` for variables, quote expansions, and handle errors explicitly.
- **Go**: Run `gofmt` and `go vet` before committing.

## License

By contributing, you agree that your contributions will be licensed under the [MIT License](LICENSE).
