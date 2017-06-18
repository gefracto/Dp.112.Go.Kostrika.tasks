package task3

import (
	"math"
	"sort"
)

type Triangle struct {
	Name string
	A    float64
	B    float64
	C    float64
}

func (T *Triangle) getP() float64 {
	return (T.A + T.B + T.C) / 2
}

func (T *Triangle) getSquare(p float64) float64 {
	S := math.Sqrt(p * (p - T.A) * (p - T.B) * (p - T.C))
	return S
}

func (T *Triangle) getName() string {
	return T.Name
}

func reverseSliceOfNames (s []string) (newS []string) {
	for i := len(s) - 1; i >=0; i-- {
		newS = append(newS, s[i])
	}
	return newS
}

func DoTask3(T []Triangle) []string {
	var names []string
	var squares []float64
	var indexes []int
	var sortedSquares []float64
	var sortedNames []string

	for i := range T {
		names = append(names, T[i].getName())
		squares = append(squares, T[i].getSquare(T[i].getP()))
	}

	sortedSquares = append(sortedSquares, squares...)
	sort.Float64s(sortedSquares)
	for _, v := range sortedSquares {
		for i2, v2 := range squares {
			if v == v2 {
				indexes = append(indexes, i2)
				squares[i2] = -1
			}
		}
	}

	for _, v := range indexes {
		sortedNames = append(sortedNames, names[v])
	}
	reversedSortedNames := reverseSliceOfNames(sortedNames)
	return reversedSortedNames

}
