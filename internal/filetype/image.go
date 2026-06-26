package filetype

import (
	"path/filepath"
	"strings"
)

func IsImage(path string) bool {
	knownSuffixes := []string{
		".png", ".jpeg", ".jpg", ".gif",
	}

	ext := strings.ToLower(filepath.Ext(path))
	for _, suffix := range knownSuffixes {
		if suffix == ext {
			return true
		}
	}

	return false
}
