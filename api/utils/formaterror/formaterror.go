package formaterror

import (
	"errors"
	"strings"
)

/*
Format Error
*/
func FormatError(err string) error {
	if strings.Contains(err, "name") {
		return errors.New("Name is already taken.")
	}
	if strings.Contains(err, "email") {
		return errors.New("Email already taken.")
	}
	return errors.New("Incorrect Details")
}
