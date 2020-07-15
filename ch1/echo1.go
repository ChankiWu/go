// Echo1 prints its command-line arguments.

package main

import (
	"fmt"
	"os"
	// "strings"
)

func main() {

	// 等价声明

	/* 	s := ""
	var s string
	var s = ""
	var s string = "" */

	var s, sep string

	// os.Args的第一个元素：os.Args[0]，是命令本身的名字
	// s[m:n]形式的切片表达式，产生从第m个元素到第n-1个元素的切片

	// 1.
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }

	// 2.

	// 每次循环迭代，range产生一对值；索引以及在该索引处的元素值。
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	// fmt.Println("range: ")

	// for i, j := range "keep" {
	// 	fmt.Println(i)
	// 	fmt.Println(j)
	// }

	/*
		range:
			0
			107 // Unicode
			1
			101
			2
			101
			3
			112
	*/

	// 3.
	// fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])

	fmt.Println("The length of the Arguments is: ", len(os.Args))
	fmt.Println("The first argument is: ", os.Args[0])
	fmt.Println(s)
}
