package problem1

import (
	"fmt"
	"testing"
)

func TestLowestInt(t *testing.T) {
	s := []int{1, 1, 3, 7}

	fmt.Println(lowestInt(s))
}
