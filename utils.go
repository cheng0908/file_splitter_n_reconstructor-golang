package main

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func displayProgressBar(current, total int) {
	progress := float64(current) / float64(total) * 100
	barLength := 40
	currentBar := int(progress / 100 * float64(barLength))

	fmt.Printf("\r[")
	for i := 0; i < currentBar; i++ {
		fmt.Print("=")
	}
	for i := currentBar; i < barLength; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("] %.2f%% (%d/%d)", progress, current, total)
}
