package pure

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/edouard-lopez/pure.go/internal/colorize"
	"github.com/edouard-lopez/pure.go/internal/constants"
	"github.com/edouard-lopez/pure.go/internal/segments/current_working_dir"
	"github.com/edouard-lopez/pure.go/internal/prompt"
	"github.com/edouard-lopez/pure.go/internal/prompt_symbol"
	"github.com/edouard-lopez/pure.go/internal/segments/go_version"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

type MockedRuntime struct {
	mock.Mock
}

func (m *MockedRuntime) Output() ([]byte, error) {
	args := m.Called()
	return []byte(args.String(0)), errors.New("Missing go binary")
}

func mockGoVersion(fakeVersion string) string {
	mockRuntime := new(MockedRuntime)
	mockRuntime.On("Output").Return(fakeVersion)
	go_version.GetGoVersion = mockRuntime.Output // redefine go_version.GetGoVersion to use the mock

	return fakeVersion
}

func TestGet_Last_Command_Succeed(t *testing.T) {
	mockGoVersion("")
	expectedPrompt := prompt.Prompt{
		CurrentWorkingDir: current_working_dir.Get(),
		Symbol:            prompt_symbol.Get(constants.ExitCodeSuccess),
		GoVersion:         go_version.Get(),
		LastStatusCommand: constants.ExitCodeSuccess,
	}
	expected := expectedPrompt.String()

	actual := Get(expectedPrompt.LastStatusCommand)

	assert.Equal(t, expected, actual)

}

func TestGet_Last_Command_Failed(t *testing.T) {
	mockGoVersion("")
	expectedPrompt := prompt.Prompt{
		CurrentWorkingDir: current_working_dir.Get(),
		Symbol:            colorize.Danger(constants.PromptSymbol),
		GoVersion:         go_version.Get(),
		LastStatusCommand: constants.ExitCodeFailure,
	}
	expected := expectedPrompt.String()

	actual := Get(expectedPrompt.LastStatusCommand)

	assert.Equal(t, expected, actual)
}

func TestGet_Current_Working_Directory(t *testing.T) {
	err := os.Chdir(os.Getenv("HOME"))
	if err != nil {
		t.Fatal(err)
	}
	// cwd, err := os.Getwd()
	// if err != nil {
	// 	t.Fatal(err)
	// }

	mockGoVersion("")
	expectedPrompt := prompt.Prompt{
		CurrentWorkingDir: current_working_dir.Get(),
		Symbol:            prompt_symbol.Get(constants.ExitCodeSuccess),
		GoVersion:         go_version.Get(),
		LastStatusCommand: constants.ExitCodeSuccess,
	}
	expected := expectedPrompt.String()

	actual := Get(expectedPrompt.LastStatusCommand)

	assert.Equal(t, expected, actual)
}

func Test_Get_Go_Version(t *testing.T) {
	mockGoVersion("1.2.3")
	expectedPrompt := prompt.Prompt{
		CurrentWorkingDir: current_working_dir.Get(),
		Symbol:            prompt_symbol.Get(constants.ExitCodeSuccess),
		GoVersion:         go_version.Get(),
		LastStatusCommand: constants.ExitCodeSuccess,
	}
	expected := expectedPrompt.String()

	actual := Get(expectedPrompt.LastStatusCommand)

	assert.Equal(t, expected, actual)
}
