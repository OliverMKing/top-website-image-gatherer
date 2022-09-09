package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	outputName = "output"
	topName    = "top"
)

func init() {
	f := gatherCmd.Flags()
	f.StringVarP(&output, outputName, "o", "", "specify the path to the directory to place screenshots in (required)")
	gatherCmd.MarkFlagRequired(outputName)
	f.IntVarP(&top, topName, "t", 10, "specify the number of top websites to retrieve screenshots from")

	rootCmd.AddCommand(gatherCmd)
}

var gatherCmd = &cobra.Command{
	Use:   "gather [flags]",
	Short: "Retrieve screenshots of top websites",
	RunE:  run,
}

// flags
var (
	output string
	top    int
)

func run(cmd *cobra.Command, args []string) error {
	fmt.Print("Hello world!")

	return nil
}
