package filescmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileTracker_HasReturnsFalseWhenEmpty(t *testing.T) {
	ft := NewFileTracker()

	require.False(t, ft.HasPath("/foo/bar"))
	require.False(t, ft.HasChecksum("7d97e98f8af710c7e7fe703abc8f639e0ee507c4"))

	path, ok := ft.GetPathByChecksum("7d97e98f8af710c7e7fe703abc8f639e0ee507c4")
	require.False(t, ok)
	require.Empty(t, path)

	paths, ok := ft.GetPathsByChecksum("7d97e98f8af710c7e7fe703abc8f639e0ee507c4")
	require.False(t, ok)
	require.Nil(t, paths)
}

func TestFileTracker_TracksSingleEntry(t *testing.T) {
	ft := NewFileTracker()
	expectedPath := "/foo/bar"
	expectedChecksum := "7d97e98f8af710c7e7fe703abc8f639e0ee507c4"

	ft.Set(expectedPath, expectedChecksum)

	require.True(t, ft.HasPath(expectedPath))
	require.True(t, ft.HasChecksum(expectedChecksum))

	path, ok := ft.GetPathByChecksum(expectedChecksum)
	require.True(t, ok)
	require.Equal(t, expectedPath, path)

	paths, ok := ft.GetPathsByChecksum(expectedChecksum)
	require.True(t, ok)
	require.Equal(t, []string{expectedPath}, paths)
}

func TestFileTracker_TracksAllPathsForChecksum(t *testing.T) {
	ft := NewFileTracker()
	expectedChecksum := "7d97e98f8af710c7e7fe703abc8f639e0ee507c4"

	ft.Set("/foo/bar", expectedChecksum)
	ft.Set("/foo/baz", expectedChecksum)

	paths, ok := ft.GetPathsByChecksum(expectedChecksum)
	require.True(t, ok)
	require.Equal(t, []string{"/foo/bar", "/foo/baz"}, paths)

	path, ok := ft.GetPathByChecksum(expectedChecksum)
	require.True(t, ok)
	require.Equal(t, "/foo/bar", path)
}

func TestFileTracker_DuplicatePathsByChecksum(t *testing.T) {
	ft := NewFileTracker()
	duplicateChecksum := "7d97e98f8af710c7e7fe703abc8f639e0ee507c4"
	uniqueChecksum := "3f786850e387550fdab836ed7e6dc881de23001b"

	ft.Set("/qux/bar", duplicateChecksum)
	ft.Set("/foo/baz", uniqueChecksum)
	ft.Set("/foo/bar", duplicateChecksum)

	duplicates := ft.DuplicatePathsByChecksum()

	require.Equal(t, map[string][]string{
		duplicateChecksum: {"/foo/bar", "/qux/bar"},
	}, duplicates)
}

func TestNewFileTracker(t *testing.T) {
	require.NotNil(t, NewFileTracker())
}
