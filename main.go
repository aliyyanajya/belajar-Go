package main

import "fmt"

func main() {
	fmt.Println("Hai teman-teman?")

	foo()

	fmt.Println("Sesuatu yang lain")

	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}

func foo() {
	fmt.Println("Saya di dalam foo")
}
