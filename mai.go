package main

import "fmt"

func main() {
	x := []int{4, 7, 3, 2, 0, 9}
	x[3] = x[1] + x[0]
	x[4] = x[2]
	fmt.Println(x[3] % x[4])
}
