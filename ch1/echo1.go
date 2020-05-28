// Echo1 prints its command-line arguments.

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	// os.Args的第一个元素：os.Args[0]，是命令本身的名字
	// s[m:n]形式的切片表达式，产生从第m个元素到第n-1个元素的切片

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println("The length of the Arguments is: ", len(os.Args))
	fmt.Println("The first argument is: ", os.Args[0])
	fmt.Println(s)
}
