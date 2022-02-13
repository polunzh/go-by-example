package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("a", "b", "/", "/", "c")
	fmt.Println("p:", p)

	fmt.Println(filepath.Join("dir1//", "c"))
	fmt.Println(filepath.Join("dir1/../dir2", "c"))

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	fmt.Println("IsAbs(dir/file)", filepath.IsAbs("dir/file"))
	fmt.Println("IsAbs(/dir/file)", filepath.IsAbs("/dir/file"))

	filename := "config.json"

	ext := filepath.Ext(filename)
	fmt.Println("ext:", ext)
	fmt.Println("filename without ext:", strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")

	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
