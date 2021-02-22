package cmdErrors

import (
	"fmt"
)

func BadArgumentsError(badArguments []string) error {
    return fmt.Errorf("Bad argument(s) (Bad formatting, bad spelling, incorrect id, etc.): %v", badArguments)
}