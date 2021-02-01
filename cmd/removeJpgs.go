package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var in string

func NewRemoveJpgsCmd() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "removeJpgs",
		Short: "something",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(cmd.OutOrStdout(), in)
			return nil
		},
	}
	cmd.Flags().StringVar(&in, "in", "", "This is a very important input.")

	return cmd
}

func init() {
	rootCmd.AddCommand(NewRemoveJpgsCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeJpgsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeJpgsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
