package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	println("Demo reverse with Go Slice")
	a := [...]int{0, 1, 2, 3, 4, 5}
	println("Before:")
	fmt.Println(a)
	reverse(a[:])
	println("After:")
	fmt.Println(a)
}
