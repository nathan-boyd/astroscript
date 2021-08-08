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
		fmt.Printf("skipping path %s, due to filter\n", path)
		return
	}

	if filepath.Ext(path) != ".jpg" {
		fmt.Printf("skipping non jpg file %s\n", path)
		return
	}

	if !strings.Contains(filepath.Base(path), fileNameSubstring) {
		fmt.Printf("skipping file missing required substring %s\n", path)
		return
	}

	fmt.Printf("removing file %s\n", path)

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
