/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	_ "embed"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/scottgreenup/organiser/internal/checksum"
	"github.com/scottgreenup/organiser/internal/commands/filescmd"
	"github.com/scottgreenup/organiser/internal/filetype"
	"github.com/scottgreenup/organiser/internal/term"
)

//go:embed wallpaperSort.long
var wallpaperSortCmdLong string

// findDuplicatesCmd represents the findDuplicates command
var wallpaperSortCmd = &cobra.Command{
	Use:   "wallpaperSort",
	Short: "wallpaperSort will organise your wallpaper",
	Long:  long,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		// Validate our inputs from our "view", then pass it down to the
		// business logic.
		dir := args[0]
		if err := validateDirectory(dir); err != nil {
			return err
		}

		// Business logic, it doesn't need to know about Cobra, just the inputs.
		return wallpaperSortActual(dir)
	},
}

func wallpaperSortActual(dir string) error {
	l := term.NewStandardLogger()
	ft := filescmd.NewFileTracker()

	if err := trackFiles(l, ft, dir); err != nil {
		return err
	}

	return nil
}

func trackFiles(l term.Logger, ft *filescmd.FileTracker, directoryPath string) error {
	// Let's walk the entire directory and process the files we've found.
	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			l.Printf("Walk: skipping %q due to error with walk: %+v\n", path, err, term.WithForegroundColor(term.FgRed))
			return nil
		}
		if ok, _ := filetype.IsDir(path); ok {
			l.Printf("Walk: skipping %q due to directory\n", path, term.WithForegroundColor(term.FgYellow))
			return nil
		}

		sumBytes, err := checksum.File(path)
		if err != nil {
			l.Printf("Walk: skipping %q due to error with checksum: %+v\n", path, err, term.WithForegroundColor(term.FgRed))
			return nil
		}

		sum := string(sumBytes)
		group := directoryPath

		if existingPath, ok := ft.GetPathByChecksum(sum, group); ok {
			l.Printf("Walk: %q is a duplicate of %q (reason: checksum)\n", path, existingPath, term.WithForegroundColor(term.FgRed))
		} else {
			ft.Set(path, string(sum), group)
		}

		return nil
	})

	return err
}

func init() {
	filesCmd.AddCommand(wallpaperSortCmd)
}

func validateDirectory(dir string) error {
	isDir, err := filetype.IsDir(dir)
	if err != nil {
		return err
	}
	if !isDir {
		return errors.Errorf("%s is not a directory", dir)
	}
	return nil
}
