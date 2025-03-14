package main

import (
	"GoIntro/tutorials"
	"fmt"
)

func main() {
	tutorials.Hello()
	fmt.Println(tutorials.Add(12, 6))
	result, err := tutorials.Divide(12, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	tutorials.RandomNumber()
}
