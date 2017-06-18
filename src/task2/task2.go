package task2

import "math"

type Envelope struct {
	Side1, Side2 float64
}

func (e1 *Envelope) isBiggerThan(e2 *Envelope) (b bool) {
	if e1.Side1*e1.Side2 > e2.Side1*e2.Side2 {
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

func DoTask2(e1 Envelope, e2 Envelope) int {
	if (e1.Side1 == e2.Side1 || e1.Side1 == e2.Side2) && (e1.Side2 == e2.Side1 || e1.Side2 == e2.Side2) {
		return 0
	}
	if e1.isBiggerThan(&e2) && (e1.findDiagonal()-e2.biggerSide() >= e2.smallerSide()) {
		return 2
	} else if e2.isBiggerThan(&e1) && (e2.findDiagonal()-e1.biggerSide() >= e1.smallerSide()) {
		return 1
	}
	return 0
}
