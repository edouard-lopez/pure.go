package prompt_symbol

import (
	"github.com/edouard-lopez/pure.go/internal/colorize"
	"github.com/edouard-lopez/pure.go/internal/config"
	"github.com/edouard-lopez/pure.go/internal/constants"
)

func Get(lastStatusCommand int) string {
	symbol := config.Get("PURE_SYMBOL_PROMPT", constants.PromptSymbol)
	if lastStatusCommand == constants.ExitCodeSuccess {
		return colorize.Success(symbol)
	}
	return colorize.Danger(symbol)
}
