package prod

import (
	"bytes"
	"math"
	"strconv"
)

func product(arr []int) []int {
	ln := len(arr)
	nw := make([]int, ln)
	for i := range nw {
		pr := 1
		for j := 0; j < ln; j++ {
			if i != j {
				pr = pr * arr[j]
			}
		}
		nw[i] = pr
	}
	return nw
}

// find the largest palindrpme made from the product of two n-digit numbers
//9009 = 91x99
func palindrome(n int) int {
	st := int(math.Pow10(n)) - 1
	mx := 0
	for i := st; i > 0; i-- {
		if i*st <= mx {
			break
		}
		for j := st; j > 0; j-- {
			pd := i * j
			if pd < mx {
				break
			}
			if pd > mx && isPalindrome(pd) {
				mx = pd
			}
		}
	}
	return mx
}

func isPalindrome(n int) bool {
	str := strconv.Itoa(n)
	rv := reverse(str)
	return rv == str
}

func reverse(s string) string {
	var buf bytes.Buffer
	b := []byte(s)
	ln := len(b)
	for i := ln - 1; i >= 0; i-- {
		buf.WriteByte(b[i])
	}
	return buf.String()
}
