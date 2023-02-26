package main

import (
	"fmt"
	"os"
	"time"
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

func doActionWithLoadingAnimation(action string) {
	animation := []string{
		"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷",
	}

	fmt.Printf("%s ", action)

	for i := 0; i < 40; i++ {
		fmt.Print(animation[i%len(animation)])
		time.Sleep(100 * time.Millisecond)
		fmt.Print("\b")
	}

	fmt.Println(" Done!")
}
