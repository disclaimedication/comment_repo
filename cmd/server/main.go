package main

import (
	"fmt"
)

func Run() error {
	fmt.Println("This is the run function.")
	return nil
}
func main() {
	fmt.Println("Welcome to the new adventure!!")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
