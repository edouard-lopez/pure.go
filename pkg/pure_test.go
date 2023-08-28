package pure

import (
	"log"
	"os"
	"testing"

	"github.com/edouard-lopez/pure.go/internal/constants"
	"github.com/edouard-lopez/pure.go/internal/current_working_dir"
	"github.com/edouard-lopez/pure.go/internal/prompt"
	"github.com/stretchr/testify/assert"
)

var cwd string

func setup() {
	cwd = changeToTmpDir()
}

func teardown() {}

func changeToTmpDir() string {
	err := os.Chdir(`/tmp`)
	if err != nil {
		log.Fatal(err)
	}

	expected, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return expected
}

func TestMain(m *testing.M) {
	setup()
	defer teardown()

	m.Run()
}

func TestGet_Last_Command_Succeed(t *testing.T) {
	prompt := prompt.Prompt{
		CurrentWorkingDir: current_working_dir.Get(),
		LastStatusCommand: constants.ExitCodeSuccess,
		Symbol:            constants.PromptSymbol,
		GoVersion:         "1.8.3",
	}
	expected := prompt.String()

	actual := Get(prompt.LastStatusCommand)

	assert.Equal(t, expected, actual)

}

func TestGet_Last_Command_Failed(t *testing.T) {
	prompt := prompt.Prompt{
		CurrentWorkingDir: current_working_dir.Get(),
		LastStatusCommand: constants.ExitCodeFailure,
		Symbol:            constants.PromptSymbol,
		GoVersion:         "1.8.3",
	}
	expected := prompt.String()

	actual := Get(prompt.LastStatusCommand)

	assert.Equal(t, expected, actual)
}

func TestGet_Current_Working_Directory(t *testing.T) {
	err := os.Chdir(os.Getenv("HOME"))
	if err != nil {
		t.Fatal(err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	prompt := prompt.Prompt{
		CurrentWorkingDir: current_working_dir.ReplaceByTilde(cwd),
		LastStatusCommand: constants.ExitCodeSuccess,
		Symbol:            constants.PromptSymbol,
		GoVersion:         "1.8.3",
	}
	expected := prompt.String()

	actual := Get(prompt.LastStatusCommand)

	assert.Equal(t, expected, actual)
}

func Test_Get_Go_Version(t *testing.T) {
	prompt := prompt.Prompt{
		CurrentWorkingDir: current_working_dir.Get(),
		LastStatusCommand: constants.ExitCodeSuccess,
		Symbol:            constants.PromptSymbol,
		GoVersion:         "1.8.3",
	}
	expected := prompt.String()

	actual := Get(prompt.LastStatusCommand)

	assert.Equal(t, expected, actual)
}
