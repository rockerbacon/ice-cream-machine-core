package assert

import (
	"regexp"
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

func Matches(
	t *testing.T,
	str string,
	pattern string,
) {
	hasMatch, err := regexp.MatchString(pattern, str)

	if (err != nil) {
		t.Fatalf("An error occured while matching the string: %s", err.Error())
	}

	if (!hasMatch) {
		t.Fatalf("'%s' does not match the pattern '%s'", str, pattern)
	}

}
