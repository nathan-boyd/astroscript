package cmd

import (
	"fmt"
	"log"
	"os"

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
}

type EnvWrapperImpl struct{}

// GetWorkingDirectory returns a fake working directory for testing
func (t *EnvWrapperImpl) GetWorkingDirectory() (wd string, err error) {
	return
}

func run(cmd *cobra.Command, args []string) error {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	log.Println(path)

	fmt.Fprintf(cmd.OutOrStdout(), directory)
	return nil
}

// NewRemoveJpgsCmd initializes an instance of a command which removes jpg files from a directory
func NewRemoveJpgsCmd(envWrapper EnvWrapper) *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "removeJpgs",
		Short: "something",
		RunE:  run,
	}

	cmd.Flags().StringVar(&directory, "dir", "", "The directory to remove JPG files from")

	return cmd
}

func init() {
	envWrapper := &EnvWrapperImpl{}
	cmd := NewRemoveJpgsCmd(envWrapper)
	rootCmd.AddCommand(cmd)
}
