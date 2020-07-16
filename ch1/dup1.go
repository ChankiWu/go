// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.

package main

import (
		"bufio"
		"fmt"
		"os"
)

func main() {
	// 声明省略 var, 注意 := 左侧如果没有声明新的变量，就产生编译错误
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan(){
		counts[input.Text()]++
	}
	
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

