package git_branch

import (
	"os/exec"
	"strings"

	"github.com/edouard-lopez/pure.go/internal/colorize"
)

// IsInsideWorkTree reports whether the current directory belongs to a git work tree.
func IsInsideWorkTree() bool {
	output, err := exec.Command("git", "rev-parse", "--is-inside-work-tree").Output()
	return err == nil && strings.TrimSpace(string(output)) == "true"
}

// parseBranch mirrors Fish's _pure_parse_git_branch:
// the branch name, or a ref name on a detached HEAD.
func parseBranch() string {
	output, err := exec.Command("git", "symbolic-ref", "--short", "HEAD").Output()
	if err != nil {
		output, err = exec.Command("git", "name-rev", "--name-only", "HEAD").Output()
		if err != nil {
			return ""
		}
	}
	return strings.TrimSpace(string(output))
}

func Get() string {
	if !IsInsideWorkTree() {
		return ""
	}

	branch := parseBranch()
	if branch == "" {
		return ""
	}
	return colorize.Mute(branch)
}
