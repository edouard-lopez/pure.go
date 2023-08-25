package constant

import (
	"testing"

	"github.com/carlmjohnson/be"
)

func TestExitCodeFailure(t *testing.T) {
	expect := 1

	actual := ExitCodeFailure

	be.Equal(t, actual, expect)
}

func TestExitCodeSuccess(t *testing.T) {
	expect := 0

	actual := ExitCodeSuccess

	be.Equal(t, actual, expect)
}
