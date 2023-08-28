package prompt_symbol

import (
	"testing"

	"github.com/edouard-lopez/pure.go/internal/colorize"
	"github.com/edouard-lopez/pure.go/internal/constants"
	"github.com/stretchr/testify/assert"
)

func Test_Prompt_Symbol_when_last_command_is_success(t *testing.T) {
	lastStatusCommand := constants.ExitCodeSuccess
	expected := colorize.Success(constants.PromptSymbol)

	actual := Get(lastStatusCommand)

	assert.Equal(t, expected, actual)
}

func Test_Prompt_Symbol_when_last_command_is_failure(t *testing.T) {
	lastStatusCommand := constants.ExitCodeFailure
	expected := colorize.Danger(constants.PromptSymbol)

	actual := Get(lastStatusCommand)

	assert.Equal(t, expected, actual)
}
