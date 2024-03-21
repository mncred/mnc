// package init provides command to initialize current dir for use with mnc.
package cmdinit

import (
	_ "embed"
	"strings"

	"github.com/spf13/cobra"
)

//go:embed help.txt
var help string

func Register() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "init <project_name>",
		Args:    cobra.ExactArgs(1),
		Aliases: []string{},
		Short:   strings.Split(help, "\n")[0],
		Long:    help,
		RunE:    run,
	}
	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	// projectName := args[0]

	return nil
}
