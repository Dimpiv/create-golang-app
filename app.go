package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CreateDirectory(name string) {
	err := os.Mkdir(name, os.ModePerm)
	checkErr(err)
}

func CreateFile(name string) {
	file, err := os.Create(name)
	checkErr(err)
	err = file.Close()
	checkErr(err)
}

func CreateDirs(listDirs []string) {
	for _, dir := range listDirs {
		CreateDirectory(dir)
	}
}

func CreateFiles(listFiles []string) {
	for _, file := range listFiles {
		CreateFile(file)
	}
}

func CreateRootDir(name *string) {
	err := os.Mkdir(*name, os.ModePerm)
	checkErr(err)
	dir, err := os.Getwd()
	checkErr(err)
	err = os.Chdir(filepath.Join(dir, *name))
	checkErr(err)
}

func CreateGoMod() {
	cmd := exec.Command("go", "mod", "init", "example.com/m/v2")
	err := cmd.Run()
	checkErr(err)
}

func CreateInitGit() {
	cmd := exec.Command("git", "init")
	err := cmd.Run()
	checkErr(err)
}

func WriteGitignore() {
	f, err := os.OpenFile(".gitignore", 1, os.ModePerm)
	checkErr(err)
	if _, err := f.Write([]byte("build\n")); err != nil {
		_ = f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func WriteReadme(name *string) {
	f, err := os.OpenFile("README.md", 1, os.ModePerm)
	checkErr(err)
	if _, err := f.Write([]byte(fmt.Sprintf("#%s\n", *name))); err != nil {
		_ = f.Close()
		log.Fatal(err)
	}
	if _, err := f.Write([]byte("`Hello World! This is New golang project!`\n")); err != nil {
		_ = f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	nameProject := flag.String("n", "NewGoProject", "Name of new golang project")
	flag.Parse()

	dirs := []string{"cmd", "internal", "api", "configs", "build"}
	files := []string{"README.md", ".gitignore"}

	CreateRootDir(nameProject)
	CreateDirs(dirs)
	CreateFiles(files)
	CreateGoMod()
	CreateInitGit()
	WriteGitignore()
	WriteReadme(nameProject)
}
