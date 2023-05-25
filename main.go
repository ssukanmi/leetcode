package main

import (
	"fmt"

	"github.com/ssukanmi/leetcode/solutions"
)

func main() {
	fmt.Println("Hello World!!!")
	input := []string{"flora", "flo"}
	ans := solutions.LongestCommonPrefi2(input)
	fmt.Println(ans)
}
