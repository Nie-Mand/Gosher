package core

import (
	"fmt"
	"time"
)

// returns the return of fn
func Took(label string) func() {
	fmt.Printf("Starting: %s\n", label)
	start := time.Now()

	return func() {
		elapsed := time.Since(start)
		fmt.Printf("%s took %s to execute\n", label, elapsed)
	}
}
