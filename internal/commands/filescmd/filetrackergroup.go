package filescmd

type FileTrackerGroup struct {
	byChecksum map[string][]string
	byPath     map[string]string
}

func NewFileTrackerGroup() *FileTrackerGroup {
	return &FileTrackerGroup{
		byChecksum: make(map[string][]string),
		byPath:     make(map[string]string),
	}
}

func (t *FileTrackerGroup) Set(path string, checksum string) {
	t.byChecksum[checksum] = append(t.byChecksum[checksum], path)
	t.byPath[path] = checksum
}

func (t *FileTrackerGroup) HasPath(path string) bool {
	if _, ok := t.byPath[path]; ok {
		return true
	}
	return false
}

func (t *FileTrackerGroup) HasChecksum(checksum string) bool {
	if _, ok := t.byChecksum[checksum]; ok {
		return true
	}
	return false
}

func (t *FileTrackerGroup) GetPathByChecksum(checksum string) (path string, ok bool) {
	paths, ok := t.byChecksum[checksum]
	if !ok || len(paths) == 0 {
		return "", false
	}
	return paths[0], true
}

func (t *FileTrackerGroup) GetPathsByChecksum(checksum string) (paths []string, ok bool) {
	paths, ok = t.byChecksum[checksum]
	return paths, ok
}
