package tools

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

func GenerateController(structName string) (string, error) {
	templatePath := "core/generate/templates/controller.tmpl"

	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", fmt.Errorf("template parsing error: %s", err)
	}

	var buf bytes.Buffer
	data := struct {
		StructName string
	}{
		StructName: structName,
	}
	err = tpl.Execute(&buf, data)
	if err != nil {
		return "", fmt.Errorf("template execution error: %s", err)
	}
	return buf.String(), nil
}

func GenerateServices(structName string) (string, error) {
	templatePath := "core/generate/templates/services.tmpl"
	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", fmt.Errorf("template parsing error: %s", err)
	}

	var buf bytes.Buffer
	data := struct {
		StructName string
	}{
		StructName: structName,
	}
	err = tpl.Execute(&buf, data)
	if err != nil {
		return "", fmt.Errorf("template execution error: %s", err)
	}
	return buf.String(), nil
}

func GenerateModel(structName string, fields strings.Builder) (string, error) {
	templatePath := "core/generate/templates/model.tmpl"
	data := struct {
		StructName string
		Fields     string
	}{
		StructName: structName,
		Fields:     fields.String(),
	}

	// Load template
	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func GenerateRequest(structName string, fields strings.Builder) (string, error) {
	templatePath := "core/generate/templates/request.tmpl"
	data := struct {
		StructName string
		Fields     string
	}{
		StructName: structName,
		Fields:     fields.String(),
	}

	// Load template
	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
func GenerateResource(structName string, fields strings.Builder) (string, error) {
	templatePath := "core/generate/templates/resource.tmpl"
	data := struct {
		StructName string
		Fields     string
	}{
		StructName: structName,
		Fields:     fields.String(),
	}

	// Load template
	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func GenerateMigration() (string, error) {
	templatePath := "core/generate/templates/migration.tmpl"
	// Load template
	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, nil)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
