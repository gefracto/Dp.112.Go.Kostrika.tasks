package task3

import (
	"errors"
	"math"
	"regexp"
	"sort"
)

type T3 struct {
	SliceOfTriangles []Triangle
}

func (T *T3) Dotask3() (err error, data []string) {
	err, data = dotask(T.SliceOfTriangles)
	return
}

type Triangle struct {
	Name string
	A    float64
	B    float64
	C    float64
}

func (T *Triangle) halfperim() float64 {
	return (T.A + T.B + T.C) / 2
}

func (T *Triangle) area(p float64) float64 {
	S := math.Sqrt(p * (p - T.A) * (p - T.B) * (p - T.C))
	return S
}

func (T *Triangle) getname() string {
	return T.Name
}

func reverseSliceOfNames(slice []string) (newSlice []string) {
	for i := len(slice) - 1; i >= 0; i-- {
		newSlice = append(newSlice, slice[i])
	}
	return newSlice
}

func checkTriangles(T []Triangle) (ok bool, err error) {
	for _, v := range T {
		pattern := "\\A[A-Z]{1}[0-9]?[A-Z]{1}[0-9]?[A-Z]{1}[0-9]?$"
		if ok, _ := regexp.MatchString(pattern, v.Name); ok == false {
			return ok, errors.New("Имя треугольника " + v.Name +
				" должно состоять из имен трех вершин\n")
		}
		if v.A >= v.B+v.C || v.A <= 0 {
			return ok, errors.New("Значение стороны \"А\" треугольника \"" +
				v.Name + "\"\nне должно быть" +
				" больше суммы сторон \"B\" и \"C\", " +
				"\nа также не может быть меньше или равно нулю\n")
		}
		if v.B >= v.A+v.C || v.B <= 0 {
			return ok, errors.New("Значение стороны \"B\" треугольника \"" +
				v.Name + "\"\nне должно быть" +
				" больше суммы сторон \"A\" и \"C\", " +
				"\nа также не может быть меньше или равно нулю\n")

		}
		if v.C >= v.A+v.B || v.C <= 0 {
			return ok, errors.New("Значение стороны \"C\" треугольника \"" + v.Name +
				"\"\nне должно быть больше суммы сторон \"B\" и \"A\", " +
				"\nа также не может быть меньше или равно нулю\n")
		}
	}
	return true, err
}

func dotask(T []Triangle) (err error, data []string) {
	if ok, err := checkTriangles(T); ok == false {
		return err, nil

	}
	var names []string
	var squares []float64
	var indexes []int
	var sortedSquares []float64
	var sortedNames []string

	for i := range T {
		names = append(names, T[i].getname())
		squares = append(squares, T[i].area(T[i].halfperim()))
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
	return err, reversedNames

}
