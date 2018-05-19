package main

import (
	"fmt"
)

func main() {
	i := 1

	// Print i's pointer
	fmt.Println(&i)

	iPointer := &i

	// Set value through i's pointer
	*iPointer = 32

	if i == 32 {
		fmt.Println("i = 32")
	}
}
