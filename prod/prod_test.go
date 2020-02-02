package prod

import (
	"fmt"
	"testing"
)

func TestRatios(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}

	fmt.Println(product(s))

	a := []int{3, 2, 1}
	fmt.Println(product(a))
}

func TestReverse(t *testing.T) {

	fmt.Println(reverse("Hello"))

}

func TestPalindrome(t *testing.T) {
	fmt.Println(palindrome(3))
}
