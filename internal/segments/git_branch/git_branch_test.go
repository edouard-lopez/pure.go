package git_branch

import (
	"os"
	"os/exec"
	"testing"

	"github.com/edouard-lopez/pure.go/internal/colorize"
	"github.com/stretchr/testify/assert"
)

func Test_Get_Git_Branch_No_Repository(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "no-repo")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	expected := ""
	actual := Get()

	assert.Equal(t, expected, actual)
}

func Test_Get_Git_Branch_Repository(t *testing.T) {
	branchName := "feature/test-branch"
	cleanup := FixtureGitBranch(t, branchName)
	defer cleanup()

	expected := colorize.Mute(branchName)
	actual := Get()

	assert.Equal(t, expected, actual)
}

func Test_Get_Git_Branch_Detached_HEAD(t *testing.T) {
	branchName := "feature/detached"
	cleanup := FixtureGitBranch(t, branchName)
	defer cleanup()
	if err := exec.Command("git", "checkout", "--detach").Run(); err != nil {
		t.Fatal(err)
	}

	expected := colorize.Mute(branchName)
	actual := Get()

	assert.Equal(t, expected, actual)
}

func Test_Get_Git_Branch_From_Subdirectory(t *testing.T) {
	branchName := "feature/subdir"
	cleanup := FixtureGitBranch(t, branchName)
	defer cleanup()
	if err := os.Mkdir("subdir", 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir("subdir"); err != nil {
		t.Fatal(err)
	}

	expected := colorize.Mute(branchName)
	actual := Get()

	assert.Equal(t, expected, actual)
}
