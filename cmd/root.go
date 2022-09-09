package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "twig",
	Short: "Top website image gatherer (TWIG) retrieves screenshots from the top websites",
	Long: `Top website image gatherer (TWIG) retrieves screenshots from the top websites

To retrieve screenshots run the 'twig gather' command

	$ twig gather -o ./example/output/path

Running this will take screenshots of the top websites and output them to the specified directory. Use the -t flag to specify the number of websites to pull from.`,
}

// Execute runs the root command and should only happen once
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
