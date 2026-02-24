#!/bin/zsh

# Git Worktree Manager Installation Script

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "🌳 Git Worktree Manager Installation"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""

# Check for Oh-My-Zsh
if [[ ! -d "$HOME/.oh-my-zsh" ]]; then
  echo "❌ Error: Oh-My-Zsh is not installed"
  echo "   Install it from: https://ohmyz.sh/"
  exit 1
fi
echo "✅ Oh-My-Zsh found"

# Check for git
if ! command -v git &>/dev/null; then
  echo "❌ Error: git is not installed"
  exit 1
fi
echo "✅ Git found ($(git --version))"

# Check for gh CLI (optional but recommended)
if command -v gh &>/dev/null; then
  echo "✅ GitHub CLI found ($(gh --version | head -1))"
else
  echo "⚠️  Warning: GitHub CLI (gh) not found"
  echo "   The 'worktree review' command requires gh CLI"
  echo "   Install it with: brew install gh"
fi

# Check for yarn (optional but recommended)
if command -v yarn &>/dev/null; then
  echo "✅ Yarn found ($(yarn --version))"
else
  echo "⚠️  Warning: Yarn not found"
  echo "   You'll need yarn for automatic dependency installation"
  echo "   Install it with: npm install -g yarn"
fi

echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "📦 Installing plugin..."
echo ""

# Get the directory where this script is located
PLUGIN_SOURCE_DIR="$(cd "$(dirname "$0")" && pwd)"
PLUGIN_NAME="git-worktree-manager"
PLUGIN_DEST_DIR="$HOME/.oh-my-zsh/custom/plugins/$PLUGIN_NAME"

# Check if plugin already installed
if [[ -L "$PLUGIN_DEST_DIR" ]]; then
  echo "⚠️  Plugin symlink already exists at: $PLUGIN_DEST_DIR"
  echo "   Removing existing symlink..."
  rm "$PLUGIN_DEST_DIR"
elif [[ -d "$PLUGIN_DEST_DIR" ]]; then
  echo "⚠️  Plugin directory already exists at: $PLUGIN_DEST_DIR"
  echo "   Please remove it manually and run this script again"
  exit 1
fi

# Create symlink
if ln -s "$PLUGIN_SOURCE_DIR" "$PLUGIN_DEST_DIR"; then
  echo "✅ Plugin symlink created successfully"
else
  echo "❌ Error: Failed to create symlink"
  exit 1
fi

echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "🎉 Installation complete!"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "📝 Next steps:"
echo ""
echo "1. Add the plugin to your ~/.zshrc plugins array:"
echo "   plugins=(git zsh-autosuggestions brew git-worktree-manager)"
echo ""
echo "2. Reload your shell:"
echo "   source ~/.zshrc"
echo ""
echo "3. Test the installation:"
echo "   worktree setup --help"
echo ""
echo "For more information, see: $PLUGIN_SOURCE_DIR/README.md"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"