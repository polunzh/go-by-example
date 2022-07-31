package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Example struct {
	Id   string
	Name string
}

type Code struct {
	Index     string
	FileName  string
	Title     string
	Practiced bool
	Link      string
}

type Template struct {
	Codes     []Code
	Completed int
	Total     int
	Progress  string
}

const SITE_URL = "https://gobyexample.com/"

const tmpl = `
# Test code of "Go by Example"

[Go by Example](https://gobyexample.com/hello-world)

## Table of Contents

Progress __{{.Completed}}/{{.Total}}__

| Order  | Codes  | Details  |
|---|---|---|
{{- range $index, $val := .Codes}}
| {{.Index}}  | {{if .Practiced }} [{{$val.Title}}](./{{$val.FileName}}) {{else}} {{$val.Title}} {{end}} | {{$val.Link}}  |
{{- end}}
`

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

	exampleBytes, err := os.ReadFile("./tools/all-examples.json")
	if err != nil {
		log.Fatal(err)
	}

	var examples []Example
	json.Unmarshal(exampleBytes, &examples)

	var codes []Code
	filesMap := make(map[string]bool)

	for _, f := range files {
		filename := f.Name()
		fi, err := os.Stat(filename)
		if err != nil {
			log.Fatal(err)
		}

		if fi.Mode().IsRegular() && filepath.Ext(f.Name()) == ".go" {
			fileBaseName := strings.TrimSuffix(filename, filepath.Ext(filename))
			filesMap[fileBaseName] = true
		}
	}

	for index, example := range examples {
		codes = append(codes, Code{Index: fmt.Sprintf("%.2d", index+1), FileName: fmt.Sprintf("%s.go", example.Id), Title: example.Name, Practiced: filesMap[example.Id], Link: fmt.Sprintf("%s%s", SITE_URL, example.Id)})
	}

	t := template.Must(template.New("tmpl").Parse(tmpl))
	f, err := os.Create("README.md")
	check(err)

	t.Execute(f, Template{Codes: codes, Completed: len(filesMap), Total: len(codes), Progress: fmt.Sprintf("%.2f%%", 100*float64(len(filesMap))/float64(len(codes)))})
	fmt.Println("Updated README!")
}
