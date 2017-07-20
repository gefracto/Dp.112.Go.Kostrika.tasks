package task1

import (
	"testing"
)

func Test_makeRow(t *testing.T) {
	b := true
	s := "*"
	w := 4
	res := makeRow(b, s, w)
	if res != "* * " {
		t.Error("Ожидался результат \"* * \", получено ", res)
	}
}

func Benchmark_makeRow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeRow(true, "*", 8)
	}
}

func Test_validate(t *testing.T) {
	var valid []int = []int{1, 3, 3, 1, 3}
	var invalid []int = []int{0, -1, -2, -299, -323245}

	for i := 0; i < len(valid); i++ {
		if _, err := validate(valid[i], valid[i]); err != nil {
			t.Error(err)
		}
	}

	for j := 0; j < len(invalid); j++ {
		if _, err := validate(invalid[j], invalid[j]); err == nil {
			t.Error("Ошибка! Нет реакции на недопустимое значение: ", invalid[j])
		}
	}
}

func Benchmark_validate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validate(i, i)
	}
}
