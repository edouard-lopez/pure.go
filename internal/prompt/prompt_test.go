package prompt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Prompt_serialize_correctly(t *testing.T) {
	prompt := Prompt{
		CurrentWorkingDir: "/home/edouard",
		LastStatusCommand: 0,
		Symbol:            "❯",
		GoVersion:         "1.2.3",
	}

	expected := "/home/edouard 🐐1.2.3\n0 ❯"

	actual := prompt.String()

	assert.Equal(t, expected, actual)
}
