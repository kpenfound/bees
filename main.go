package main

import (
	"fmt"
	"os"
)

func main() {
	// Simulate
	fmt.Println("Bees!")
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Run with argument bee or start")
		os.Exit(0)
	}

	if args[0] == "start" {
		NewWorld()
	}

	if args[0] == "bee" {
		r := NewRedis()
		id := os.Getenv("NOMAD_JOB_ID")
		b, _ := r.GetBee(id)
		b.Lifecycle()
	}
	os.Exit(0)
}
