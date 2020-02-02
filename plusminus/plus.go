package plusminus

import "sort"

func ratios(arr []int) (p float32, m float32, z float32) {
	for i := range arr {
		if arr[i] > 0 {
			p++
		} else if arr[i] < 0 {
			m++
		} else {
			z++
		}
	}
	t := float32(len(arr))
	return p / t, m / t, z / t
}

func twosum(arr []int, t int) []int {
	sort.Ints(arr)
	f := 0
	l := len(arr) - 1
	for {
		if arr[f]+arr[l] == t {
			break
		} else if arr[f]+arr[l] < t {
			f++
		} else {
			l--
		}
	}
	return []int{f, l}
}
