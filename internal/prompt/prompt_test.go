package prompt

import (
	"testing"

	"github.com/edouard-lopez/pure.go/internal/colorize"
	"github.com/edouard-lopez/pure.go/internal/prompt_symbol"
	"github.com/stretchr/testify/assert"
)

func Test_Prompt_serialize_correctly(t *testing.T) {
	lastStatusCommand := 1
	prompt := Prompt{
		CurrentWorkingDir: "/home/edouard",
		LastStatusCommand: lastStatusCommand,
		Symbol:            prompt_symbol.Get(lastStatusCommand),
		GoVersion:         "1.2.3",
	}

	expected := "/home/edouard 🐐1.2.3\n" + colorize.Danger("❯")

	actual := prompt.String()

	assert.Equal(t, expected, actual)
}
