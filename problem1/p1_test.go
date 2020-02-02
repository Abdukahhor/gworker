package problem1

import (
	"testing"
)

func TestP1(t *testing.T) {
	s := []int{13, 10, 3, 7}

	if !p1(s, 17) {
		t.Error("false")
	}
}
