package assert

import (
	fmt "fmt"
	regexp "regexp"
)

type TestController interface {
	Fatal(...any);
}

func Equals[T comparable](
	t TestController,
	actual T,
	expected T,
) {
	if (actual != expected) {
		t.Fatal(
			fmt.Sprintf(
				"Expected value to equal %v, but got %v",
				expected,
				actual,
			),
		)
	}
}

// func IsNotNil[T comparable](
// 	t TestController,
// 	actual *T,
// ) {
// 	if (actual == nil) {
// 		t.Fatal(
// 			"Received nil value",
// 		)
// 	}
// }

func ErrorEquals(
	t TestController,
	err error,
	expectedError error,
) {
	if (err == nil) {
		t.Fatal("Function did not return error")
		return
	}

	if (err.Error() != expectedError.Error()) {
		t.Fatal(
			fmt.Sprintf(
				"Expected error with message '%s' but received '%s'",
				expectedError.Error(),
				err.Error(),
			),
		)
	}
}

func Matches(
	t TestController,
	str string,
	pattern string,
) {
	hasMatch, err := regexp.MatchString(pattern, str)

	if (err != nil) {
		t.Fatal(
			fmt.Sprintf(
				"An error occured while matching the string: %s",
				err.Error(),
			),
		)
		return
	}

	if (!hasMatch) {
		t.Fatal(
			fmt.Sprintf(
				"'%s' does not match the pattern '%s'",
				str,
				pattern,
			),
		)
	}

}
