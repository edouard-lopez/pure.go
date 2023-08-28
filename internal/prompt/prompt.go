package prompt

import (
	"strings"
	"text/template"

	"github.com/edouard-lopez/pure.go/internal/constants"
)

type Prompt struct {
	CurrentWorkingDir string
	LastStatusCommand int
	Symbol            string
	GoVersion         string
}

func (prompt Prompt) String() string {
	var renderedMessage strings.Builder

	promptTemplate := template.Must(template.New("prompt").Parse(constants.PromptLayout))
	err := promptTemplate.Execute(&renderedMessage, prompt)
	if err != nil {
		panic(err)
	}

	return renderedMessage.String()
}
