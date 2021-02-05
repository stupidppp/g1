package math

import (
	"testing"
)

func TestFactor(t *testing.T) {
	if Factor(3) != 6 {
		t.Fatal(`Factor(3) != 6`)
	}
}
