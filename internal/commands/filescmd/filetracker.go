package filescmd

import "sort"

type FileTracker struct {
	groups map[string]*FileTrackerGroup
}

func NewFileTracker() *FileTracker {
	return &FileTracker{
		groups: make(map[string]*FileTrackerGroup),
	}
}

func (d *FileTracker) HasPath(path string, group string) bool {
	if _, ok := d.groups[group]; !ok {
		return false
	}
	return d.groups[group].HasPath(path)
}

func (d *FileTracker) HasPathAnywhere(path string) bool {
	for _, g := range d.groups {
		if g.HasPath(path) {
			return true
		}
	}
	return false
}

func (d *FileTracker) HasChecksum(checksum string, group string) bool {
	if _, ok := d.groups[group]; !ok {
		return false
	}
	return d.groups[group].HasChecksum(checksum)
}

func (d *FileTracker) HasChecksumAnywhere(checksum string) bool {
	for _, g := range d.groups {
		if g.HasChecksum(checksum) {
			return true
		}
	}
	return false
}

func (d *FileTracker) HasChecksumAnywhereExcept(checksum string, groupToSkip string) bool {
	for groupName, g := range d.groups {
		if groupName == groupToSkip {
			continue
		}
		if g.HasChecksum(checksum) {
			return true
		}
	}
	return false
}

func (d *FileTracker) GetPathByChecksum(checksum string, group string) (path string, ok bool) {
	if _, ok := d.groups[group]; !ok {
		return "", false
	}
	return d.groups[group].GetPathByChecksum(checksum)
}

func (d *FileTracker) GetPathByChecksumAnywhere(checksum string) (path string, ok bool) {
	for _, g := range d.groups {
		if path, ok := g.GetPathByChecksum(checksum); ok {
			return path, true
		}
	}
	return "", false
}

func (d *FileTracker) DuplicatePathsByChecksum() map[string][]string {
	pathsByChecksum := make(map[string][]string)

	for _, group := range d.groups {
		for checksum, paths := range group.byChecksum {
			pathsByChecksum[checksum] = append(pathsByChecksum[checksum], paths...)
		}
	}

	duplicates := make(map[string][]string)
	for checksum, paths := range pathsByChecksum {
		if len(paths) > 1 {
			sort.Strings(paths)
			duplicates[checksum] = paths
		}
	}

	return duplicates
}

func (d *FileTracker) Set(path string, checksum string, group string) {
	if _, ok := d.groups[group]; !ok {
		d.groups[group] = NewFileTrackerGroup()
	}

	d.groups[group].Set(path, checksum)
}
