package core

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func GetJson(url string) interface{} {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var result interface{}
	err = json.Unmarshal(content, &result)
	if err != nil {
		panic(err)
	}
	return result
}

func WriteFile(filePath string, content string) error {
	parentDirPath := path.Dir(filePath)
	err := os.MkdirAll(parentDirPath, os.ModeDir)
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
		dirHandler(dir, parents)
	}

	parents = append(parents, dir.Name)

	for _, f := range dir.Files {
		if fileHandler != nil {
			fileHandler(&f, parents)
		}
	}

	for _, d := range dir.Directories {
		traverseRecursive(&d, parents, fileHandler, dirHandler)
	}

	return nil

}
