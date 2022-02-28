package main

import "fmt"

// 阶乘

func Rescuvie(n int, a int) int {
	if n == 0 {
		return a
	}
	return Rescuvie(n-1, a*n)
}

func F(n, a1, a2 int) int {
	if n == 0 {
		return a1
	}
	return F(n-1, a2, a1+a2)
}

func main() {
	fmt.Println("尾递归", Rescuvie(5, 1))
	fmt.Println("斐波那契数列", F(5, 1, 1))
}
