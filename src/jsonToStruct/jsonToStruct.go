package jsonToStruct

import (
	t2 "task2"
	t3 "task3"
	"io/ioutil"
)

type Data struct {
	T1 struct {
		Width, Height int
		Symbol string
	}
	T2 struct {
		Env1, Env2 t2.Envelope
	}
	T3 struct {
		SliceOfTriangles []t3.Triangle
	}
	T4 struct {
		Number int
	}
	T5 struct {
		Min, Max int
	}
	T6 struct {
		Length, MaxSquare int
	}
	T7 struct {
		File string
	}
}

func getFile(f string) (string, bool) {
	ok := false
	contents, _ := ioutil.ReadFile(f)
	str := string(contents)
	return str, ok
}


