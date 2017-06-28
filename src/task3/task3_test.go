package task3

import (
	"testing"
)

func Test_halfperim(t *testing.T) {

}

func Benchmark_halfperim(b *testing.B) {

}

func Test_area(t *testing.T) {

}

func Benchmark_area(b *testing.B) {

}

func Test_getname(t *testing.T) {

}

func Benchmark_getname(b *testing.B) {

}

func Test_reverseSliceOfNames(t *testing.T) {

}

func Benchmark_reverseSliceOfNames(b *testing.B) {

}

func Test_checkTriangles(t *testing.T) {

}

func Benchmark_checkTriangles(b *testing.B) {

}

func Test_Dotask(t *testing.T) {

}

func Benchmark_Dotask(b *testing.B) {
	var aaa Triangle = Triangle{"ABC",10,7,5}
	var bbb Triangle = Triangle{"DEF",43,60,71}
	var ccc Triangle = Triangle{"GHI",143,60,100}
	var ddd []Triangle = []Triangle{aaa,bbb,ccc}

	for i:=0; i<b.N; i++ {
		Dotask(ddd)
	}
}

