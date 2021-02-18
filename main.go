package main

import (
	"fmt"
	"os"
)

func main() {
	// Simulate
	fmt.Println("Bees!")
	w := NewWorld()
	w.Simulate()
	os.Exit(0)
}
