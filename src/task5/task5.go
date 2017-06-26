package task5

import (
	"errors"
	"strconv"
	"strings"
)

type T5 struct {
	Min, Max int
}

func (T *T5) Dotask5() (err error, data interface{}) {
	return Dotask(T.Min, T.Max)
}

// обязательно переделать работу со строками на инты

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
		R.winnersName = "method 1 "
	} else if R.method2Count > R.method1Count {
		R.winnersName = "method 2 "
	}
}

func makeSixDigits(num int) []int {
	numSlice := make([]int, 6)
	for i := 5; num > 0; i-- {
		numSlice[i] = num % 10
		num = num / 10
	}
	return numSlice
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
	var numSl []int
	var sumLeft, sumRight int

	numSl = makeSixDigits(n)

	leftN := numSl[:3]
	rightN := numSl[3:]

	sumLeft = getSum(leftN)
	sumRight = getSum(rightN)

	if sumLeft == sumRight {
		return true
	}

	return b
}

func getSum(n []int) (sum int) {
	for _, v := range n {
		sum += v
	}
	return
}

func method2(n int) (b bool) {
	var numSl []int
	var odds, evens []int

	numSl = makeSixDigits(n)

	for _, v := range numSl {
		if v%2 == 0 {
			odds = append(odds, v)

		} else {
			evens = append(evens, v)
		}
	}

	if getSum(odds) == getSum(evens) {
		return true

	}

	return b
}

func Dotask(min, max int) (err error, data Results) {
	if min < 0 || max < 0 {
		return errors.New("Значения должны быть больше нуля"), data

	} else if min > max {
		return errors.New("Минимальное значение не должно быть больше максимального"), data

	} else {
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
		return err, Res
	}
}
