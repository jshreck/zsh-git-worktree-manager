package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/tmbtech/zsh-git-worktree-manager/tui/internal/data"
	"github.com/tmbtech/zsh-git-worktree-manager/tui/internal/model"
)

func main() {
	tuiData, err := data.ParseFromReader(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(2)
	}

	// Stdin was a pipe (consumed by ParseFromReader) and stdout is captured
	// by the calling shell's $() subshell. Open /dev/tty for both keyboard
	// input and TUI rendering so Bubble Tea talks directly to the terminal.
	ttyIn, err := os.Open("/dev/tty")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening terminal for input: %v\n", err)
		os.Exit(2)
	}
	defer ttyIn.Close()

	ttyOut, err := os.OpenFile("/dev/tty", os.O_WRONLY, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening terminal for output: %v\n", err)
		os.Exit(2)
	}
	defer ttyOut.Close()

	m := model.New(tuiData)
	p := tea.NewProgram(m,
		tea.WithAltScreen(),
		tea.WithInput(ttyIn),
		tea.WithOutput(ttyOut),
	)

	finalModel, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running TUI: %v\n", err)
		os.Exit(2)
	}

	// Print the selected action to stdout (captured by the shell wrapper).
	if result, ok := finalModel.(model.Model); ok {
		if action := result.Selected(); action != "" {
			fmt.Print(action)
			os.Exit(0)
		}
	}

	// User quit without selecting.
	os.Exit(1)
}
