package main

import "fmt"
import "golang-book/ques3/sleep"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "from1"
			sleep.Sleep(2)
		}
	}()
	go func() {
		for {
			c2 <- "from2"
			sleep.Sleep(3)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()
	var input string
	fmt.Scanln(&input)
}
