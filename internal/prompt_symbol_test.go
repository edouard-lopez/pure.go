package prompt_symbol

import (
	"testing"

	"github.com/carlmjohnson/be"
)

func TestSymbol(t *testing.T) {
	expect := "❯"

	actual := Get()

	be.Equal(t, actual, expect)
}
