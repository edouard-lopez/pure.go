package git_dirty

import (
	"os"
	"os/exec"
	"testing"

	"github.com/edouard-lopez/pure.go/internal/colorize"
	"github.com/stretchr/testify/assert"
)

// chdir moves into dir. t.Chdir needs Go 1.24, and go.mod targets Go 1.20.
func chdir(t *testing.T, dir string) {
	t.Helper()
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
}

// fixtureGitRepo creates an empty git repository with one commit and moves into it.
func fixtureGitRepo(t *testing.T) {
	t.Helper()
	chdir(t, t.TempDir())
	for _, args := range [][]string{
		{"init", "--quiet"},
		{"config", "user.email", "test@example.com"},
		{"config", "user.name", "Test User"},
		{"commit", "--quiet", "--allow-empty", "--message", "init"},
	} {
		if err := exec.Command("git", args...).Run(); err != nil {
			t.Fatal(err)
		}
	}
}

func Test_Get_Git_Dirty_Clean_Repository(t *testing.T) {
	fixtureGitRepo(t)

	assert.Equal(t, "", Get())
}

func Test_Get_Git_Dirty_Untracked_File(t *testing.T) {
	fixtureGitRepo(t)
	if err := os.WriteFile("file.txt", []byte("untracked"), 0o644); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, colorize.Mute("*"), Get())
}

func Test_Get_Git_Dirty_Staged_File(t *testing.T) {
	fixtureGitRepo(t)
	if err := os.WriteFile("file.txt", []byte("staged"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := exec.Command("git", "add", "file.txt").Run(); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, colorize.Mute("*"), Get())
}

func Test_Get_Git_Dirty_Uses_PURE_SYMBOL_GIT_DIRTY(t *testing.T) {
	fixtureGitRepo(t)
	if err := os.WriteFile("file.txt", []byte("untracked"), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PURE_SYMBOL_GIT_DIRTY", "!")

	assert.Equal(t, colorize.Mute("!"), Get())
}

func Test_Get_Git_Dirty_Outside_Repository(t *testing.T) {
	chdir(t, t.TempDir())

	assert.Equal(t, "", Get())
}
