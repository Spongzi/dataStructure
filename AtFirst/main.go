package main

import "fmt"

// 汉诺塔问题

var total = 0

func tower(n int, a string, b string, c string) {
	if n == 1 {
		total = total + 1
		fmt.Println(a, "->", c)
		return
	}
	tower(n-1, a, c, b)
	total = total + 1
	fmt.Println(a, "->", c)
	tower(n-1, b, a, c)
}

func main() {
	n := 64  // 64个盘子
	a := "a" // 杆子A
	b := "b" // 杆子B
	c := "c" // 杆子C
	tower(n, a, b, c)
	// 当n=1时，移动次数为1
	// 当n=2时，移动次数为3
	// 当n=3时，移动次数为7
	// 当n=4时，移动次数为15
	fmt.Println(total)
}
