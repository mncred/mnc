// package cmd provides root command and other commands registration.
package cmd

import (
	_ "embed"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mncred/mnc/cmd/describe"
)

//go:embed help.txt
var help string

func Register() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mnc",
		Short: strings.Split(help, "\n")[0],
		Long:  help,
	}

	cmd.AddCommand(describe.Register())

	return cmd
}
