package git_branch

import (
	"os/exec"
	"strings"

	"github.com/edouard-lopez/pure.go/internal/colorize"
)

// parseBranch mirrors Fish's _pure_parse_git_branch:
// the branch name, or a ref name on a detached HEAD.
// `symbolic-ref -q` fails quietly on a detached HEAD and outside a repository,
// so the `name-rev` fallback also covers the "not a repository" case.
func parseBranch() string {
	output, err := exec.Command("git", "symbolic-ref", "-q", "--short", "HEAD").Output()
	if err != nil {
		output, err = exec.Command("git", "name-rev", "--name-only", "HEAD").Output()
		if err != nil {
			return ""
		}
	}
	return strings.TrimSpace(string(output))
}

func Get() string {
	branch := parseBranch()
	if branch == "" {
		return ""
	}
	return colorize.Mute(branch)
}
