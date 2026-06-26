package cmd

import (
	_ "embed"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/scottgreenup/organiser/internal/checksum"
	"github.com/scottgreenup/organiser/internal/commands/filescmd"
	"github.com/scottgreenup/organiser/internal/filetype"
	"github.com/scottgreenup/organiser/internal/term"
)

//go:embed findDuplicates.long
var long string

var ignoreSameDir bool

// findDuplicatesCmd represents the findDuplicates command
var findDuplicatesCmd = &cobra.Command{
	Use:   "findDuplicates",
	Short: "Identifies files that are the same or similar.",
	Long:  long,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		if err := assertReadableDirectories(args); err != nil {
			return err
		}

		return findDuplicatesCmdActual(args)
	},
}

func findDuplicatesCmdActual(dirPaths []string) error {
	l := term.NewStandardErrorLogger()
	ft := filescmd.NewFileTracker()

	for _, dirPath := range dirPaths {
		if err := processDirectory(l, ft, dirPath); err != nil {
			return err
		}
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(ft.DuplicatePathsByChecksum())
}

func processDirectory(l term.Logger, ft *filescmd.FileTracker, directoryPath string) error {
	// Let's walk the entire directory and process the files we've found.
	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			l.Printf("Walk: skipping %q due to error with walk: %+v\n", path, err, term.WithForegroundColor(term.FgRed))
			return nil
		}
		if ok, _ := filetype.IsDir(path); ok {
			l.Printf("Walk: skipping %q as it is a directory\n", path, term.WithForegroundColor(term.FgYellow))
			return nil
		}

		sum, err := checksum.FileDigest(path)
		if err != nil {
			l.Printf("Walk: skipping %q due to error with checksum: %+v\n", path, err, term.WithForegroundColor(term.FgRed))
			return nil
		}

		if existingPath, ok := ft.GetPathByChecksum(sum, directoryPath); ok {
			l.Printf("Walk: %q is a duplicate of %q (reason: checksum)\n", path, existingPath, term.WithForegroundColor(term.FgRed))
		}
		ft.Set(path, sum, directoryPath)

		return nil
	})

	return err
}

func init() {
	filesCmd.AddCommand(findDuplicatesCmd)
	findDuplicatesCmd.Flags().BoolVar(&ignoreSameDir, "ignore-same-dir", false, "Ignore duplicates found in the same directory")
}

func assertReadableDirectories(dirs []string) error {
	for _, dir := range dirs {
		isDir, err := filetype.IsDir(dir)
		if err != nil {
			return err
		}
		if !isDir {
			return errors.Errorf("%s is not a directory", dir)
		}
	}
	return nil
}
