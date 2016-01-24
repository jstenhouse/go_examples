package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("Evolution!")
	case "windows":
		fmt.Println("Doors!")
	default:
		fmt.Println("IDK!")
	}
}
