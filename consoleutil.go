package main

import (
	"os"
	"os/exec"
	"runtime"
)

func consoleClear() {
	currOS := runtime.GOOS
	switch currOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
