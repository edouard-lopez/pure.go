package prompt

import (
	"strings"
	"text/template"
)

const PromptLayout = "{{.CurrentWorkingDir}}{{if .GoVersion}} 🐐{{.GoVersion}}{{end}}\n{{.Symbol}}"

type Prompt struct {
	CurrentWorkingDir string
	LastStatusCommand int
	Symbol            string
	GoVersion         string
}

func (prompt Prompt) String() string {
	var renderedMessage strings.Builder

	promptTemplate := template.Must(template.New("prompt").Parse(PromptLayout))
	err := promptTemplate.Execute(&renderedMessage, prompt)
	if err != nil {
		panic(err)
	}

	return renderedMessage.String()
}
