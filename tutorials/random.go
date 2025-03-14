package tutorials

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomNumber() {
	rand.Seed(time.Now().UnixNano())
	minimum := 10
	maximum := 30
	num := rand.Intn(maximum-minimum+1) + minimum
	var guess int

	for num != guess {
		fmt.Println("Guess the number:")
		fmt.Scan(&guess) // needs error handling
	}
}
