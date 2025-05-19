package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
	"time"

	"gopkg.in/yaml.v2"
)

type IndexYaml struct {
	Entries map[string][]ChartEntry `yaml:"entries"`
}

type ChartEntry struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	AppVersion  string `yaml:"appVersion"`
	Description string `yaml:"description"`
}

var readmeTemplate = `# Gwynbliedd Helm Charts Repository

## Overview

This repository contains Helm charts for various applications and services. These charts simplify the deployment of applications on Kubernetes clusters.

## Usage

### Adding the Repository

To add this Helm repository to your local Helm configuration:

` + "`" + `bash
helm repo add gwynbliedd https://gutsbutcher.github.io/helm-charts/
helm repo update
` + "`" + `

### Available Charts

| Chart Name | Version | App Version | Description | Installation Command |
|-----------|---------|-------------|-------------|----------------------|
{{- range .Entries }}
{{- range . }}
| {{ .Name }} | {{ .Version }} | {{ if .AppVersion }}{{ .AppVersion }}{{ else }}N/A{{ end }} | {{ if .Description }}{{ .Description }}{{ else }}N/A{{ end }} | ` + "`" + `helm install {{ .Name }} gwynbliedd/{{ .Name }}` + "`" + ` |
{{- end }}
{{- end }}

### Installation

To install a specific chart:

` + "`" + `bash
helm install <release-name> gwynbliedd/<chart-name>
` + "`" + `

## Last Updated

*Automatically generated on {{ .GeneratedTime }}*

## License
see the [LICENSE](LICENSE) file for details.

`

func main() {
	// Read index.yaml
	data, err := ioutil.ReadFile("index.yaml")
	if err != nil {
		fmt.Println("Error reading index.yaml:", err)
		os.Exit(1)
	}

	// Parse YAML
	var indexYaml IndexYaml
	err = yaml.Unmarshal(data, &indexYaml)
	if err != nil {
		fmt.Println("Error parsing YAML:", err)
		os.Exit(1)
	}

	// Prepare template data
	templateData := struct {
		Entries        map[string][]ChartEntry
		GeneratedTime  string
	}{
		Entries:       indexYaml.Entries,
		GeneratedTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	// Create template with custom functions
	funcMap := template.FuncMap{
		"default": func(defaultVal, val string) string {
			if val == "" {
				return defaultVal
			}
			return val
		},
	}

	// Create template
	tmpl, err := template.New("readme").Funcs(funcMap).Parse(readmeTemplate)
	if err != nil {
		fmt.Println("Error creating template:", err)
		os.Exit(1)
	}

	// Create README.md
	file, err := os.Create("README.md")
	if err != nil {
		fmt.Println("Error creating README.md:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Execute template
	err = tmpl.Execute(file, templateData)
	if err != nil {
		fmt.Println("Error executing template:", err)
		os.Exit(1)
	}

	fmt.Println("README.md generated successfully")
}
