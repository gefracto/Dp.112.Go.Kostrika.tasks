package task2

import (
	"encoding/json"
	"errors"
	"github.com/gefracto/kostrika-go-tasks/src/tools"
	"math"
)

type T2 struct {
	E1 Envelope
	E2 Envelope
}

func init() {
	tools.RememberMe(2, Dotask)
}

func Dotask(js []byte) (err error, data []byte) {

	var t []Envelope
	json.Unmarshal(js, &t)
	// fmt.Println(t[0], t[1])

	A, B := t[0].Width, t[0].Height
	C, D := t[1].Width, t[1].Height

	areEqual := (A == C || A == D) && (B == C || B == D)

	if A <= 0 || B <= 0 || C <= 0 || D <= 0 {

		return errors.New("Ошибка!\nСторона конверта " +
			"не может быть меньше или равно 0"), []byte("0")

	} else if areEqual {
		return nil, []byte("0")

	} else if t[0].goesIn(&t[1]) {
		return nil, []byte("1")

	} else if t[1].goesIn(&t[0]) {
		return nil, []byte("2")
	}
	return nil, []byte("1")
}

type Envelope struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func (e *Envelope) isBigger(e2 *Envelope) (b bool) {
	A, B := e.Width, e.Height
	C, D := e2.Width, e2.Height

	if A*B > C*D {
		return true
	}
	return
}

func (e *Envelope) biggerS() (f float64) {
	if e.Width >= e.Height {
		return e.Width

	}
	return e.Height
}

func (e *Envelope) smallerS() (f float64) {
	if e.Width < e.Height {
		return e.Width

	}
	return e.Height
}

func (e *Envelope) diagonal() (f float64) {
	f = float64(math.Sqrt(e.Width*e.Width + e.Height*e.Height))
	return
}

func (e *Envelope) goesIn(e2 *Envelope) bool {
	x := e2.isBigger(e)
	y := e2.diagonal()-e.biggerS() >= e.smallerS()
	return x && y
}

// func dotask(e1 Envelope, e2 Envelope) (err error, data int) {
// 	A, B := e1.Width, e1.Height
// 	C, D := e2.Width, e2.Height

// 	areEqual := (A == C || A == D) && (B == C || B == D)

// 	if A <= 0 || B <= 0 || C <= 0 || D <= 0 {

// 		return errors.New("Ошибка!\nСторона конверта " +
// 			"не может быть меньше или равно 0"), 0

// 	} else if areEqual {
// 		return nil, 0

// 	} else if e1.goesIn(&e2) {
// 		return nil, 1

// 	}
// 	return nil, 2
// }
