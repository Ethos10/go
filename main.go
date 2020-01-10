package main

import "fmt"

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	x := 1
	y := 2
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}
