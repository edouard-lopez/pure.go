package constants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExitCodeFailure(t *testing.T) {
	expected := 1

	actual := ExitCodeFailure

	assert.Equal(t, expected, actual)
}

func TestExitCodeSuccess(t *testing.T) {
	expected := 0

	actual := ExitCodeSuccess

	assert.Equal(t, expected, actual)
}

func TestPromptSymbol(t *testing.T) {
	expected := "❯"

	actual := PromptSymbol

	assert.Equal(t, expected, actual)
}
