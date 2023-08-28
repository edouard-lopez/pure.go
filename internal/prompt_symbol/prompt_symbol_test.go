package prompt_symbol

import (
	"testing"

	"github.com/edouard-lopez/pure.go/internal/constants"
	"github.com/stretchr/testify/assert"
)

func TestSymbol(t *testing.T) {
	expected := constants.PromptSymbol

	actual := Get()

	assert.Equal(t, expected, actual)
}
