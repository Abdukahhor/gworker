package problem1

import "sort"

func p1(s []int, k int) bool {
	ln := len(s)
	if ln < 2 {
		return false
	}
	sort.Ints(s)
	f := 0
	l := ln - 1
	for {
		if s[f]+s[l] == k {
			return true
		} else if s[f]+s[l] < k {
			f++
		} else {
			l--
		}
	}
}
