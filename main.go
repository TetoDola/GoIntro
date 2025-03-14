package main

import (
	"GoIntro/tutorials"
	"fmt"
)

func main() {
	tutorials.Hello()
	fmt.Println(tutorials.Add(12, 6))
	result, err := fmt.Println(tutorials.Divide(12, 0))
	if err != nil {
		println(err)
	} else {
		println(result)
	}
}
