package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var inPath string
var outPath string

func consolodateImages(cmd *cobra.Command, args []string) (err error) {

	if _, err := appFs.Stat(inPath); err != nil {
		return fmt.Errorf("directory does not exist %s", directory)
	}

	err = filepath.Walk(inPath, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() {
			newPath := fmt.Sprintf("%s/%s", outPath, info.Name())
			return copy(path, newPath)
		}

		fmt.Fprintf(cmd.OutOrStdout(), directory)

		return err

	})

	return
}

func copy(src, dst string) error {

	fmt.Println(src)

	in, err := afero.ReadFile(appFs, src)
	if err != nil {
		return err
	}

	out, err := appFs.Create(dst)
	if err != nil {
		return err
	}

	defer out.Close()

	_ = afero.WriteFile(appFs, dst, in, 0644)
	if err != nil {
		return err
	}

	return out.Close()
}

// NewConsolodateCmd returns a consolodateImages cmd
func NewConsolodateCmd() *cobra.Command {

	// consolodateImagesCmd represents the consolodateImages command
	cmd := &cobra.Command{
		Use:   "consolodateImages",
		Short: "bar",
		Long: `A longer description that spans multiple lines and likely contains examples
		  to quickly create a Cobra application.`,

		RunE: func(cmd *cobra.Command, args []string) (err error) {
			return consolodateImages(cmd, args)
		},
	}

	cmd.Flags().StringVar(&inPath, "inPath", "", "The directory to copy jpg files from")
	rootCmd.MarkFlagRequired("inPath")

	cmd.Flags().StringVar(&outPath, "outPath", "", "The directory to copy jpg files to")
	rootCmd.MarkFlagRequired("outPath")

	return cmd
}

func init() {
	cmd := NewConsolodateCmd()
	rootCmd.AddCommand(cmd)
}
