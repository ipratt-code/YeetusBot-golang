package cmdErrors

import (
	"fmt"
	"strings"
)
// returns human readable errors
func FriendlyErrors(err error) error {
	if err.Error() == `HTTP 403 Forbidden, {"message": "Missing Permissions", "code": 50013}` {
		return fmt.Errorf("You dont have the permissions to do this command!")
	}else if strings.Contains(strings.ToLower(err.Error()), "404 not found") {
		return fmt.Errorf("The content could not be found! Please try something else.")
	}
	return err
}