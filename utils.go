package main

import "strings"

func isEbook(path string) bool {
	if strings.HasSuffix(path, ".pdf") || strings.HasSuffix(path, ".epub") {
		return true
	}

	return false
}
