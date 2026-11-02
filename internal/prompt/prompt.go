package prompt

import (
	"strings"
	"text/template"
)

// PromptLayout starts with a blank line to separate the prompt from previous output, like Fish's new_line.
const PromptLayout = "\n" +
	`{{.CurrentWorkingDir}}` +
	`{{if .GitBranch}} {{.GitBranch}}{{.GitDirty}}{{end}}` +
	`{{if .GoVersion}} 🐐{{.GoVersion}}{{end}}` +
	"\n" +
	`{{.Symbol}}`

type Prompt struct {
	CurrentWorkingDir string
	LastStatusCommand int
	Symbol            string
	GoVersion         string
	GitBranch         string
	GitDirty          string
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
