package per

import "fmt"

var Format = "%s, err: %w"

func Errorf(err error, format string, a ...interface{}) error {
	return Error(err, fmt.Sprintf(format, a...))
}

func Error(err error, msg string) error {
	if err != nil {
		return fmt.Errorf(Format, msg, err)
	}
	return nil
}
