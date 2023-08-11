package assert

import (
	testing "testing"
)

func Equals[T comparable](
	t *testing.T,
	actual T,
	expected T,
) {
	if (actual != expected) {
		t.Fatalf(
			"Expected value to equal %v, but got %v",
			expected,
			actual,
		)
	}
}

func IsNotNil[T comparable](
	t *testing.T,
	actual *T,
) {
	if (actual == nil) {
		t.Fatal(
			"Received nil value",
		)
	}
}

func ErrorEquals(
	t *testing.T,
	err error,
	expectedError error,
) {
	if (err == nil) {
		t.Fatal("Function did not return error")
	}

	if (err.Error() != expectedError.Error()) {
		t.Fatalf(
			"Expected error with message '%s' but received '%s'",
			err.Error(),
			expectedError.Error(),
		)
	}
}
