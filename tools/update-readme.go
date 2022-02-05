package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const SITE_URL = "https://gobyexample.com/"

const tmpl = `
# Test code of "Go by Example"

[Go by Example](https://gobyexample.com/hello-world)

## Table of Contents

| Order  | Codes  | Details  |
|---|---|---|
{{- range $index, $val := .}}
| {{$index}}  | [{{$val.Title}}](./{{$val.FileName}}) | {{$val.Link}}  |
{{- end}}
`

type Code struct {
	FileName string
	Title string
	Link string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	var codes []Code
	for _, f := range files {
		filename := f.Name()
		fi, err := os.Stat(filename)
		if err != nil {
			log.Fatal(err)
		}

		if fi.Mode().IsRegular() && filepath.Ext(f.Name()) == ".go" {
			fileBaseName := strings.TrimSuffix(filename, filepath.Ext(filename))
			title := strings.ReplaceAll(fileBaseName, "-", " ")
			title = strings.Title(title)

			codes = append(codes, Code{FileName: filename, Title: title, Link: fmt.Sprintf("%s%s", SITE_URL, fileBaseName)})
		}
	}

	t := template.Must(template.New("tmpl").Parse(tmpl))
	f, err := os.Create("README.md")
	check(err)

	t.Execute(f, codes)
}
