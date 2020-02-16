package main

import "strings"

var (
	// EbookExtensions are the valid extensions which can be handled by Curator
	EbookExtensions = []string{
		".pdf",
		".epub",
	}
)

func isEbook(path string) bool {
	for _, ext := range EbookExtensions {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}

	return false
}
