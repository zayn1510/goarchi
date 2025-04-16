package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func Upgrade() {
	binaryName := "goarchi"
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}

	goPath, err := exec.LookPath("go")
	if err != nil {
		fmt.Println("❌ Go compiler not found in PATH.")
		return
	}

	// Rebuild the binary
	build := exec.Command(goPath, "build", "-o", binaryName, "cli/main.go")
	build.Stdout = os.Stdout
	build.Stderr = os.Stderr
	err = build.Run()
	if err != nil {
		fmt.Println("❌ Failed to build:", err)
		return
	}

	// Handle Windows separately
	if runtime.GOOS == "windows" {
		fmt.Println("✅ Build completed. Please manually replace", binaryName, "in a directory listed in your PATH.")
		return
	}

	// Target binary path on Unix systems
	dest := "/usr/local/bin/goarchi"

	if isBinaryRunning(dest) {
		fmt.Println("❌ The binary is currently running. Please stop it first using: pkill -f goarchi")
		return
	}

	err = os.Rename(binaryName, dest)
	if err != nil {
		// Fallback to manual copy if rename fails
		err = copyFile(binaryName, dest)
		if err != nil {
			fmt.Println("❌ Failed to upgrade binary:", err)
			fmt.Println("🔧 Try running: sudo env \"PATH=$PATH\" go run cli/main.go upgrade")
			return
		}
		_ = os.Remove(binaryName)
	}

	fmt.Println("🚀 Goarchi has been successfully upgraded to the latest version!")
}
