package jsonToStruct

import (
	t2 "task2"
	t3 "task3"
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
