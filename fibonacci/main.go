package main

import "fmt"

func fibo(x uint) uint {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return fibo(x-1) + fibo(x-2)
}
func main() {
	var x uint
	fmt.Println("Enter the nth number in fibonacci series: ")
	fmt.Scan(&x)
	f := fibo(x)
	fmt.Println("The", x, "th number in fibonacci series is : ", f)
}
