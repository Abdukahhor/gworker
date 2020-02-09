package problem1

func lowestInt(arr []int) int {
	ln := len(arr)
	m := make([]bool, ln+1)

	for i := range arr {
		if arr[i] > 0 && arr[i] <= ln {
			m[arr[i]] = true
		}
	}

	for i := 1; i <= ln; i++ {
		if !m[i] {
			return i
		}
	}

	return ln + 1
}
