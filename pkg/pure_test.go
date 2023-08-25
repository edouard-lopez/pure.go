package pure

import (
	"fmt"
	"testing"

	"github.com/carlmjohnson/be"
	"github.com/edouard-lopez/pure.go/internal/constant"
)

func TestGet_Last_Command_Succeed(t *testing.T) {
	lastStatusCommand := constant.ExitCodeSuccess

	expect := fmt.Sprintf("%v%v", lastStatusCommand, constant.PromptSymbol)
	actual := Get(lastStatusCommand)

	be.Equal(t, actual, expect)

}

func TestGet_Last_Command_Failed(t *testing.T) {
	lastStatusCommand := constant.ExitCodeFailure

	expect := fmt.Sprintf("%v%v", lastStatusCommand, constant.PromptSymbol)
	actual := Get(lastStatusCommand)

	fmt.Println(actual)
	be.Equal(t, actual, expect)
}
