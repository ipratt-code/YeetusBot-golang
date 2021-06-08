package cmdErrors

import (
	"fmt"
	"strings"
)
// returns human readable errors
func FriendlyErrors(err error) error {
	if err.Error() == `HTTP 403 Forbidden, {"message": "Missing Permissions", "code": 50013}` {
		return fmt.Errorf("The bot does not have the permission to do this command!")
	}else if strings.Contains(strings.ToLower(err.Error()), "404 not found") {
		return fmt.Errorf("The content could not be found! Please try something else.")
	}else if strings.Contains(strings.ToLower(err.Error()), "403 forbidden") {
		return fmt.Errorf("The content is private! Please try something else.")
	}
	return err
}