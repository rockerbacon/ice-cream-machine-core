package assert

import (
	assert "rockerbacon/ice-cream-machine-core/pkg/assert"
	errors "errors"
	testing "testing"
)

type Call struct {
	arguments []any;
}

type TestingMock struct {
	fatalCalls []Call;
}

func (self *TestingMock)Fatal(arguments... any) {
	self.fatalCalls = append(self.fatalCalls, Call{ arguments: arguments })
}

func NewTestingMock() TestingMock {
	return TestingMock{
		fatalCalls: make([]Call, 0),
	}
}

func TestEqualsDoesNothingWhenValuesAreEqual(t *testing.T) {
	testingMock := NewTestingMock()
	assert.Equals(&testingMock, 2, 2)

	if (len(testingMock.fatalCalls) != 0) {
		t.Fatalf(
			"Fatal was called %d times. First call used arguments: %+v",
			len(testingMock.fatalCalls),
			testingMock.fatalCalls[0].arguments,
		)
	}
}

func TestEqualsCallsFatalWhenValuesAreDifferent(t *testing.T) {
	testingMock := NewTestingMock()
	assert.Equals(&testingMock, 2, 3)

	if (len(testingMock.fatalCalls) != 1) {
		t.Fatalf("Fatal called incorrect number of times: %d", len(testingMock.fatalCalls))
	}

	if (testingMock.fatalCalls[0].arguments[0] != "Expected value to equal 3, but got 2") {
		t.Fatalf("Fatal called with wrong arguments: %+v", testingMock.fatalCalls[0].arguments)
	}
}

func TestErrorEqualsDoesNothingWhenErrorsAreEqual(t *testing.T) {
	testingMock := NewTestingMock()
	assert.ErrorEquals(&testingMock, errors.New("Hello World!"), errors.New("Hello World!"))

	if (len(testingMock.fatalCalls) != 0) {
		t.Fatalf(
			"Fatal was called %d times. First call used arguments: %+v",
			len(testingMock.fatalCalls),
			testingMock.fatalCalls[0].arguments,
		)
	}
}

func TestErrorEqualsCallsFatalWhenErrorsHaveDifferentMessages(t *testing.T) {
	testingMock := NewTestingMock()
	assert.ErrorEquals(&testingMock, errors.New("Hello World!"), errors.New("Bye World!"))

	if (len(testingMock.fatalCalls) != 1) {
		t.Fatalf("Fatal called incorrect number of times: %d", len(testingMock.fatalCalls))
	}

	if (testingMock.fatalCalls[0].arguments[0] != "Expected error with message 'Bye World!' but received 'Hello World!'") {
		t.Fatalf("Fatal called with wrong arguments: %+v", testingMock.fatalCalls[0].arguments)
	}
}

func TestErrorEqualsCallsFatalWhenActualErrorIsNil(t *testing.T) {
	testingMock := NewTestingMock()
	assert.ErrorEquals(&testingMock, nil, errors.New("Hello World!"))

	if (len(testingMock.fatalCalls) != 1) {
		t.Fatalf("Fatal called incorrect number of times: %d", len(testingMock.fatalCalls))
	}

	if (testingMock.fatalCalls[0].arguments[0] != "Function did not return error") {
		t.Fatalf("Fatal called with wrong arguments: %+v", testingMock.fatalCalls[0].arguments)
	}
}

func TestMatchesDoesNothingWhenStringMatchesExpectedPattern(t *testing.T) {
	testingMock := NewTestingMock()
	assert.Matches(&testingMock, "abcdedcba", "^[a-e]+$")

	if (len(testingMock.fatalCalls) != 0) {
		t.Fatalf(
			"Fatal was called %d times. First call used arguments: %+v",
			len(testingMock.fatalCalls),
			testingMock.fatalCalls[0].arguments,
		)
	}
}

func TestMatchesCallsFatalWhenStringDoesNotMatchExpectedPattern(t *testing.T) {
	testingMock := NewTestingMock()
	assert.Matches(&testingMock, "abcdefghij", "^[a-e]+$")

	if (len(testingMock.fatalCalls) != 1) {
		t.Fatalf("Fatal called incorrect number of times: %d", len(testingMock.fatalCalls))
	}

	if (testingMock.fatalCalls[0].arguments[0] != "'abcdefghij' does not match the pattern '^[a-e]+$'") {
		t.Fatalf("Fatal called with wrong arguments: %+v", testingMock.fatalCalls[0].arguments)
	}
}

func TestMatchesCallsFatalWhenThePatternCannotBeParsed(t *testing.T) {
	testingMock := NewTestingMock()
	assert.Matches(&testingMock, "abcdefghij", "a(b")

	if (len(testingMock.fatalCalls) != 1) {
		t.Fatalf("Fatal called incorrect number of times: %d", len(testingMock.fatalCalls))
	}

	if (testingMock.fatalCalls[0].arguments[0] != "An error occured while matching the string: error parsing regexp: missing closing ): `a(b`") {
		t.Fatalf("Fatal called with wrong arguments: %+v", testingMock.fatalCalls[0].arguments)
	}
}

