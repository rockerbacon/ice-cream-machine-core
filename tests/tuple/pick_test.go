package tuple

import (
	assert "rockerbacon/ice-cream-machine-core/pkg/assert"
	testing "testing"
	tuple "rockerbacon/ice-cream-machine-core/pkg/tuple"
)

func returnMultValues() (int32, string) {
	return 3, "Hello World!"
}

func TestCanPickFirstFunctionReturn(t *testing.T) {
	assert.Equals(
		t,
		tuple.First(returnMultValues()),
		3,
	)
}

func TestCanPickSecondFunctionReturn(t *testing.T) {
	assert.Equals(
		t,
		tuple.Second(returnMultValues()),
		"Hello World!",
	)
}

