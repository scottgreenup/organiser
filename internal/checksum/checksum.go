package checksum

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"

	"github.com/scottgreenup/organiser/internal/filetype"
)

//type Checksum interface {
//	Raw() []byte
//	Digest() string
//}
//
//type FileChecksum struct {
//}
//
//func File_new(path string) (Checksum, error) {
//
//	isDir, err := filetype.IsDir(path)
//	if err != nil {
//		return nil, errors.WithMessagef(err, "unable to checksum %q", path)
//	}
//	if isDir {
//		return nil, errors.Errorf("unable to checksum %q as it is a directory", path)
//	}
//
//	return &FileChecksum{}
//}
//
//func (c *FileChecksum) Raw() []byte {
//
//}
//
//func (c *FileChecksum) Digest() string {
//
//}

// FileDigest calculates the MD5 checksum of the file at the given path as a hexadecimal string.
func FileDigest(path string) (string, error) {
	s, err := File(path)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", s), nil
}

// File calculates the MD5 checksum of the file at the given path as raw bytes.
func File(path string) ([]byte, error) {
	isDir, err := filetype.IsDir(path)
	if err != nil {
		return nil, errors.WithMessagef(err, "unable to checksum %q", path)
	}
	if isDir {
		return nil, errors.Errorf("unable to checksum %q as it is a directory", path)
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fileHash := md5.New()
	buffer := make([]byte, 1000*1000)
	isEOF := false
	for !isEOF {
		readCount, err := f.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				isEOF = true
			} else {
				return nil, errors.WithMessagef(err, "could no longer read %q", path)
			}
		}
		if readCount == 0 {
			break
		}

		hashedCount, err := fileHash.Write(buffer[0:readCount])
		if err != nil {
			return nil, errors.WithMessagef(err, "could not hash output from %q", path)
		}

		if hashedCount != readCount {
			return nil, errors.Errorf("Read %d bytes, but hashed %d bytes from %q", readCount, hashedCount, path)
		}
	}

	return fileHash.Sum(nil), nil
}
