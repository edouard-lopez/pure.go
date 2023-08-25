package pure

import (
	"fmt"

	"github.com/edouard-lopez/pure.go/internal/constants"
	"github.com/edouard-lopez/pure.go/internal/current_working_dir"
	prompt_symbol "github.com/edouard-lopez/pure.go/internal/prompt_symbol"
)

func Get(lastStatusCommand int) string {
	return fmt.Sprintf(constants.PromptLayout,
		current_working_dir.Get(),
		lastStatusCommand,
		prompt_symbol.Get(),
	)
}
