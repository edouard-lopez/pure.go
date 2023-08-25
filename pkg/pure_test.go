package pure

import (
	"testing"

	"github.com/carlmjohnson/be"
)

func PureGet_Last_Command_Succeed(t *testing.T) {
	expect := "❯0"
	lastStatusCommand := 0

	actual := Get(lastStatusCommand)

	be.Equal(t, actual, expect)

}

func PureGet_Last_Command_Failed(t *testing.T) {
	expect := "❯1"
	lastStatusCommand := 1
	actual := Get(lastStatusCommand)

	be.Equal(t, actual, expect)

}
