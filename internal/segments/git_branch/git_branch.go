package git_branch

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/edouard-lopez/pure.go/internal/colorize"
)

func Get() string {
	gitDir := ".git"

	// Check if .git directory exists
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		return ""
	}

	// Get the current branch name
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	raw_branch := string(output)
	branch := raw_branch[:len(raw_branch)-1] // Remove the newline character
	return colorize.Mute(fmt.Sprintf("%s", branch))
}
