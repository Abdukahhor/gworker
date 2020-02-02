package plusminus

import "testing"

import "fmt"

func TestRatios(t *testing.T) {
	s := []int{1, 1, 3, -1, -7, 0}
	fmt.Println(ratios(s))
}

func TestTwosum(t *testing.T) {
	s := []int{2, 7, 11, 15}
	fmt.Println(twosum(s, 9))
}
