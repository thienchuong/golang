package main

import (
	"fmt"
	"os"
)

func main() {
	const name, age = "kim", 222222
	n, err := fmt.Fprintf(os.Stdout, "%s is %d years old.\n", name, age)

	//n and err return values from Fprintf are those returned by the underlaying io.Writer

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
	}
	fmt.Printf("%d bytes written.\n", n)
}
