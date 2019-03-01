package main

import "fmt"

func main() {
	fmt.Println("Hai teman-teman?")

	foo()

	fmt.Println("Sesuatu yang lain")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func bar() {
	fmt.Println("Keluar Program !")
}

func foo() {
	fmt.Println("Saya di dalam foo")
}
