package task5

import (
	"strconv"
	"strings"
)

type Results struct {
	winnersName  string
	method1Count int
	method2Count int
}

func (R *Results) method1Won() int {
	R.method1Count += 1
	return R.method1Count
}

func (R *Results) method2Won() int {
	R.method2Count += 1
	return R.method2Count
}

func (R *Results) setWinnersName() {
	if R.method1Count > R.method2Count {
		R.winnersName = "method 1"
	} else if R.method2Count > R.method1Count {
		R.winnersName = "method 2"
	}
}

func makeSixDigits(n int) (s string) {
	s = strings.Repeat("0", 6-len(strconv.Itoa(n)))
	s += strconv.Itoa(n)
	return
}

func method1(n int) (b bool) {
	var s string
	if n < 100000 {
		s = makeSixDigits(n)
	} else {
		s = strconv.Itoa(n)
	}
	leftInt, _ := strconv.Atoi(s[:3])
	rightInt, _ := strconv.Atoi(s[3:])
	if leftInt == rightInt {
		b = true
	} else {
		b = false
	}

	return b
}

func getSum(n []int) (sum int) {
	for _, i := range n {
		sum += i
	}
	return
}

func method2(n int) (b bool) {
	var s string
	var odds, evens []int
	if n < 100000 {
		s = makeSixDigits(n)
	} else {
		s = strconv.Itoa(n)
	}
	var str []string = strings.Split(s, "")
	for _, k := range str {
		i, _ := strconv.Atoi(string(k))
		if i%2 == 0 {
			odds = append(odds, i)
		} else {
			evens = append(evens, i)
		}
	}
	if getSum(odds) == getSum(evens) {
		b = true
	} else {
		b = false
	}
	return b
}

func LuckyTickets(min, max int) Results {
	//numbers := getNumbers(min, max)
	var Res Results
	for i := min; i <= max; i++ {
		if method1(i) {
			Res.method1Won()
		}

		if method2(i) {
			Res.method2Won()
		}
	}
	Res.setWinnersName()

	return Res

}
