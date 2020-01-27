package main

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

var bookExt = []string{".pdf", ".epub"}

// TODO: Convert file into mobi
func convertFile(path, outputDir string) error {
	log.Printf("Converting: %s in %s", path, outputDir)

	var fileName string

	for _, ext := range bookExt {
		if strings.HasSuffix(path, ext) {
			fileName = strings.TrimSuffix(path, ext)
			break
		}
	}

	cFileName := fmt.Sprintf("%s.mobi", filepath.Base(fileName))
	cPath := filepath.Join(outputDir, cFileName)

	// Execute ebook-convertor
	cmd := exec.Command("ebook-convert", path, cPath)

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
