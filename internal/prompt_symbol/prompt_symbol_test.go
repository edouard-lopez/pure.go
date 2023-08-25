package prompt_symbol

import (
	"testing"

	"github.com/carlmjohnson/be"
	"github.com/edouard-lopez/pure.go/internal/constants"
)

func TestSymbol(t *testing.T) {
	expect := constants.PromptSymbol

	actual := Get()

	be.Equal(t, actual, expect)
}
