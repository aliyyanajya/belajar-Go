package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("Data not complete!")
	if err != nil {
		fmt.Println(err)
	}
}
