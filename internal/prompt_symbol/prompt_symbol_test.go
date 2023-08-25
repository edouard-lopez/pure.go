package prompt_symbol

import (
	"testing"

	"github.com/carlmjohnson/be"
	"github.com/edouard-lopez/pure.go/internal/constant"
)

func TestSymbol(t *testing.T) {
	expect := constant.PromptSymbol

	actual := Get()

	be.Equal(t, actual, expect)
}
