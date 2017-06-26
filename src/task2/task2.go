package task2

import (
	"math"
	"errors"
)

// укоротить длинные куски кода

type T2 struct {
	E1 Envelope
	E2 Envelope
}

func (T *T2) Dotask2() (err error, data interface{}) {

	err, data = Dotask(T.E1, T.E2)
	return
}

type Envelope struct {
	Side1, Side2 float64
}

func (e *Envelope) isBiggerThan(e2 *Envelope) (b bool) {
	if e.Side1*e.Side2 > e2.Side1*e2.Side2 {
		return true
	}
	return false
}

func (e *Envelope) biggerSide() (f float64) {
	if e.Side1 >= e.Side2 {
		return e.Side1

	} else if e.Side1 < e.Side2 {
		return e.Side2
	}

	return
}

func (e *Envelope) smallerSide() (f float64) {
	if e.Side1 < e.Side2 {
		return e.Side1

	} else if e.Side2 <= e.Side1 {
		return e.Side2
	}

	return
}

func (e *Envelope) findDiagonal() (f float64) {
	f = float64(math.Sqrt(e.Side1*e.Side1 + e.Side2*e.Side2))
	return
}

func Dotask(e1 Envelope, e2 Envelope) (err error, data int) {
	A, B := e1.Side1, e1.Side2
	C, D := e2.Side1, e2.Side2
	if A <= 0 || B <= 0 || C <= 0 || D <= 0 {
		return errors.New("Ошибка!\nСторона конверта не может быть меньше или равно 0"), 0

	} else if (A == C || A == D) && (B == C || B == D) {
		return nil, 0

	} else if e1.isBiggerThan(&e2) && (e1.findDiagonal()-e2.biggerSide() >= e2.smallerSide()) {
		return nil, 2

	} else if e2.isBiggerThan(&e1) && (e2.findDiagonal()-e1.biggerSide() >= e1.smallerSide()) {
		return nil, 1
	}

	return nil, 0
}
