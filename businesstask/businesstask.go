package businesstask

import "errors"

var (
	ErrMissingRequiredKey = errors.New("required key is missing")
	ErrUnexpectedKey      = errors.New("unexpected key")
)

func Validate(input string) error {
	return nil
}
