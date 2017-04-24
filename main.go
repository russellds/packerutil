package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("PACKER_LOG", "1")
	fmt.Println("\nPACKER_LOG:", os.Getenv("PACKER_LOG"))
	os.Setenv("PACKER_LOG_PATH", "packer.log")
	fmt.Println("\nPACKER_LOG_PATH:", os.Getenv("PACKER_LOG_PATH"))

	/*
		fmt.Println("\n")
		for _, e := range os.Environ() {
			fmt.Println(e)
		}
	*/
}
