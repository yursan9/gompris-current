package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

func SetupTemplate() *template.Template {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	templateFile := filepath.Join(home, ".config", "mpris-current", "template")

	var templateString string
	content, err := ioutil.ReadFile(templateFile)
	if err != nil {
		templateString = "{{.Status}}: {{.Title}} - {{.Artist}}"
	} else {
		templateString = string(content)
	}

	return template.Must(template.New("player-status").Parse(templateString))
}

func CreateFile() *os.File {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}

	path := filepath.Join(cacheDir, "mpris-current")
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	return f
}
