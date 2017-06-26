package task2

import (
	"errors"
	"math"
)

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

func (e *Envelope) isBigger(e2 *Envelope) (b bool) {
	A, B := e.Side1, e.Side2
	C, D := e2.Side1, e2.Side2

	if A*B > C*D {
		return true
	}
	return
}

func (e *Envelope) biggerS() (f float64) {
	if e.Side1 >= e.Side2 {
		return e.Side1

	}
	return e.Side2
}

func (e *Envelope) smallerS() (f float64) {
	if e.Side1 < e.Side2 {
		return e.Side1

	}
	return e.Side2
}

func (e *Envelope) diagonal() (f float64) {
	f = float64(math.Sqrt(e.Side1*e.Side1 + e.Side2*e.Side2))
	return
}

func (e *Envelope) goesIn(e2 *Envelope) bool {
	x := e2.isBigger(e)
	y := e2.diagonal()-e.biggerS() >= e.smallerS()
	return x && y
}

func Dotask(e1 Envelope, e2 Envelope) (err error, data int) {
	A, B := e1.Side1, e1.Side2
	C, D := e2.Side1, e2.Side2

	areEqual := (A == C || A == D) && (B == C || B == D)

	if A <= 0 || B <= 0 || C <= 0 || D <= 0 {

		return errors.New("Ошибка!\nСторона конверта " +
			"не может быть меньше или равно 0"), 0

	} else if areEqual {
		return nil, 0

	} else if e1.goesIn(&e2){
		return nil, 1

	}
	return nil, 2
}
