package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

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

// EnvWrapper abstracts the operating system and file system away from the application
type EnvWrapper interface {
	GetWorkingDirectory() (wb string, err error)
	DirectoryExists(path string) (directoryExists bool)
	GetFilePathSeperator() (seperator rune)
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

// GetFilePathSeperator returns the file path seperator for the operating system
func (t *EnvWrapperImpl) GetFilePathSeperator() (seperator rune) {
	return os.PathSeparator
}

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

	s := strings.Split(directory, string(envWrapper.GetFilePathSeperator()))
	if !sliceInSlice(s, subDirectories[:]) {
		return fmt.Errorf(SubdirectoryNotfoundMessage)
	}

	return
}

// NewRemoveJpgCmd initializes an instance of a command which removes jpg files from a directory
func NewRemoveJpgCmd(envWrapper EnvWrapper) *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "removeJpg",
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
	cmd := NewRemoveJpgCmd(envWrapper)

	rootCmd.AddCommand(cmd)
}
