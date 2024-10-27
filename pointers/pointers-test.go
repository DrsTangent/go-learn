package main

import "fmt"

func swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}

func main() {
	x := 1
	y := 2
	fmt.Println("x: ", x, ", y: ", y)

	swap(&x, &y)

	fmt.Println("x: ", x, ", y: ", y)
}
