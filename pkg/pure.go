package pure

import (
	"fmt"

	prompt_symbol "github.com/edouard-lopez/pure.go/internal/prompt_symbol"
)

func Get(lastStatusCommand int) string {
	return fmt.Sprintf("%v%v",
		lastStatusCommand,
		prompt_symbol.Get(),
	)
}
