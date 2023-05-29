package main

import (
	"fmt"

	"github.com/ssukanmi/leetcode/solutions"
)

func main() {
	fmt.Println("Hello World!!!")
	input := []string{"flora", "flo"}
	ans := solutions.LongestCommonPrefix2(input)
	fmt.Println(ans)

	for i := 0; i < 10; i += 2 {
		fmt.Println(i)
	}

	output := solutions.GetConcatenation([]int{1, 3, 2, 1})
	fmt.Println(output)
}
