package main

import (
	"filereader"
	"fmt"
	"strconv"
	"strings"
	// "strings"
	// "regexp"
	// "strings"
)

func main() {
	mainDirectory := Directory{}
	mainDirectory.Name = "main"
	mainDirectory.Directories = make([]Directory, 0)
	mainDirectory.Files = make([]File, 0)

	input := filereader.ReadFile("input.txt")
	currentDirectory := &mainDirectory

	for i := 0; i < len(input); i++ {
		trimmedString := strings.Replace(input[i], "$", "", -1)

		if string(trimmedString[0]) == " " {
			trimmedString = trimmedString[1:]
		}

		if trimmedString == "cd /" {
			currentDirectory = &mainDirectory
		} else if trimmedString == "ls" {

		} else if trimmedString[0:3] == "dir" {
			directoryName := strings.Split(input[i], " ")[1]
			directory := currentDirectory.Append(directoryName)
			currentDirectory.Directories = append(currentDirectory.Directories, *directory)
		} else if trimmedString[0:2] == "cd" && trimmedString[3:] != ".." {
			directoryResult := GetDirectoryOrNil(currentDirectory.Directories, trimmedString[3:])
			currentDirectory = directoryResult
		} else if trimmedString[0:2] == "cd" && trimmedString[3:] == ".." {
			currentDirectory = currentDirectory.ParentDirectory
		} else {
			splittedBySpace := strings.Split(input[i], " ")
			size, _ := strconv.Atoi(splittedBySpace[0])

			file := new(File)
			file.size = size
			file.name = splittedBySpace[1]
			currentDirectory.Files = append(currentDirectory.Files, *file)
		}
	}

	fmt.Print("gg")
}

func GetDirectoryOrNil(directories []Directory, name string) *Directory {
	for _, v := range directories {
		if v.Name == name {
			return &v
		}
	}

	return nil
}

type Directory struct {
	Name            string
	Directories     []Directory
	Files           []File
	ParentDirectory *Directory
}

func (dir Directory) IsMain() bool {
	return dir.ParentDirectory == nil
}

func (parent *Directory) Append(name string) *Directory {
	result := new(Directory)
	result.Name = name
	result.ParentDirectory = parent

	return result
}

type File struct {
	name string
	size int
}

type asd struct {
	addd File
}
