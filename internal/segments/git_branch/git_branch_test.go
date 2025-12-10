package git_branch

import (
	"os"
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
