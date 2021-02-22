package cmdErrors

import (
	"fmt"
)

func NeedRequiredArgumentsError(requiredArguments []string) error {
    return fmt.Errorf("Missing required argument(s): %v", requiredArguments)
}