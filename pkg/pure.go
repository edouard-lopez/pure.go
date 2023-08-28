package pure

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/edouard-lopez/pure.go/internal/current_working_dir"
	"github.com/edouard-lopez/pure.go/internal/prompt"
	"github.com/edouard-lopez/pure.go/internal/prompt_symbol"
	"github.com/edouard-lopez/pure.go/internal/segments/go_version"
)

func Get(lastStatusCommand int) string {
	_prompt := prompt.Prompt{
		CurrentWorkingDir: current_working_dir.Get(),
		LastStatusCommand: lastStatusCommand,
		Symbol:            prompt_symbol.Get(),
		GoVersion:         go_version.Get(),
	}

	var renderedMessage strings.Builder
	promptTemplate := template.Must(template.New("prompt").Parse(prompt.PromptLayout))
	err := promptTemplate.Execute(&renderedMessage, _prompt)
	if err != nil {
		panic(err)
	}

	return fmt.Sprint(_prompt.String())
}
