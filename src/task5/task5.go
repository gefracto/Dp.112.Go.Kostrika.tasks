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

func getInts(s string) []int {
	var sl []int
	slStr := strings.Split(s, "")
	for _, i := range slStr {
		l, _ := strconv.Atoi(i)
		sl = append(sl, l)
	}
	return sl
}

func method1(n int) (b bool) {
	var s string
	var sumLeft, sumRight int
	if n < 100000 {
		s = makeSixDigits(n)
	} else {
		s = strconv.Itoa(n)
	}
	leftN := getInts(s[:3])
	rightN := getInts(s[3:])

	sumLeft = getSum(leftN)
	sumRight = getSum(rightN)

	if sumLeft == sumRight {
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

func DoTask5(min, max int) Results {
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
