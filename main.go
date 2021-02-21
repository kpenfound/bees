package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Simulate
	fmt.Println("Bees!")
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Run with argument bee, start, or watch")
		os.Exit(0)
	}

	if args[0] == "start" {
		NewWorld()
	}

	if args[0] == "bee" {
		r := NewRedis()
		id := os.Getenv("NOMAD_JOB_NAME")
		b, _ := r.GetBee(id)
		b.Lifecycle()
	}

	if args[0] == "watch" {
		r := NewRedis()
		id := "watcher"
		loc := Location{X: 0, Y: 0}
		setupInterrupt()

		for {
			m, err := r.See(loc, int(math.Max(WorldX, WorldY)), id)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			fmt.Printf("%+v\n", m)
		}
	}
	os.Exit(0)
}

func setupInterrupt() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}
