package per

import "fmt"

// Format string to use in Error messages.
// It MUST contain %s for message and %s or %w for error.
var Format = "%s, err: %w"

// Errorf will return error with additional formatted message if provided err is.
func Errorf(err error, format string, a ...interface{}) error {
	return Error(err, fmt.Sprintf(format, a...))
}

// Error will return error wrapped with message formatted with Format if provided err is not nil.
func Error(err error, msg string) error {
	if err != nil {
		return fmt.Errorf(Format, msg, err)
	}
	return nil
}
