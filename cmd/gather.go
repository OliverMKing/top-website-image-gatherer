package cmd

import (
	"top-website-image-gatherer/pkg/gather"
	"top-website-image-gatherer/pkg/screenshot"
	"top-website-image-gatherer/pkg/site"

	"github.com/spf13/cobra"
)

const (
	outputName = "output"
	numName    = "sites-number"
	offsetName = "offset"
)

func init() {
	f := gatherCmd.Flags()
	f.StringVarP(&output, outputName, "o", "", "specify the path to the directory to place screenshots in (required)")
	gatherCmd.MarkFlagRequired(outputName)
	f.IntVarP(&num, numName, "n", 10, "specify the number of top websites to retrieve screenshots from")
	f.IntVarP(&offset, offsetName, "s", 0, "specify the number of websites to skip from the top of the list")

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
	num    int
	offset int
)

func run(cmd *cobra.Command, args []string) error {
	s := site.Top(num, offset)
	ss := screenshot.New()
	g := gather.New(s, ss)
	return g.Gather("./")
}
