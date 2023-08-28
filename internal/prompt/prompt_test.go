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
		GoVersion:         "1.13.4",
	}

	expected := "/home/edouard\n0 ❯"

	actual := prompt.String()

	assert.Equal(t, expected, actual)
}
