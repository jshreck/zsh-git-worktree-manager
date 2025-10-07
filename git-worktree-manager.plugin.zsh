# Git Worktree Manager - Oh-My-Zsh Plugin
# Provides comprehensive git worktree management functions

# Get the directory where this plugin is installed
PLUGIN_DIR="${0:h}"

# Source all function files
for func_file in "${PLUGIN_DIR}"/functions/*; do
    source "${func_file}"
done

# Create alias for worktree_pull
alias wtp='worktree_pull'