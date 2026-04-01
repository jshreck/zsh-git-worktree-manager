# Security Policy

## Reporting a Vulnerability

If you discover a security vulnerability in this project, please report it responsibly.

**Do not open a public issue.** Instead, email the maintainer directly or use [GitHub's private vulnerability reporting](https://github.com/tmbtech/zsh-git-worktree-manager/security/advisories/new).

Please include:

- A description of the vulnerability
- Steps to reproduce
- Potential impact
- Suggested fix (if any)

You can expect an initial response within 7 days.

## Scope

This project executes shell commands and interacts with git repositories on the user's machine. Security concerns include:

- Command injection via crafted branch names or repository URLs
- Unintended file operations (deletion, overwriting)
- Credential or secret exposure through environment file copying

## Supported Versions

Only the latest version on the `main` branch is actively supported with security fixes.
