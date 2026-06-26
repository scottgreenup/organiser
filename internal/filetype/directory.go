package filetype

import (
	"os"

	"github.com/pkg/errors"
)

// IsDir sort of checks if the path is a directory.
// TODO: All callers need to handle the error cases properly
func IsDir(path string) (isDir bool, err error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, errors.WithMessagef(err, "could not stat %s", path)
	}
	return fi.IsDir(), nil
}
