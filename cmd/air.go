package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func RunDev() {
	airPath, err := exec.LookPath("air")
	if err != nil {
		fmt.Println("Air is not installed. Please install it using: go install github.com/cosmtrek/air@latest")
		return
	}

	if _, err := os.Stat(".air.toml"); os.IsNotExist(err) {
		fmt.Println(".air.toml' not found, running: air init ...")
		cmdInit := exec.Command(airPath, "init")
		cmdInit.Stdout = os.Stdout
		cmdInit.Stderr = os.Stderr
		err := cmdInit.Run()
		if err != nil {
			fmt.Println("Failed to run 'air init': " + err.Error())
			return
		}
		fmt.Println(".air.toml has been initialized.")
	}
	cmd := exec.Command(airPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin // biar bisa CTRL+C
	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to run air: " + err.Error())
		return
	}
}
