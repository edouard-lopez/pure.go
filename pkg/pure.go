package pure

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/edouard-lopez/pure.go/internal/constants"
	"github.com/edouard-lopez/pure.go/internal/current_working_dir"
	"github.com/edouard-lopez/pure.go/internal/prompt"
	prompt_symbol "github.com/edouard-lopez/pure.go/internal/prompt_symbol"
)

func Get(lastStatusCommand int) string {
	prompt := prompt.Prompt{
		CurrentWorkingDir: current_working_dir.Get(),
		LastStatusCommand: lastStatusCommand,
		Symbol:            prompt_symbol.Get(),
		GoVersion:         prompt_symbol.Get(),
	}

	var renderedMessage strings.Builder
	promptTemplate := template.Must(template.New("prompt").Parse(constants.PromptLayout))
	err := promptTemplate.Execute(&renderedMessage, prompt)
	if err != nil {
		panic(err)
	}

	return fmt.Sprint(prompt.String())
}
