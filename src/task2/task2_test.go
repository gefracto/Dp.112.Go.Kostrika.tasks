package task2

import (
	"testing"
)

func Test_isBigger(t *testing.T) {

}

func Benchmark_isBigger(b *testing.B) {

}

func Test_biggerS(t *testing.T) {

}

func Benchmark_biggerS(b *testing.B) {

}

func Test_smallerS(t *testing.T) {

}

func Benchmark_smallerS(b *testing.B) {

}

func Test_diagonal(t *testing.T) {

}

func Benchmark_diagonal(b *testing.B) {

}

func Test_goesIn(t *testing.T) {

}

func Benchmark_goesIn(b *testing.B) {

}

func Test_dotask(t *testing.T) {

}

func Benchmark_dotask(b *testing.B) {
	var e, e2 Envelope

	e.Side1, e.Side2 = 21, 15
	e2.Side1, e2.Side2 = 3.5, 22.3

	for i := 0; i < b.N; i++ {
		dotask(e, e2)
	}
}
