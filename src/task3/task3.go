package task3

import (
	"math"
	"sort"
	"regexp"
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

func reverseSliceOfNames(s []string) (newS []string) {
	for i := len(s) - 1; i >= 0; i-- {
		newS = append(newS, s[i])
	}
	return newS
}

/*func isTriangle(a,b,c float64) bool {
	if
}*/

func checkTriangles(T []Triangle) (ok bool, reason string) {
	for _, v := range T {
		if ok, _ := regexp.MatchString("\\A[A-Z]{3}$", v.Name); ok == false {
			return false, "Имя треугольника " + v.Name + " должно состоять из имен трех вершин\n"

		} else if v.A >= v.B+v.C || v.A <= 0 {
			return false, "Значение стороны \"А\" треугольника \"" + v.Name +
				"\"\nне должно быть больше суммы сторон \"B\" и \"C\", " +
				"\nа также не может быть меньше или равно нулю\n"

		} else if v.B >= v.A+v.C || v.B <= 0 {
			return false, "Значение стороны \"B\" треугольника \"" + v.Name +
				"\"\nне должно быть больше суммы сторон \"A\" и \"C\", " +
				"\nа также не может быть меньше или равно нулю\n"

		} else if v.C >= v.A+v.B || v.C <= 0 {
			return false, "Значение стороны \"C\" треугольника \"" + v.Name +
				"\"\nне должно быть больше суммы сторон \"B\" и \"A\", " +
				"\nа также не может быть меньше или равно нулю\n"
		}
	}
	return true, ""
}

func DoTask3(T []Triangle) (ok bool, data []string, reason string) {
	if ok, reason := checkTriangles(T); ok == false {
		return ok, nil, reason

	} else {
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

		reversedNames := reverseSliceOfNames(sortedNames)
		return ok, reversedNames, reason

	}
}
