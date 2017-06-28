package task7

import (
	"testing"
)

func Test_getLength(t *testing.T) {
	num := 12345
	len := 5
	if getLength(num) != len {
		t.Errorf("Ожидалось %d, "+
			"получено %d", len, getLength(num))
	}
}

func Benchmark_getLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getLength(i)
	}
}

func Test_validateFileName(t *testing.T) {
	name1 := "context"
	name2 := "someother"
	err1 := validateFileName(name1)
	err2 := validateFileName(name2)

	if err1 != nil {
		t.Errorf("Функция должна обрабатывать имя context")
	}
	if err2 == nil {
		t.Errorf("Функция не должна пропускать имя, отличное от context")
	}

}

func Benchmark_validateFileName(b *testing.B) {

}

func Test_canBeRead(t *testing.T) {

}

func Benchmark_canBeRead(b *testing.B) {

}

func Test_validateHasData(t *testing.T) {

}

func Benchmark_validateHasData(b *testing.B) {

}

func Test_hasTwoNums(t *testing.T) {

}

func Benchmark_hasTwoNums(b *testing.B) {

}

func Test_hasOneNum(t *testing.T) {

}

func Benchmark_hasOneNum(b *testing.B) {

}

func Test_splitData(t *testing.T) {

}

func Benchmark_splitData(b *testing.B) {

}

func Test_validateNums(t *testing.T) {

}

func Benchmark_validateNums(b *testing.B) {

}

func Test_caseTwoNums(t *testing.T) {

}

func Benchmark_caseTwoNums(b *testing.B) {

}

func Test_caseOneNum(t *testing.T) {

}

func Benchmark_caseOneNum(b *testing.B) {

}

func Test_dotask(t *testing.T) {
}

func Benchmark_dotask(b *testing.B) {
	file := "context"
	for i := 0; i < b.N; i++ {
		dotask(file)
	}
}
