#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
TUI_DIR="$PROJECT_ROOT/tui"
BIN_DIR="$PROJECT_ROOT/bin"

mkdir -p "$BIN_DIR"

cd "$TUI_DIR"

LDFLAGS="-s -w"

# Cross-compile targets.
TARGETS=(
  "darwin/arm64"
  "darwin/amd64"
  "linux/amd64"
  "linux/arm64"
)

echo "Building worktree-tui..."

for target in "${TARGETS[@]}"; do
  os="${target%/*}"
  arch="${target#*/}"
  output="$BIN_DIR/worktree-tui-${os}-${arch}"
  echo "  ${os}/${arch} -> $(basename "$output")"
  GOOS="$os" GOARCH="$arch" go build -ldflags="$LDFLAGS" -o "$output" ./cmd/worktree-tui/
done

# Symlink the current platform binary for local use.
current_os="$(uname -s | tr '[:upper:]' '[:lower:]')"
current_arch="$(uname -m)"
case "$current_arch" in
  x86_64) current_arch="amd64" ;;
  aarch64) current_arch="arm64" ;;
esac

platform_bin="$BIN_DIR/worktree-tui-${current_os}-${current_arch}"
if [[ -f "$platform_bin" ]]; then
  ln -sf "$(basename "$platform_bin")" "$BIN_DIR/worktree-tui"
  echo ""
  echo "Linked bin/worktree-tui -> $(basename "$platform_bin")"
fi

echo ""
echo "Done. Binaries in $BIN_DIR/"
