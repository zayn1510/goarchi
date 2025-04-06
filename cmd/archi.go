package cmd

import (
	"fmt"
	"github.com/zayn1510/goarchi/config"
	"github.com/zayn1510/goarchi/database/migrations"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
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

		content := fmt.Sprintf(`package controllers

import "github.com/gin-gonic/gin"

type %s struct{}

func (ctrl *%s) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from %s",
	})
}
`, structName, structName, structName)

		// Buat folder-nya dulu kalau belum ada
		dir := "app/controllers/" + strings.Join(parts[:len(parts)-1], "/")
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Println("Gagal membuat folder:", err)
			return
		}
		// Simpan file
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			fmt.Println("Gagal membuat controller:", err)
			return
		}

		fmt.Println("Controller berhasil dibuat:", filePath)
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

		content := fmt.Sprintf(`package services

type %s struct{}
`, structName)

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
		fmt.Println("Service berhasil dibuat:", filePath)
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

		content := fmt.Sprintf(`package requests

import (
	"github.com/gin-gonic/gin"
)

type %s struct {
%s}

func Bind%s(c *gin.Context) (*%s, error) {
	var req %s
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
`, structName, fieldsBuilder.String(), structName, structName, structName)

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

		fmt.Println("Request berhasil dibuat:", filePath)
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

		content := fmt.Sprintf(`package resources


type %s struct {
%s}

`, structName, fieldsBuilder.String())

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

		fmt.Println("Resource berhasil dibuat:", filePath)
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
				// Ini struct relasi
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
		content := fmt.Sprintf(`package models
import "time"
type %s struct {
%s}
`, structName, fieldsBuilder.String())

		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Println("Gagal membuat folder:", err)
			return
		}

		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			fmt.Println("Gagal membuat model:", err)
			return
		}

		fmt.Println("Model berhasil dibuat:", filePath)
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

		content := `package migrations

import (
	"gorm.io/gorm"
)

func Up(db *gorm.DB) error {
	// TODO: implement migration
	return nil
}

func Down(db *gorm.DB) error {
	// TODO: implement rollback
	return nil
}
`

		if err := os.MkdirAll("database/migrations", os.ModePerm); err != nil {
			fmt.Println("Gagal membuat folder:", err)
			return
		}

		if err := os.WriteFile(fileName, []byte(content), 0644); err != nil {
			fmt.Println("Gagal membuat migration:", err)
			return
		}

		fmt.Println("Migration berhasil dibuat:", fileName)
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
