package main

import (
	"fmt"
	"time"
)

// ch4.1 ch4.2
func array(){
	var a [3]int = [3]int{1, 2, 3}             // array of 3 integers
	// fmt.Println(a[0])        // print the first element
	// fmt.Println(a[len(a)-1]) // print the last element, a[2]

	for i := 0; i < 5; i++ {
		fmt.Printf("%d, ", i)
	}

	fmt.Println()

	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
		/* output
		0 1
		1 2
		2 3
		*/
	}

	// Print the elements only.
	// for _, v := range a {
	// 	fmt.Printf("%d\n", v)
	// }

	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"

	// slice
	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...) // append the slice x
	fmt.Println(x)      // "[1 2 3 4 5 6 1 2 3 4 5 6]"
}
/*
	var s []int    // len(s) == 0, s == nil
	s = nil        // len(s) == 0, s == nil
	s = []int(nil) // len(s) == 0, s == nil
	s = []int{}    // len(s) == 0, s != nil
	一个零值的slice等于nil
*/

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// ch4.3
func makeMap(){
	// 在向map存数据前必须先创建map。
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34
	if val, ok := ages["alice"]; ok {
		fmt.Println(val) // "31"
	}
	
	delete(ages, "alice") // remove element ages["alice"]

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
		// output: charlie 34
	}
}

// ch4.4	
type tree struct {
	value       int
	left, right *tree
}

type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

type Point struct{ X, Y int }

func main(){
	array()

	// test for reverse
	// s := []int{0, 1, 2, 3, 4, 5}
	// // Rotate s left by two positions. : -> i-j-1
	// reverse(s[:2])
	// fmt.Println(s) // [1 0 2 3 4 5]
	// reverse(s[2:])
	// fmt.Println(s) // [1 0 5 4 3 2]
	// reverse(s)
	// fmt.Println(s) // [2 3 4 5 0 1]

	makeMap()

	var dilbert Employee
	dilbert.Salary = 100

	p := Point{1, 2}
	pp := &Point{1, 2} // equals to pp := new(Point); *pp = Point{1, 2}
	fmt.Println(dilbert) // default value {0   0001-01-01 00:00:00 +0000 UTC  100 0}
	fmt.Println(p)	// {1 2}
	fmt.Println(pp) // &{1 2}
} 