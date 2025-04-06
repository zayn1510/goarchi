package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Gunakan format: go run cmd/make.go controller NamaController")
		return
	}

	command := os.Args[1]
	name := os.Args[2]

	switch command {
	case "controller":
		createController(name)
	case "model":
		createModel(name)
	case "service":
		createService(name)
	case "request":
		createRequest(name)
	case "resource":
		createResponse(name)
	case "all":
		createAll(name)
	case "seeder":
		createSeeder(name)
	case "create-seed":
		runSeeder(name)
	default:
		fmt.Println("Perintah tidak dikenal:", command)
	}
}

func createAll(name string) {
	createModel(name)
	createService(name)
	createRequest(name)
	createResponse(name)
	createController(name)
}

func createController(name string) {
	basePath := "app/controllers/"
	parts := strings.Split(name, "/")

	var folderPath, controllerName string

	if len(parts) > 1 {
		folderPath = basePath + strings.Join(parts[:len(parts)-1], "/")
		controllerName = parts[len(parts)-1]
	} else {
		folderPath = basePath
		controllerName = parts[0]
	}

	fileName := strings.ToLower(controllerName) + "_controller.go"
	fullPath := folderPath + "/" + fileName

	// Buat folder jika belum ada
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.MkdirAll(folderPath, os.ModePerm)
	}

	// Cek apakah file sudah ada
	if _, err := os.Stat(fullPath); err == nil {
		fmt.Println("File controller sudah ada:", fullPath)
		return
	}

	// Template isi file controller
	content := fmt.Sprintf(`package controllers

import "fmt"

type %sController struct {}

func (c *%sController) Index() {
	fmt.Println("Controller %s berjalan...")
}
`, controllerName, controllerName, controllerName)

	// Simpan file
	err := os.WriteFile(fullPath, []byte(content), 0644)
	if err != nil {
		fmt.Println("Gagal membuat controller:", err)
		return
	}

	fmt.Println("Controller berhasil dibuat di:", fullPath)
}
func createSeeder(name string) {
	basePath := "app/seeders/"
	parts := strings.Split(name, "/")

	var folderPath, controllerName string

	if len(parts) > 1 {
		folderPath = basePath + strings.Join(parts[:len(parts)-1], "/")
		controllerName = parts[len(parts)-1]
	} else {
		folderPath = basePath
		controllerName = parts[0]
	}

	fileName := strings.ToLower(controllerName) + "_seeder.go"
	fullPath := folderPath + "/" + fileName

	// Buat folder jika belum ada
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.MkdirAll(folderPath, os.ModePerm)
	}

	// Cek apakah file sudah ada
	if _, err := os.Stat(fullPath); err == nil {
		fmt.Println("File seeder sudah ada:", fullPath)
		return
	}

	// Template isi file controller
	content := fmt.Sprintf(`package main
`)

	// Simpan file
	err := os.WriteFile(fullPath, []byte(content), 0644)
	if err != nil {
		fmt.Println("Gagal membuat seeder:", err)
		return
	}

	fmt.Println("Seeder berhasil dibuat di:", fullPath)
}
func createService(name string) {
	basePath := "app/services/"
	parts := strings.Split(name, "/")

	var folderPath, controllerName string

	if len(parts) > 1 {
		folderPath = basePath + strings.Join(parts[:len(parts)-1], "/")
		controllerName = parts[len(parts)-1]
	} else {
		folderPath = basePath
		controllerName = parts[0]
	}

	fileName := strings.ToLower(controllerName) + "_service.go"
	fullPath := folderPath + "/" + fileName

	// Buat folder jika belum ada
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.MkdirAll(folderPath, os.ModePerm)
	}

	// Cek apakah file sudah ada
	if _, err := os.Stat(fullPath); err == nil {
		fmt.Println("File Service sudah ada:", fullPath)
		return
	}

	// Template isi file controller
	content := fmt.Sprintf(`package services

import "fmt"

type %sService struct {}

func (c *%sService) Index() {
	fmt.Println("Controller %s berjalan...")
}
`, controllerName, controllerName, controllerName)

	// Simpan file
	err := os.WriteFile(fullPath, []byte(content), 0644)
	if err != nil {
		fmt.Println("Gagal membuat service:", err)
		return
	}

	fmt.Println("Service berhasil dibuat di:", fullPath)
}
func createRequest(name string) {
	basePath := "app/requests/"
	parts := strings.Split(name, "/")

	var folderPath, controllerName string

	if len(parts) > 1 {
		folderPath = basePath + strings.Join(parts[:len(parts)-1], "/")
		controllerName = parts[len(parts)-1]
	} else {
		folderPath = basePath
		controllerName = parts[0]
	}

	fileName := strings.ToLower(controllerName) + "_request.go"
	fullPath := folderPath + "/" + fileName

	// Buat folder jika belum ada
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.MkdirAll(folderPath, os.ModePerm)
	}

	// Cek apakah file sudah ada
	if _, err := os.Stat(fullPath); err == nil {
		fmt.Println("File Request sudah ada:", fullPath)
		return
	}

	// Template isi file controller
	// Template request (validasi)
	const requestTemplate = `package requests
	type {{.Name}}Request struct {
		Name  string ` + "`validate:\"required\"`" + `
		Email string ` + "`validate:\"required,email\"`" + `
	}`

	// Simpan file
	writeFileFromTemplate(fullPath, requestTemplate, name)
	fmt.Println("Request berhasil dibuat di:", fullPath)
}
func createResponse(name string) {
	basePath := "app/resources/"
	parts := strings.Split(name, "/")

	var folderPath, controllerName string

	if len(parts) > 1 {
		folderPath = basePath + strings.Join(parts[:len(parts)-1], "/")
		controllerName = parts[len(parts)-1]
	} else {
		folderPath = basePath
		controllerName = parts[0]
	}

	fileName := strings.ToLower(controllerName) + "_resource.go"
	fullPath := folderPath + "/" + fileName

	// Buat folder jika belum ada
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.MkdirAll(folderPath, os.ModePerm)
	}

	// Cek apakah file sudah ada
	if _, err := os.Stat(fullPath); err == nil {
		fmt.Println("File Resources sudah ada:", fullPath)
		return
	}

	// Template isi file controller
	// Template request (validasi)
	const requestTemplate = `package resources

	type {{.Name}}Resource struct {
		Name  string ` + "`validate:\"required\"`" + `
		Email string ` + "`validate:\"required,email\"`" + `
	}`

	// Simpan file
	writeFileFromTemplate(fullPath, requestTemplate, name)
	fmt.Println("Resource berhasil dibuat di:", fullPath)
}
func createModel(name string) {
	basePath := "app/models/"
	parts := strings.Split(name, "/")

	var folderPath, modelName string

	if len(parts) > 1 {
		folderPath = basePath + strings.Join(parts[:len(parts)-1], "/")
		modelName = parts[len(parts)-1]
	} else {
		folderPath = basePath
		modelName = parts[0]
	}

	fileName := strings.ToLower(modelName) + ".go"
	fullPath := folderPath + "/" + fileName

	// Buat folder jika belum ada
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.MkdirAll(folderPath, os.ModePerm)
	}

	// Cek apakah file sudah ada
	if _, err := os.Stat(fullPath); err == nil {
		fmt.Println("File model sudah ada:", fullPath)
		return
	}

	// Template isi file model
	content := fmt.Sprintf(`package %s
	import (
	"github.com/joho/godotenv"
	"os"
	"log"
	)
	type %s struct {
		ID   uint   `+"`gorm:\"primaryKey\" json:\"id\"`"+`
		Name string `+"`gorm:\"type:varchar(100)\" json:\"name\"`"+`
	}
	
	func (%s) TableName() string {
		errenv := godotenv.Load()
		if errenv != nil {
			log.Fatal(errenv)
		}
		DB_PREFIX := os.Getenv("DB_PREFIX")
		return DB_PREFIX+"_%s"
	}
	`, "models", modelName, modelName, strings.ToLower(modelName))

	// Simpan file
	err := os.WriteFile(fullPath, []byte(content), 0644)
	if err != nil {
		fmt.Println("Gagal membuat model:", err)
		return
	}
	err = os.Chmod(fullPath, 0644)
	if err != nil {
		fmt.Println("Gagal mengatur permission:", err)
	}
	rootPackage, err := getRootPackage()
	if err != nil {
		fmt.Println("Gagal mendapatkan root package:", err)
		return
	}
	fmt.Println("Model berhasil dibuat di:", fullPath)
	AddingMigration(rootPackage, modelName)
	AddingMethod(rootPackage, modelName)
}

func getRootPackage() (string, error) {
	file, err := os.Open("go.mod")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}

	return "", fmt.Errorf("module not found in go.mod")
}
func writeFileFromTemplate(filePath, tmpl string, name string) {
	t := template.Must(template.New("file").Parse(tmpl))
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Gagal membuat file:", err)
		return
	}
	defer f.Close()

	err = t.Execute(f, struct{ Name string }{Name: name})
	if err != nil {
		fmt.Println("Gagal menulis template:", err)
	}
}
func AddingMigration(rootPackage, modelName string) {
	migrationPath := "app/migrations/"
	timestamp := time.Now().Format("20060102150405") // Format: YYYYMMDDHHMMSS
	migrationFile := migrationPath + strings.ToLower(modelName) + "_" + timestamp + ".go"
	if _, err := os.Stat(migrationPath); os.IsNotExist(err) {
		os.MkdirAll(migrationPath, os.ModePerm)
	}
	if _, err := os.Stat(migrationFile); err == nil {
		fmt.Println("⚠️ File migration sudah ada:", migrationFile)
	} else {

		// Buat file migration
		CreateMigrationFile(rootPackage, modelName, migrationFile)
	}
}

func CreateMigrationFile(rootPackage, modelName, migrationFile string) {

	// Buat isi file migration secara dinamis
	migrationContent := fmt.Sprintf(`package migrations

	import (
		"gorm.io/gorm"
		"%s/app/models"
	)

	func Migrate%s(db *gorm.DB) error {
		return db.AutoMigrate(&models.%s{})
	}
`, rootPackage, modelName, modelName)

	err := os.WriteFile(migrationFile, []byte(migrationContent), 0644)
	if err != nil {
		fmt.Println("Gagal membuat migration:", err)
		return
	}

	fmt.Println("Migration berhasil dibuat di:", migrationFile)
}

func AddingMethod(rootPackage, modelName string) {
	mainFile := "app/migrate/migrate.go"

	// Baca isi file main.go
	content, err := os.ReadFile(mainFile)
	if err != nil {
		fmt.Println("Gagal membaca migrate.go:", err)
		return
	}
	mainCode := string(content)
	migrationImport := fmt.Sprintf(`"%s/app/migrations"`, rootPackage)

	if strings.Contains(mainCode, "import (") {
		if !strings.Contains(mainCode, migrationImport) {
			mainCode = strings.Replace(mainCode, "import (", "import (\n"+migrationImport, 1)
		}
	} else {
		lastImportIdx := strings.LastIndex(mainCode, "import ")
		if lastImportIdx != -1 {
			endImportIdx := strings.Index(mainCode[lastImportIdx:], "\n")
			if endImportIdx != -1 {
				insertPos := lastImportIdx + endImportIdx
				mainCode = mainCode[:insertPos] + "\n" + migrationImport + mainCode[insertPos:]
			}
		}
	}
	// **Tambahkan pemanggilan `generate<Model>()` di `main()`**
	generateCall := fmt.Sprintf("\tgenerate%s(db)", modelName)
	if !strings.Contains(mainCode, generateCall) {
		mainCode = strings.Replace(mainCode, "fmt.Println(\"Semua migrasi selesai!\")", generateCall+"\n\tfmt.Println(\"Semua migrasi selesai!\")", 1)
	}

	// **Tambahkan fungsi `generate<Model>()` jika belum ada**
	generateFunc := fmt.Sprintf(`
func generate%s(db *gorm.DB) {
		fmt.Println("Migrasi tabel %s...")
		if err := migrations.Migrate%s(db); err != nil {
			log.Fatalf("Gagal migrasi %s: %%v", err)
		}
		fmt.Println("Migrasi %s selesai.")
}`, modelName, modelName, modelName, modelName, modelName)
	if !strings.Contains(mainCode, fmt.Sprintf("func generate%s(", modelName)) {
		mainCode += generateFunc
	}
	// Simpan kembali main.go yang sudah diperbarui
	if err := os.WriteFile(mainFile, []byte(mainCode), 0644); err != nil {
		fmt.Println("Gagal memperbarui migrate.go:", err)
		return
	}
	fmt.Println("Fungsi generate" + modelName + " berhasil ditambahkan ke migrate.go")
}

func runSeeder(name string) {
	// Tentukan path file seeder yang ingin dijalankan
	seederPath := "app/seeders/" + strings.ToLower(name) + "_seeder.go"

	// Cek apakah file seeder ada
	if _, err := os.Stat(seederPath); err != nil {
		fmt.Println("Seeder tidak ditemukan:", seederPath)
		return
	}

	// Jalankan seeder dengan menjalankan perintah go run
	cmd := exec.Command("go", "run", seederPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Gagal menjalankan seeder:", err)
		return
	}

	fmt.Println("Seeder berhasil dijalankan:", name)
}
