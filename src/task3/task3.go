// S = sqrt(p*(p-a)(p-b)(p-c))
// p = (a+b+c)/2
 
package task3
 
import (
    "sort"
    "math"
    "fmt"
)
 
type Triangle struct {
    Name string
    A float64
    B float64
    C float64
}
 
//type Triangle []Triangle
 
func calculate(T Triangle) (string, float64) {
    p := (T.A + T.B + T.C)/2
    s := math.Sqrt(p * (p - T.A) * (p - T.B) * (p - T.C))
    return T.Name, s
}
 
func SortTriangles(sl []Triangle) []string {
    var names []string
    m := make(map[string]float64)
    for i := 0; i < len(sl); i++ {
        k, v := calculate(sl[i])
        m[k] = v
    }
 
    //Inverting maps
    invMap := make(map[float64]string, len(m))
    for k,v := range m {
        invMap[v] = k
    }
 
     //Sorting
    sortedKeys := make([]float64, len(invMap))
    var i int = 0
    for k := range invMap {
        sortedKeys[i] = k
        i++
    }
    sort.Slice(sortedKeys, func(i, j int) bool { return sortedKeys[i] < sortedKeys[j] })
 
    for i := range sortedKeys {
        names = append(names, invMap[float64(i)])
        fmt.Println(invMap[float64(i)])
    }
    return names
}