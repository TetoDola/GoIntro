package tutorials

import (
	"fmt"
	"time"
)

func Allocate() {
	n := 100000000
	testslice := []int{}
	testslice2 := make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testslice, n))
	fmt.Printf("Total time with preallocation: %v", timeLoop(testslice2, n))
	//
}

func timeLoop(slice []int, n int) time.Duration {
	t0 := time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	t1 := time.Since(t0)
	return t1

}
