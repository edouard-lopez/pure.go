package prompt_symbol

import (
	"github.com/edouard-lopez/pure.go/internal/colorize"
	"github.com/edouard-lopez/pure.go/internal/constants"
)

func Get(lastStatusCommand int) string {
	if lastStatusCommand == constants.ExitCodeSuccess {
		return colorize.Success(constants.PromptSymbol)
	}
	return colorize.Danger(constants.PromptSymbol)
}
