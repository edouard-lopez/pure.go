package pure

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/edouard-lopez/pure.go/internal/constants"
	"github.com/edouard-lopez/pure.go/internal/current_working_dir"
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
	lastStatusCommand := constants.ExitCodeSuccess

	expected := fmt.Sprintf(constants.PromptLayout, cwd, lastStatusCommand, constants.PromptSymbol)
	actual := Get(lastStatusCommand)

	assert.Equal(t, expected, actual)

}

func TestGet_Last_Command_Failed(t *testing.T) {
	lastStatusCommand := constants.ExitCodeFailure

	expected := fmt.Sprintf(constants.PromptLayout, cwd, lastStatusCommand, constants.PromptSymbol)
	actual := Get(lastStatusCommand)

	fmt.Println(actual)
	assert.Equal(t, expected, actual)
}

func TestGet_Current_Working_Directory(t *testing.T) {
	lastStatusCommand := constants.ExitCodeSuccess
	err := os.Chdir(os.Getenv("HOME"))
	if err != nil {
		t.Fatal(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	expected := fmt.Sprintf(constants.PromptLayout,
		current_working_dir.ReplaceByTilde(cwd),
		lastStatusCommand,
		constants.PromptSymbol,
	)
	actual := Get(lastStatusCommand)

	assert.Equal(t, expected, actual)

}
