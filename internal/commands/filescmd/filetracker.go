package filescmd

import "sort"

type FileTracker struct {
	byChecksum map[string][]string
	byPath     map[string]string
}

func NewFileTracker() *FileTracker {
	return &FileTracker{
		byChecksum: make(map[string][]string),
		byPath:     make(map[string]string),
	}
}

func (d *FileTracker) HasPath(path string) bool {
	_, ok := d.byPath[path]
	return ok
}

func (d *FileTracker) HasChecksum(checksum string) bool {
	_, ok := d.byChecksum[checksum]
	return ok
}

func (d *FileTracker) GetPathByChecksum(checksum string) (path string, ok bool) {
	paths, ok := d.byChecksum[checksum]
	if !ok || len(paths) == 0 {
		return "", false
	}
	return paths[0], true
}

func (d *FileTracker) GetPathsByChecksum(checksum string) (paths []string, ok bool) {
	paths, ok = d.byChecksum[checksum]
	return paths, ok
}

func (d *FileTracker) DuplicatePathsByChecksum() map[string][]string {
	duplicates := make(map[string][]string)
	for checksum, paths := range d.byChecksum {
		if len(paths) > 1 {
			sort.Strings(paths)
			duplicates[checksum] = paths
		}
	}

	return duplicates
}

func (d *FileTracker) Set(path string, checksum string) {
	d.byChecksum[checksum] = append(d.byChecksum[checksum], path)
	d.byPath[path] = checksum
}
