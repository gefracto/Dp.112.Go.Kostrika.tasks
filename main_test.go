package main

import (
	"github.com/gefracto/kostrika-go-tasks/src/tools"
	"testing"
)

/*func TestMain(m *testing.M) {
	main()
}*/

func Benchmark_runall(b *testing.B) {

	_, data := tools.GetData("data.json")

	for j := 0; j < b.N; j++ {
		for i := 0; i < 100; i++ {
			runall(data)
		}
	}
}
