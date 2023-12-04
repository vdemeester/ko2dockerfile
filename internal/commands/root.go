package commands

import (
	"github.com/spf13/cobra"
)

var Root = New()

func New() *cobra.Command {
	var verbose bool
	root := &cobra.Command{
		Use:               "ko",
		Short:             "Rapidly iterate with Go, Containers, and Kubernetes.",
		SilenceUsage:      true, // Don't show usage on errors
		DisableAutoGenTag: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// if verbose {
			// 	logs.Warn.SetOutput(os.Stderr)
			// 	logs.Debug.SetOutput(os.Stderr)
			// }
			// logs.Progress.SetOutput(os.Stderr)
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	root.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable debug logs")

	// AddKubeCommands(root)
	// root.AddCommand(…)

	return root
}
