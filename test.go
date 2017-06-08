package main

import (
	"fmt"
	"os"
)

func run() (err error) {
	// add a new comment line
	// 日本語普通に行けるはず
	fmt.Println("hi")
	fmt.Println("lo")

	return
}

func main() {
	var err error

	err = run()

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
