package main

import "fmt"

func main() {

	c := make(chan int, 3)
	numbers := []int{1, 2, 3, 4, 5}

	go func() {
		for _, num := range numbers {
			c <- num
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
