package problem1

import (
	"strconv"
	"strings"
)

func findIntersection(strArr []string) string {

	// code goes here
	if len(strArr) != 2 {
		return ""
	}
	rtn := make([]string, 0)
	str0 := strings.Split(strArr[0], ",")
	//str1 := strings.Split(strArr[1], ",")
	for i := range str0 {
		if strings.Index(strArr[1], str0[i]) >= 0 {
			rtn = append(rtn, str0[i])
		}
	}
	return strings.Join(rtn, ", ")

}
func missingDigit(str string) string {

	equal := "="
	if strings.Index(str, equal) == -1 {
		return ""
	}
	spl := strings.Split(str, " ")

	isX := 0
	num1 := 0
	num2 := 0
	for i := range spl {
		num, err := strconv.Atoi(spl[i])
		if err != nil {
			isX = i
			continue
		}
		if num1 == 0 {
			num1 = num
		} else {
			num2 = num
		}

	}

	switch spl[1] {
	case "+":
		if isX > 2 {
			return strconv.Itoa(num1 + num2)
		}
		return strconv.Itoa(num2 - num1)
	case "-":
		if isX > 2 {
			return strconv.Itoa(num1 - num2)
		} else if isX == 0 {
			return strconv.Itoa(num1 + num2)
		}
		return strconv.Itoa(num1 - num2)
	case "*":
		if isX > 2 {
			return strconv.Itoa(num1 * num2)
		}
		return strconv.Itoa(num2 / num1)
	case "/":
		if isX > 2 {
			return strconv.Itoa(num1 / num2)
		} else if isX == 0 {
			return strconv.Itoa(num1 * num2)
		}
		return strconv.Itoa(num1 / num2)
	}
	// code goes here
	return ""

}

//[]string {"baseball", "a,all,b,ball,bas,base,cat,code,d,e,quit,z"}
//Output: base,ball
func wordSplit(strArr []string) string {
	if len(strArr) != 2 {
		return ""
	}
	str0 := strings.Split(strArr[1], ",")

	ln := len(str0)
	for i := 0; i < ln; i++ {
		for j := 0; j < ln; j++ {
			if str0[i]+str0[j] == strArr[0] {
				return str0[i] + "," + str0[j]
			}
		}
	}
	// code goes here
	return ""

}
