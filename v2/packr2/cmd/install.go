package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:                "install",
	Short:              "Wraps the go install command with packr",
	DisableFlagParsing: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var cargs []string
		for _, a := range args {
			if a == "--legacy" {
				globalOptions.Legacy = true
				continue
			}
			cargs = append(cargs, a)
		}
		if err := pack(); err != nil {
			return errors.WithStack(err)
		}
		return goCmd("install", cargs...)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
