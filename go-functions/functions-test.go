package main

import (
	"fmt"
	"os"
)

func multiValueReturner() (int, int) {
	return 5, 6
}

func sumAll(args ...int) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func closer() int {
	var a, b int = 1, 2

	sumFunc := func(x int, y int) int {
		return x + y
	}

	sum := sumFunc(a, b)

	return sum
}

func factorial(x int) int {
	if x <= 1 {
		return 1
	}

	return x * factorial(x-1)
}

/*
This has 3 advantages:
(1) it keeps our Close call near our Open call so it's easier to understand,
(2) if our function had multiple return statements
(perhaps one in an if and one in an else) Close will happen before both of them
(3) deferred functions are run even if a run-time panic occurs.
*/
func deferFunctionTest() {
	file, err := os.Open("../files.txt")
	if err != nil {
		fmt.Println("Error Occured while opening file")
		// handle the error here
		return
	}
	defer file.Close()
	// get the file stats
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// read the file
	bs := make([]byte, stat.Size()) //Make an array of bytes equal to the size of file//
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	fmt.Println(bs)
	str := string(bs) // convert bytes to string
	fmt.Println(str)
}

func recoverPanic() {
	defer func() {
		str := recover()
		fmt.Println("error: ", str)
	}()
	panic("Not Impemented")
}

func main() {
	//Multi Value Returner Test//
	fmt.Println("Multiple Value Returner Function Test: ")
	x, y := multiValueReturner()
	fmt.Println(x)
	fmt.Println(y)

	//Variadic Function Test//
	fmt.Println("Variadic Function Test: ")
	sum := sumAll(x, y, 3, 4)
	fmt.Println("sum:", sum)

	list := []int{1, 2, 50, 3}
	fmt.Println("Slice Sum: ", sumAll(list...))

	//Closure, declaring function inside another function//
	fmt.Println("Closure Function Test: ")
	fmt.Println("Closure Function Value: ", closer())

	//Recusion//
	fmt.Println("Recursion Function Test: ")
	fmt.Println("Factorial of 5: ", factorial(5))

	//Defer Keyword Test//
	fmt.Println("Defer keyword Test: ")
	deferFunctionTest()

	//Panic and Recover//
	fmt.Println("Panic and Recover Test: ")
	recoverPanic()
}
