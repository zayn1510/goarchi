package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/zayn1510/goarchi/config"
	"github.com/zayn1510/goarchi/core/tools"
	"github.com/zayn1510/goarchi/database/migrations"
	"os"
	"strings"
	"time"
)

// generate controller
var makeControllerCmd = &cobra.Command{
	Use:   "controller [name]",
	Short: "Generate a new controller",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := strings.ToLower(args[0])

		parts := strings.Split(path, "/")
		name := parts[len(parts)-1] // "user"

		structName := strings.Title(name) + "Controller"
		filePath := fmt.Sprintf("app/controllers/%s_controller.go", path)
		buf, err := tools.GenerateController(structName)
		if err != nil {
			fmt.Println("Failed to execute template:", err)
			return
		}
		content := buf

		// Buat folder-nya dulu kalau belum ada
		dir := "app/controllers/" + strings.Join(parts[:len(parts)-1], "/")
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Println("Failed to create folder", err)
			return
		}
		// Simpan file
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			fmt.Println("Failed to create controller", err)
			return
		}

		fmt.Printf("%s\n  → %s\n",
			color.HiGreenString("✅ Controller created successfully!"),
			color.YellowString(filePath),
		)
	},
}

var makeServiceCmd = &cobra.Command{
	Use:   "service [name]",
	Short: "Generate a new service",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := strings.ToLower(args[0])

		parts := strings.Split(path, "/")
		name := parts[len(parts)-1] // "user"

		structName := strings.Title(name) + "Service"
		filePath := fmt.Sprintf("app/services/%s_service.go", path)
		content, err := tools.GenerateServices(structName)
		if err != nil {
			fmt.Println("Failed to execute template:", err)
			return
		}
		// Buat folder-nya dulu kalau belum ada
		dir := "app/services/" + strings.Join(parts[:len(parts)-1], "/")
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Println("Gagal membuat folder:", err)
			return
		}
		// Simpan file
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			fmt.Println("Gagal membuat service:", err)
			return
		}

		fmt.Printf("%s\n  → %s\n",
			color.HiGreenString("✅ Service created successfully!"),
			color.YellowString(filePath),
		)
	},
}

// generate request
var makeRequestCmd = &cobra.Command{
	Use:   "request [name] [fields]",
	Short: "Generate a new request with optional fields",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := strings.ToLower(args[0]) // Misalnya: users/user
		// Ambil nama terakhir untuk nama struct
		parts := strings.Split(path, "/")
		name := parts[len(parts)-1] // "user"

		structName := strings.Title(name) + "Request"
		filePath := fmt.Sprintf("app/requests/%s_request.go", path)

		// Parsing fields
		var fieldsBuilder strings.Builder
		for _, fieldArg := range args[1:] {
			parts := strings.Split(fieldArg, ":")
			if len(parts) != 2 {
				fmt.Printf("Field '%s' tidak valid. Gunakan format name:type\n", fieldArg)
				return
			}
			fieldName := strings.Title(parts[0])
			fieldType := parts[1]
			fieldsBuilder.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\" validate:\"required\"`\n", fieldName, fieldType, parts[0]))
		}
		content, err := tools.GenerateRequest(structName, fieldsBuilder)
		if err != nil {
			fmt.Println("Failed to execute template:", err)
			return
		}
		dir := "app/requests/" + strings.Join(parts[:len(parts)-1], "/")
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Println("Gagal membuat folder:", err)
			return
		}

		// Simpan file
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			fmt.Println("Gagal membuat request:", err)
			return
		}
		fmt.Printf("%s\n  → %s\n",
			color.HiGreenString("✅ Request created successfully!"),
			color.YellowString(filePath),
		)
	},
}

// generate resource
var makeResourceCmd = &cobra.Command{
	Use:   "resource [name] [fields]",
	Short: "Generate a new resource with optional fields",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := strings.ToLower(args[0])
		parts := strings.Split(path, "/")
		name := parts[len(parts)-1]

		structName := strings.Title(name) + "Resource"
		filePath := fmt.Sprintf("app/resources/%s_resource.go", path)

		// Parsing fields
		var fieldsBuilder strings.Builder
		for _, fieldArg := range args[1:] {
			parts := strings.Split(fieldArg, ":")
			if len(parts) != 2 {
				fmt.Printf("Field '%s' tidak valid. Gunakan format name:type\n", fieldArg)
				return
			}
			fieldName := strings.Title(parts[0])
			fieldType := parts[1]
			fieldsBuilder.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\"`\n", fieldName, fieldType, parts[0]))
		}

		content, err := tools.GenerateResource(structName, fieldsBuilder)
		if err != nil {
			fmt.Println("Failed to execute template:", err)
			return
		}
		dir := "app/resources/" + strings.Join(parts[:len(parts)-1], "/")
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Println("Gagal membuat folder:", err)
			return
		}

		// Simpan file
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			fmt.Println("Gagal membuat resource:", err)
			return
		}

		fmt.Printf("%s\n  → %s\n",
			color.HiGreenString("✅ Request created successfully!"),
			color.YellowString(filePath),
		)
	},
}

var makeModelCmd = &cobra.Command{
	Use:   "model [path] [fields]",
	Short: "Generate a new model with fields and GORM tags",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := strings.ToLower(args[0])
		parts := strings.Split(path, "/")
		name := parts[len(parts)-1]
		structName := strings.Title(strings.TrimSuffix(name, "s"))

		dir := "app/models/" + strings.Join(parts[:len(parts)-1], "/")
		filePath := fmt.Sprintf("%s/%s.go", dir, name)
		var fieldsBuilder strings.Builder

		for _, fieldArg := range args[1:] {
			parts := strings.Split(fieldArg, ":")
			if len(parts) < 2 {
				fmt.Printf("Field '%s' tidak valid. Gunakan format name:type;tag1;tag2 atau Struct:foreignKey:Field\n", fieldArg)
				return
			}

			if strings.ToUpper(parts[0][:1]) == parts[0][:1] && parts[1] == "foreignKey" {
				structName := parts[0]
				foreignKey := parts[2]

				fieldsBuilder.WriteString(fmt.Sprintf("\t%s %s `gorm:\"foreignKey:%s\"`\n", structName, structName, foreignKey))
				continue
			}

			// Field biasa
			fieldName := strings.Title(parts[0])
			tagParts := strings.Split(parts[1], ";")
			fieldType := tagParts[0]
			gormTags := strings.Join(tagParts[1:], ";")

			jsonTag := strings.ToLower(parts[0])
			var tagBuilder strings.Builder

			if gormTags != "" {
				tagBuilder.WriteString(fmt.Sprintf("gorm:\"%s\" ", gormTags))
			}
			tagBuilder.WriteString(fmt.Sprintf("json:\"%s\"", jsonTag))

			fieldsBuilder.WriteString(fmt.Sprintf("\t%s %s `%s`\n", fieldName, fieldType, tagBuilder.String()))
		}
		fieldsBuilder.WriteString("\tCreatedAt time.Time `json:\"created_at\"`\n")
		fieldsBuilder.WriteString("\tUpdatedAt time.Time `json:\"updated_at\"`\n")
		content, err := tools.GenerateModel(structName, fieldsBuilder)
		if err != nil {
			fmt.Println("Gagal membuat model:", err)
			return
		}

		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Println("Gagal membuat folder:", err)
			return
		}

		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			fmt.Println("Gagal membuat model:", err)
			return
		}

		fmt.Printf("%s\n  → %s\n",
			color.HiGreenString("✅ Model created successfully!"),
			color.YellowString(filePath),
		)
	},
}
var makeMigrationCmd = &cobra.Command{
	Use:   "migration [name]",
	Short: "Generate a new migration file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		timestamp := time.Now().Format("20060102150405")
		fileName := fmt.Sprintf("database/migrations/%s_%s.go", timestamp, name)
		content, err := tools.GenerateMigration()
		if err != nil {
			fmt.Println("Failed to execute template:", err)
		}

		if err := os.MkdirAll("database/migrations", os.ModePerm); err != nil {
			fmt.Println("Gagal membuat folder:", err)
			return
		}

		if err := os.WriteFile(fileName, []byte(content), 0644); err != nil {
			fmt.Println("Gagal membuat migration:", err)
			return
		}

		fmt.Printf("%s\n  → %s\n",
			color.HiGreenString("✅ migration created successfully!"),
			color.YellowString(fileName),
		)
	},
}
var migrateCmd = &cobra.Command{
	Use:   "migrate [direction]",
	Short: "Run migrations (up or down)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		direction := args[0]

		// Connect DB (ubah sesuai config project kamu)
		db := config.GetDB()

		for _, migration := range migrations.AllMigrations {
			fmt.Println("Migrasi:", migration.Name)

			var migErr error
			if direction == "up" {
				migErr = migration.Up(db)
			} else if direction == "down" {
				migErr = migration.Down(db)
			} else {
				fmt.Println("Gunakan 'up' atau 'down'")
				return
			}

			if migErr != nil {
				fmt.Println("Gagal:", migErr)
				return
			}
			fmt.Println("Sukses:", migration.Name)
		}
	},
}

func init() {
	makeCmd.AddCommand(makeControllerCmd)
	makeCmd.AddCommand(makeServiceCmd)
	makeCmd.AddCommand(makeRequestCmd)
	makeCmd.AddCommand(makeResourceCmd)
	makeCmd.AddCommand(makeModelCmd)
	makeCmd.AddCommand(makeMigrationCmd)
	makeCmd.AddCommand(migrateCmd)
}
