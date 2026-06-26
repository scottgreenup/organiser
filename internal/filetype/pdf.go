package filetype

import (
	"path/filepath"
	"strings"
)

func IsPDF(path string) bool {
	knownSuffixes := []string{
		".pdf",
	}

	ext := strings.ToLower(filepath.Ext(path))
	for _, suffix := range knownSuffixes {
		if suffix == ext {
			return true
		}
	}

	return false
}
