package main

import "fmt"

func main() {
	printHelloWorld(10)
}

func printHelloWorld(j int) {
	i := 0
	for i <= j {
		fmt.Printf("Hello World %d times\n", i)
		i++
	}
}
