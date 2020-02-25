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

//"3x + 12 = 46"
//output 4
// Input:"4 - 2 = x"
// Output:2
// Input:"1x0 * 12 = 1200"
// Output:0
func missingDigit(str string) string {

	equal := "="
	if strings.Index(str, equal) == -1 {
		return ""
	}
	spl := strings.Split(str, " ")
	if len(spl) < 5 {
		return ""
	}
	isX := 0
	num1 := 0
	num2 := 0
	nm := ""
	for i := range spl {
		num, err := strconv.Atoi(spl[i])
		if err != nil && strings.Index(spl[i], "x") != -1 {
			isX = i
			nm = spl[i]
			continue
		}
		if num1 == 0 {
			num1 = num
		} else {
			num2 = num
		}

	}
	res := ""
	switch spl[1] {
	case "+":
		if isX > 2 {
			res = strconv.Itoa(num1 + num2)
			break
		}
		res = strconv.Itoa(num2 - num1)
	case "-":
		if isX > 2 {
			res = strconv.Itoa(num1 - num2)
			break
		} else if isX == 0 {
			res = strconv.Itoa(num1 + num2)
			break
		}
		res = strconv.Itoa(num1 - num2)
	case "*":
		if isX > 2 {
			res = strconv.Itoa(num1 * num2)
			break
		}
		res = strconv.Itoa(num2 / num1)
	case "/":
		if isX > 2 {
			res = strconv.Itoa(num1 / num2)
			break
		} else if isX == 0 {
			res = strconv.Itoa(num1 * num2)
			break
		}
		res = strconv.Itoa(num1 / num2)
	}
	if len(nm) > 1 {
		ix := strings.Index(nm, "x")
		return res[ix : ix+1]
	}
	// code goes here
	return res

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
