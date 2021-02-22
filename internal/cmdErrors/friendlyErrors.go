package cmdErrors

import (
	"fmt"
)
// returns human readable errors
func FriendlyErrors(err error) error {
	if err.Error() == `HTTP 403 Forbidden, {"message": "Missing Permissions", "code": 50013}` {
		return fmt.Errorf("You dont have the permissions to do this command!")
	}
	return err
}