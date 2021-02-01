package cmd

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// the directory to be operated on
var directory string

// a string postfixed to each file name
var fileNameSubstring = "_thn"

// a list of subDirectories which the application is allowed to delete jpgs in
var subDirectories = [...]string{
	"Light",
	"Dark",
	"Bias",
	"Flat",
}

// EnvWrapper abstracts the operating system and file system away from the application
type EnvWrapper interface {
	GetWorkingDirectory() (wb string, err error)
	DirectoryExists(path string) (directoryExists bool)
}

// EnvWrapperImpl is an implementation of an EnvWrapper
type EnvWrapperImpl struct{}

// GetWorkingDirectory returns a fake working directory for testing
func (t *EnvWrapperImpl) GetWorkingDirectory() (wd string, err error) {
	return os.Getwd()
}

// DirectoryExists returns true if directory exists
func (t *EnvWrapperImpl) DirectoryExists(path string) (directoryExists bool) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func run(cmd *cobra.Command, args []string, envWrapper EnvWrapper) (err error) {

	if 0 == len(directory) {
		fmt.Fprintf(cmd.OutOrStdout(), "optional directory argument not provided, using current working directory")
		directory, err = envWrapper.GetWorkingDirectory()
		if nil != err {
			return errors.Wrap(err, "failed to get working directory")
		}
	}

	if !envWrapper.DirectoryExists(directory) {
		return fmt.Errorf("directory does not exist %s", directory)
	}

	return
}

// NewRemoveJpgsCmd initializes an instance of a command which removes jpg files from a directory
func NewRemoveJpgsCmd(envWrapper EnvWrapper) *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "removeJpgs",
		Short: "something",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			return run(cmd, args, envWrapper)
		},
	}

	cmd.Flags().StringVar(&directory, "dir", "", "The directory to remove JPG files from")

	return cmd
}

func init() {
	envWrapper := &EnvWrapperImpl{}
	cmd := NewRemoveJpgsCmd(envWrapper)

	rootCmd.AddCommand(cmd)
}
