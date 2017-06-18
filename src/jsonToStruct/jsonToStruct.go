package jsonToStruct

import (
	t2 "task2"
	t3 "task3"
)

type Task1 struct {
	Width, Height int
	Symbol string
}

type Task2 struct {
	Env1, Env2 t2.Envelope
}

type Task3 struct {
	SliceOfTriangles []t3.Triangle
}

type Task4 struct {
	Number int
}

type Task5 struct {
	Min, Max int
}

type Task6 struct {
	Length, MaxSquare int
}

type Task7 struct {
	File string
}

type Data struct {
	t1 *Task1
	t2 *Task2
	t3 *Task3
	t4 *Task4
	t5 *Task5
	t6 *Task6
	t7 *Task7
}
