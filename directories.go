package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func listDir(c []fs.DirEntry) {
	for _, entry := range c {
		prefix := ""
		if entry.IsDir() {
			prefix = "d"
		} else {
			prefix = "-"
		}

		fmt.Println(prefix, entry.Name())
	}
}

func main() {
	dir := "/tmp/subdir"
	err := os.Mkdir(dir, 0755)
	check(err)

	defer os.RemoveAll("/tmp/subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("/tmp/subdir/file1")
	err = os.MkdirAll("/tmp/subdir/parent/child", 0755)
	check(err)

	createEmptyFile("/tmp/subdir/parent/file2")
	createEmptyFile("/tmp/subdir/parent/file3")
	createEmptyFile("/tmp/subdir/parent/child/file4")

	c, err := os.ReadDir("/tmp/subdir/parent")
	check(err)
	listDir(c)
	fmt.Println("Listing subdir/parent")

	err = os.Chdir("/tmp/subdir/parent/child")
	pwd, err:= os.Getwd()
	check(err)
	fmt.Println("Current directory:", pwd)
	c, err = os.ReadDir(".")
	listDir(c)

	err = os.Chdir("../../..")
	check(err)

	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

func visit(p string, info os.FileInfo, err error) error{
	if err != nil {
		return err
	}

	fmt.Println(" ", p, info.IsDir())
	return nil
}
