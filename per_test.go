package per

import (
	"errors"
	"testing"
)

func Test(t *testing.T) {
	err := errors.New("oh no!")
	if Error(err, "wrong :(") == nil {
		t.Error("expected error")
	}
	if Errorf(err, "error code: %d", 123) == nil {
		t.Error("expected error")
	}
	if err = nil; Error(err, "in case of error") != nil {
		t.Error("expected nil")
	}

	Format = "MSG: %s ERR : %s"
	err = errors.New("deadline exceeded")
	if Error(err, "create user, failed").Error() != "MSG: create user, failed ERR : deadline exceeded" {
		t.Error("expected format change")
	}
}
