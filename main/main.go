package main

import (
	//t3 "task3"
	//t5 "task5"
	//t6 "task6"
	//"fmt"
	//t6 "task6"
	t4 "task4"
	"fmt"
	//"strconv"
	//"strings"
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



	fmt.Println(t4.DoTask4(121212))
	s := "hello"
	n := "olleh"
	for i,j := range s {
		fmt.Println(i,string(j))
	}
	if "h" == string(n[len(s)-1]) {
		fmt.Println("HFEHEFHUF")
	}
}
