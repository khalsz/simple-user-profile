package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Println(nums, " ")

	total := 0

	for _, num := range nums {

		total += num
	}
	fmt.Println(total)
}

// func main() {
// 	// i := 1
// 	// for i <= 3 {
// 	// 	fmt.Println(i)
// 	// 	i = i + 1
// 	// }

// 	// for j := 0; j < 3; j++ {
// 	// 	fmt.Println(j)
// 	// }

// 	// for i := range 3 {
// 	// 	fmt.Println("range:", i)
// 	// }

// 	// for n := range 6 {
// 	// 	if n%2 == 0 {
// 	// 		continue
// 	// 	}
// 	// 	fmt.Println(n)
// 	// }

// 	// var a [5]int

// 	// b := [5]int{1, 2, 3, 4, 5}
// 	// fmt.Println(b)

// 	// c := [...]int{1, 2, 3, 4, 5, 6, 6, 7}
// 	// fmt.Println(c)

// 	// var s []string
// 	// fmt.Println("uninit:", s, s == nil, len(s) == 0)

// 	// s = make([]string, 3)
// 	// fmt.Println("empt:", s, "len:", len(s), "cap:", cap(s))

// 	// s[0] = "a"
// 	// s[1] = "b"
// 	// s[2] = "c"
// 	// fmt.Println(s[2])
// 	// fmt.Println(len(s))

// 	// c := make([]string, len(s))
// 	// copy(c, s)
// 	// fmt.Println("cpy:", c)

// 	// l := s[2:5]
// 	// fmt.Println("sl1", l)

// 	// m := make(map[string]int)

// 	// m["k1"] = 1
// 	// m["k2"] = 2

// 	// fmt.Println("map", m)
// 	// fmt.Println("k1", m["k1"])

// 	// v1 := m["k1"]
// 	// fmt.Println("v1", v1)

// 	// delete(m, "k2")
// 	// fmt.Println("map:", m)

// 	// clear(m)
// 	// fmt.Println("map", m)

// 	// kk, pr := m["k1"]

// 	// fmt.Println("pr:", pr, kk)

// 	// n := map[string]int{"foo": 1, "bar": 2}
// 	// fmt.Println("n", n)

// 	// n2 := map[string]int{"foo": 1, "bar": 2}
// 	// if maps.Equal(n, n2) {
// 	// 	fmt.Println("n == n2")
// 	// }

// 	// res := plus(1, 2)
// 	// fmt.Println("1+2 =", res)

// 	// res = plusPlus(1, 2, 3)
// 	// fmt.Println("1+2+3 =", res)

// 	sum(1, 2, 3, 4)

// }
