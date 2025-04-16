package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func RunInstall() {
	binaryName := "goarchi"
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}

	// Cari path binary `go`
	goPath, err := exec.LookPath("go")
	if err != nil {
		fmt.Println("❌ Go compiler tidak ditemukan di PATH.")
		return
	}

	// Build binary
	cmd := exec.Command(goPath, "build", "-o", binaryName, "cli/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("❌ Gagal build:", err)
		return
	}

	// Pindahkan binary ke path global
	if runtime.GOOS == "windows" {
		fmt.Println("✅ Build berhasil. Silakan pindahkan", binaryName, "ke direktori dalam PATH Windows kamu secara manual.")
		return
	}

	// Lokasi tujuan (Linux/macOS)
	dest := "/usr/local/bin/goarchi"
	err = os.Rename(binaryName, dest)
	if err != nil {
		// Jika gagal rename (mungkin beda mount), lakukan copy manual
		err = copyFile(binaryName, dest)
		if err != nil {
			fmt.Println("❌ Gagal memindahkan binary:", err)
			fmt.Println("🔧 Coba jalankan dengan: sudo env \"PATH=$PATH\" go run cli/main.go install")
			return
		}
		// Hapus file setelah copy
		_ = os.Remove(binaryName)
	}

	fmt.Println("✅ Goarchi berhasil terinstall di /usr/local/bin/goarchi")
}
func isBinaryRunning(path string) bool {
	out, err := exec.Command("lsof", path).Output()
	if err != nil {
		return false
	}
	return len(out) > 0
}
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func(in *os.File) {
		err := in.Close()
		if err != nil {
			fmt.Printf("error closing file: %v", err)
		}
	}(in)

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			fmt.Printf("error closing file: %v", err)
		}
	}(out)

	_, err = in.Stat()
	if err != nil {
		return err
	}

	_, err = out.ReadFrom(in)
	if err != nil {
		return err
	}

	return out.Chmod(0755)
}
