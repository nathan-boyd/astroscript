package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var directory string

var fileNameSubstring = "_thn"

// only delete jpgs from one of the following subDirectories
var subDirectories = [...]string{"Light", "Dark", "Bias", "Flat"}

func run(cmd *cobra.Command, args []string) error {
	fmt.Fprintf(cmd.OutOrStdout(), directory)
	return nil
}

// NewRemoveJpgsCmd initializes an instance of a command which removes jpg files from a directory
func NewRemoveJpgsCmd() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "removeJpgs",
		Short: "something",
		RunE:  run,
	}
	cmd.Flags().StringVar(&directory, "dir", "", "The directory to remove JPG files from")

	return cmd
}

func init() {
	cmd := NewRemoveJpgsCmd()
	rootCmd.AddCommand(cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeJpgsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeJpgsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
