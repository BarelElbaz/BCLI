package main

import (
	"fmt"
	"os"
	"time"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: no action specified.")
		os.Exit(1)
	}

	action := os.Args[1]
	fmt.Printf("Performing %s...\n", action)

	switch action {
	case "build":
		doActionWithLoadingAnimation("Building")
	case "deploy":
		doActionWithLoadingAnimation("Deploying")
	case "run":
		doActionWithLoadingAnimation("Running")
	case "upgrade":
		doActionWithLoadingAnimation("Upgrading")
	case "install":
		doActionWithLoadingAnimation("Installing")
	default:
		fmt.Printf("Error: invalid action '%s'.\n", action)
		os.Exit(1)
	}
}

func doActionWithLoadingAnimation(action string) (bool) {
	fmt.Printf("%s:\n", action)

	// Define the loading animation
	animation := []string{
		"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷",
	}

	// Start the progress bar
	fmt.Print("[")
	for i := 0; i < 40; i++ {
		fmt.Print(" ")
	}
	fmt.Print("]\r[")
	for i := 0; i < 40; i++ {
		fmt.Print(" ")
	}
	fmt.Print("]\r")

	// Loop to update the progress bar and display the loading animation
	for i := 0; i < 100; i++ {
		index := i % len(animation)
		if i < 0 || i / 2 < 0 || 39-i/2 < 0 {
			continue
		}
		fmt.Printf("\033[1;34m%s\033[0m[%-40s]%3d%%\r", animation[index], strings.Repeat("=", i/2)+">"+strings.Repeat(" ", 39-i/2), i)
		time.Sleep(30 * time.Millisecond)
	}

	// Print the "Done!" message
	fmt.Println("\n\033[1;32mDone!\033[0m")
	return true
}
