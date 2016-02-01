package main

import "fmt"

type fizzBuzzer struct {
	N int
	S string
}

func main() {
	c := generate()
	c = filter(c, 3, "Fizz")
	c = filter(c, 4, "Buzz")

	for {
		num := <-c
		fmt.Printf("%d %s\n", num.N, num.S)

		if num.N >= 100 {
			break
		}
	}
}

func generate() <-chan fizzBuzzer {
	c := make(chan fizzBuzzer)
	go func() {
		for i := 1; ; i++ {
			c <- fizzBuzzer{i, ""}
		}
	}()
	return c
}

func filter(c <-chan fizzBuzzer, divisor int, label string) <-chan fizzBuzzer {
	out := make(chan fizzBuzzer)
	go func() {
		for {
			num := <-c
			if num.N%divisor == 0 {
				num.S += label
			}
			out <- num
		}
	}()
	return out
}
