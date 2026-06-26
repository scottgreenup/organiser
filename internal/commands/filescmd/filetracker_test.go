package filescmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileTrackerGroup_CanSetOnceAndCheck(t *testing.T) {
	// Given a FileTrackerGroup with an entry
	ftg := NewFileTrackerGroup()
	expectedPath := "/foo/bar"
	expectedChecksum := "7d97e98f8af710c7e7fe703abc8f639e0ee507c4"
	ftg.Set("/foo/bar", expectedChecksum)

	// When we check for the path/checksum
	hasChecksumResult := ftg.HasChecksum(expectedChecksum) // has value
	hasPathResult := ftg.HasPath(expectedPath)             // has key

	// Then we expect them to exist
	require.True(t, hasChecksumResult)
	require.True(t, hasPathResult)
}

func TestFileTrackerGroup_TracksAllPathsForChecksum(t *testing.T) {
	ftg := NewFileTrackerGroup()
	expectedChecksum := "7d97e98f8af710c7e7fe703abc8f639e0ee507c4"

	ftg.Set("/foo/bar", expectedChecksum)
	ftg.Set("/foo/baz", expectedChecksum)

	paths, ok := ftg.GetPathsByChecksum(expectedChecksum)
	require.True(t, ok)
	require.Equal(t, []string{"/foo/bar", "/foo/baz"}, paths)

	path, ok := ftg.GetPathByChecksum(expectedChecksum)
	require.True(t, ok)
	require.Equal(t, "/foo/bar", path)
}

func TestFileTrackerGroup_HasReturnsFalse(t *testing.T) {
	// Given a FileTrackerGroup with no entries
	ftg := NewFileTrackerGroup()
	expectedPath := "/foo/bar"
	expectedChecksum := "7d97e98f8af710c7e7fe703abc8f639e0ee507c4"

	// When we check for the path/checksum
	hasChecksumResult := ftg.HasChecksum(expectedChecksum)
	hasPathResult := ftg.HasPath(expectedPath)

	// Then we expect them to exist
	require.False(t, hasChecksumResult)
	require.False(t, hasPathResult)
}

func TestFileTracker_HasPath(t *testing.T) {
	// Given a FileTracker with no entries
	ft := NewFileTracker()
	expectedPath := "/foo/bar"
	expectedChecksum := "7d97e98f8af710c7e7fe703abc8f639e0ee507c4"
	expectedGroup := "default"

	// Then we can get it back out
	require.False(t, ft.HasPath(expectedPath, expectedGroup))

	// When we set an entry
	ft.Set(expectedPath, expectedChecksum, expectedGroup)

	// Then we can get it back out
	require.True(t, ft.HasPath(expectedPath, expectedGroup))
}

func TestFileTracker_SingleEntryComprehensive(t *testing.T) {
	// Given a FileTracker with no entries
	ft := NewFileTracker()
	expectedPath := "/foo/bar"
	expectedChecksum := "7d97e98f8af710c7e7fe703abc8f639e0ee507c4"
	expectedGroup := "default"

	// Then we shouldn't be able to get that entry yet.
	require.False(t, ft.HasPath(expectedPath, expectedGroup))
	require.False(t, ft.HasPathAnywhere(expectedPath))
	require.False(t, ft.HasChecksum(expectedChecksum, expectedGroup))
	require.False(t, ft.HasChecksumAnywhere(expectedChecksum))
	require.False(t, ft.HasChecksumAnywhereExcept(expectedChecksum, expectedGroup))
	require.False(t, ft.HasChecksumAnywhereExcept(expectedChecksum, expectedGroup+"a"))

	// When we set an entry
	ft.Set(expectedPath, expectedChecksum, expectedGroup)

	// Then we should be able to get that entry via the various methods...
	require.True(t, ft.HasPath(expectedPath, expectedGroup))
	require.True(t, ft.HasPathAnywhere(expectedPath))
	require.True(t, ft.HasChecksum(expectedChecksum, expectedGroup))
	require.True(t, ft.HasChecksumAnywhere(expectedChecksum))

	// Also, excluding that group should return false
	require.False(t, ft.HasChecksumAnywhereExcept(expectedChecksum, expectedGroup))

	// But, excluding a different group means we should find it
	require.True(t, ft.HasChecksumAnywhereExcept(expectedChecksum, expectedGroup+"a"))
}

func TestFileTracker_DuplicatePathsByChecksum(t *testing.T) {
	ft := NewFileTracker()
	duplicateChecksum := "7d97e98f8af710c7e7fe703abc8f639e0ee507c4"
	uniqueChecksum := "3f786850e387550fdab836ed7e6dc881de23001b"

	ft.Set("/foo/bar", duplicateChecksum, "group-a")
	ft.Set("/foo/baz", uniqueChecksum, "group-a")
	ft.Set("/qux/bar", duplicateChecksum, "group-b")

	duplicates := ft.DuplicatePathsByChecksum()

	require.Equal(t, map[string][]string{
		duplicateChecksum: {"/foo/bar", "/qux/bar"},
	}, duplicates)
}

func TestNewFileTracker(t *testing.T) {
	require.NotNil(t, NewFileTracker())
}

func TestNewFileTrackerGroup(t *testing.T) {
	require.NotNil(t, NewFileTrackerGroup())
}
