package main

import (
	t7 "task7"
	//t5 "task5"
	//t6 "task6"
	//"fmt"
	//t6 "task6"
	//t4 "task4"
	"fmt"
	//"strconv"
	//"strings"
	//t2 "task2"
	//t1 "task1"
	"io/ioutil"
)


func main() {
	//fmt.Println(t1.ChessBoardNice(6, 6, " o ", "   "))
	//fmt.Println(t1.ChessBoardVeryNice(19, 4))

	//t := t4.Palindrom(44565478)
	//fmt.Printf(t[:len(t)])
	//fmt.Printf(t[:1])
	/*
		for k,v := range t {
			fmt.Println(k, string(v))
		}
		if strings.ContainsAny(string(t), "4") {
			fmt.Println(true)
		} else {
			fmt.Println(false)
		} */

	//fmt.Println(t2.Envelope(22,2,15,20))
	//fmt.Println(t6.NumberSequence(10, 124))
	//fmt.Println(t7.Fibo())
	//s := t7.Fibo("../src/src/task7/context")

	/*
	s := t7.Fibo("4")
	for _, v := range s {
		fmt.Println(v)
	}
	*/

	//fmt.Println(t6.NumberSequence(05, 7643))

	//m := make([]Triangle)
	//m[0] = Triangle{"DFG",3,4,5}
	/* m := make([]t3.Triangle, 0)
	m = append(m, t3.Triangle{Name: "FDF", A: 2, B: 3, C: 3})
	m = append(m, t3.Triangle{Name: "ABC", A: 1, B: 2, C: 3})
	m = append(m, t3.Triangle{Name: "FHH", A: 34, B: 45, C: 23})
	fmt.Println(m)
	fmt.Println(t3.SortTriangles(m)) */

	//fmt.Println(t5.LuckyTickets(0, 1))

	//t6.DoTask6(1020, 2345)
	//fmt.Println(t2.DoTask2(t2.Envelope{21,15}, t2.Envelope{3.5,22.3}))
	/*s := make([]t3.Triangle, 1)
	s = append(s, t3.Triangle{"sff",10,10,10},t3.Triangle{"ghj",10,10,10}, t3.Triangle{"sre",34,43,20}, t3.Triangle{"abc",2,3,2})
	fmt.Println(t3.DoTask3(s))


	var a []float64
	var b []float64
	a = append(a, 1,2,3,4,5)
	b = append(b, a...)
	fmt.Println(a,b)
	//fmt.Println(t2.DoTask2(t2.Envelope{15,21}, t2.Envelope{3.5, 22.3}))

	//fmt.Println(t4.DoTask4(121212))
*/

	fmt.Println(t7.DoTask7("context"))
	fmt.Println(t7.DoTask7("context2"))

	contents, _ := ioutil.ReadFile("context")
	str:= string(contents)
	fmt.Println(str, len(str))
}
