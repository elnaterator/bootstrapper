package core

import (
	"bufio"
	"log"
	"os"
	"path"
	"strings"
)

type Directory struct {
	Name        string
	Files       []File
	Directories []Directory
}

type File struct {
	Name    string
	Content string
}

func BuildFile(name string, content string) File {
	return File{
		Name:    name,
		Content: content,
	}
}

func WriteFile(filePath string, content string) error {
	parentDirPath := path.Dir(filePath)
	err := os.MkdirAll(parentDirPath, 0755)
	if err != nil {
		return err
	}
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	w := bufio.NewWriter(f)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()
	_, err = w.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

func Traverse(
	dir *Directory,
	fileHandler func(*File, []string) error,
	dirHandler func(*Directory, []string) error,
) error {
	return traverseRecursive(dir, []string{}, fileHandler, dirHandler)
}

func traverseRecursive(
	dir *Directory,
	parents []string,
	fileHandler func(*File, []string) error,
	dirHandler func(*Directory, []string) error,
) error {

	if dir == nil {
		log.Fatal("invalid directory structure")
	}

	if dirHandler != nil {
		err := dirHandler(dir, parents)
		if err != nil {
			return err
		}
	}

	parents = append(parents, dir.Name)

	for _, f := range dir.Files {
		if fileHandler != nil {
			err := fileHandler(&f, parents)
			if err != nil {
				return err
			}
		}
	}

	for _, d := range dir.Directories {
		err := traverseRecursive(&d, parents, fileHandler, dirHandler)
		if err != nil {
			return err
		}
	}

	return nil

}

func WriteFileHandler(file *File, parents []string) error {
	filePath := strings.Join(append(parents, file.Name), string(os.PathSeparator))
	err := WriteFile(filePath, file.Content)
	if err != nil {
		return err
	}
	return nil
}
