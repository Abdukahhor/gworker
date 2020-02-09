package pairs

type choise func(int, int) int

type makeChoise func(choise) int

func cons(a, b int) makeChoise {
	p := func(i choise) int {
		return i(a, b)
	}
	return p
}

func car(do makeChoise) int {
	chooseRight := func(a, b int) int {
		return a
	}
	return do(chooseRight)
}

func cdr(do makeChoise) int {
	chooseLeft := func(a, b int) int {
		return b
	}
	return do(chooseLeft)
}
