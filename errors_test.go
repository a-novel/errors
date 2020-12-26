package errors

import "testing"

func TestError(t *testing.T) {
	// Should create an error.
	err := New("err1", "something happened")
	if err == nil {
		t.Fatalf("error not created")
	}

	// Should print error message.
	if err.Error() != "something happened" {
		t.Errorf("unexpected error message : got '%s' instead of 'something happened'", err.Error())
	}

	if err.ID != "err1" {
		t.Errorf("unexpected error id : got '%s' instead of 'err1'", err.ID)
	}
}
