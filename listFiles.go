package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func getPath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return path.Dir(filename)
}

func getGoFiles() ([]string, error) {
	here := getPath()
	pattern := "*.go"
	var matches []string
	err := filepath.Walk(here, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func listGoFiles() {
	files, err := getGoFiles()
	if err != nil {
		fmt.Print(err)
	}
	for _, s := range files {
		s = strings.TrimPrefix(s, "\\")
		fmt.Println(strings.TrimSuffix(s, ".go"))
	}
}
