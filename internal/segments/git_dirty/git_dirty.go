package git_dirty

import (
	"os/exec"
	"strings"

	"github.com/edouard-lopez/pure.go/internal/colorize"
	"github.com/edouard-lopez/pure.go/internal/config"
	"github.com/edouard-lopez/pure.go/internal/constants"
)

// isDirty mirrors Fish's _pure_prompt_git_dirty: any untracked, unstaged or staged change.
func isDirty() bool {
	output, err := exec.Command("git", "status", "--porcelain", "--ignore-submodules").Output()
	return err == nil && strings.TrimSpace(string(output)) != ""
}

// Get returns the dirty symbol when the repository has changes, else an empty string.
func Get() string {
	if !isDirty() {
		return ""
	}
	return colorize.Mute(config.Get("PURE_SYMBOL_GIT_DIRTY", constants.GitDirtySymbol))
}
