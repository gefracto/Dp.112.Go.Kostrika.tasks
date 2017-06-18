package task2

import "math"

/*
func bigger(a, b, c, d int) int {
	ab := a*b
	cd := c*d
	if ab > cd {
		return ab
	} else if ab < cd {
		return cd
	} else {
		return 0
	}
} */

func Envelope3(a, b, c, d int) float64 {
	if (a > c) && (b > d) {
		return 2
	} else if (a < c) && (b < d) {
		return 1
	} else if (a == c) && (b == d) {
		return 0
	}

	if (math.Sqrt(float64(a*a+b*b)) - float64(d)) > float64(c) {
		return 2
	} else {
		return 1
	}
}

type Envelope struct {
	a, b float64
}

func DoTask2(e1 Envelope, e2 Envelope) int {

}