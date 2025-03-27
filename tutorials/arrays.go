package tutorials

import (
	"fmt"
	"time"
)

func Array() {
	start := time.Now()
	var array [100000000000]uint8
	for i := 0; i < len(array); i++ {
		array[i] = 1
	}
	end := time.Since(start)
	fmt.Println(end)

}
