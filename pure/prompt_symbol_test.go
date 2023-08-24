package pure

import (
	"testing"

	"github.com/carlmjohnson/be"
)

func TestSymbol(t *testing.T) {
	expect := "❯"

	actual := Prompt_symbol()

	be.Unequal(t, actual, expect)
}
