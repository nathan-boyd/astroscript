package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var appFs = afero.NewOsFs()

// the directory to be operated on
var directory string

// a string postfixed to each file name
var fileNameSubstring = "_thn"

// a list of subDirectories which the application is allowed to delete jpg in
var subDirectories = [...]string{
	"Light",
	"Dark",
	"Bias",
	"Flat",
}

// SubdirectoryNotfoundMessage is the error message returned when a required subdirectory is not found in the input path
var SubdirectoryNotfoundMessage = fmt.Sprintf("%s %s", "required subdirectory not found, path must contain one of the following sub-directories", strings.Join(subDirectories[:], ", "))

func stringInSlice(incString string, incList []string) bool {
	for _, b := range incList {
		if b == incString {
			return true
		}
	}
	return false
}

func sliceInSlice(sliceOne []string, sliceTwo []string) bool {
	for _, v1 := range sliceOne {
		if stringInSlice(v1, sliceTwo) {
			return true
		}
	}
	return false
}

func run(cmd *cobra.Command, args []string) (err error) {

	if _, err := appFs.Stat(directory); err != nil {
		return fmt.Errorf("directory does not exist %s", directory)
	}

	err = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		return removeJpgs(path)
	})

	return
}

func removeJpgs(path string) (err error) {

	// check that path contains one of the required astro sub directories
	s := strings.Split(path, "/")
	if !sliceInSlice(s, subDirectories[:]) {
		return
	}

	if filepath.Ext(path) != ".jpg" {
		return
	}

	if !strings.Contains(filepath.Base(path), fileNameSubstring) {
		return
	}

	return appFs.Remove(path)
}

// NewRemoveJpgCmd initializes an instance of a command which removes jpg files from a directory
func NewRemoveJpgCmd() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "removeJpg",
		Short: "something",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			return run(cmd, args)
		},
	}

	cmd.Flags().StringVar(&directory, "dir", "", "The directory to remove JPG files from")
	rootCmd.MarkFlagRequired("dir")

	return cmd
}

func init() {
	cmd := NewRemoveJpgCmd()
	rootCmd.AddCommand(cmd)
}
