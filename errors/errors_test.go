// Package errors implements functions to manipulate errors.
package errors

import (
	"testing"
)

func TestNewEqual(t *testing.T) {
	const tmpErr Error = "tmp"
	// Different allocations should be equal.
	if tmpErr != New("tmp") {
		t.Errorf(`New("tmp") != New("tmp")`)
	}
}
