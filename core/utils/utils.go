package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func SwitchLayout(layout string) error {
	// Send notification
	notifyCmd := exec.Command("notify-send", "Layout Switch", layout)
	if err := notifyCmd.Run(); err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}

	// Switch keyboard layout
	args := strings.Fields(layout)
	setxkbmapArgs := append([]string{"setxkbmap"}, args...)
	layoutCmd := exec.Command(setxkbmapArgs[0], setxkbmapArgs[1:]...)
	if err := layoutCmd.Run(); err != nil {
		return fmt.Errorf("failed to switch layout: %w", err)
	}

	return nil
}

func RunProgram(command []string, mainLayout, secondaryLayout string) int {
	// Switch to secondary layout
	if err := SwitchLayout(secondaryLayout); err != nil {
		fmt.Fprintf(os.Stderr, "Error switching to secondary layout: %v\n", err)
		return 1
	}

	// Run the command and wait for it to finish
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()

	// Switch back to main layout
	if err := SwitchLayout(mainLayout); err != nil {
		fmt.Fprintf(os.Stderr, "Error switching back to main layout: %v\n", err)
	}

	// Get exit code
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
				return status.ExitStatus()
			}
		}
		return 1
	}

	return 0
}
