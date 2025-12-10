package git_branch

import (
	"os"
	"os/exec"
	"testing"
)

// FixtureGitBranch creates a temporary git repository and checks out the specified branch.
// Returns a cleanup function to remove the temp directory.
func FixtureGitBranch(t *testing.T, branchName string) func() {
	t.Helper()

	tempDir, err := os.MkdirTemp("", "git-repo")
	if err != nil {
		t.Fatal(err)
	}

	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("git", "init")
	err = cmd.Run()
	if err != nil {
		t.Fatal(err)
	}

	// Configure git identity (required in CI environments)
	cmd = exec.Command("git", "config", "user.email", "test@example.com")
	err = cmd.Run()
	if err != nil {
		t.Fatal(err)
	}

	cmd = exec.Command("git", "config", "user.name", "Test User")
	err = cmd.Run()
	if err != nil {
		t.Fatal(err)
	}

	cmd = exec.Command("git", "commit", "--allow-empty", "-m", "Initial commit")
	err = cmd.Run()
	if err != nil {
		t.Fatal(err)
	}

	cmd = exec.Command("git", "checkout", "-b", branchName)
	err = cmd.Run()
	if err != nil {
		t.Fatal(err)
	}

	return func() {
		os.RemoveAll(tempDir)
	}
}
