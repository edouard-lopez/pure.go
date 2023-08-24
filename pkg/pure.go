package pure

import (
	"fmt"

	prompt_symbol "github.com/edouard-lopez/pure.go/internal"
)

func Get(lastStatusCommand int) string {
	return prompt_symbol.Get() + fmt.Sprint(lastStatusCommand)
}
