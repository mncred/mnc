// package describe provides command to identify every minecraft file.
package describe

import (
	_ "embed"
	"encoding/json"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mncred/mnc/internal/describer"
)

//go:embed help.txt
var help string

func Register() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "describe filename",
		Args:    cobra.ExactArgs(1),
		Aliases: []string{"desc"},
		Short:   strings.Split(help, "\n")[0],
		Long:    help,
		RunE:    run,
	}
	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	filename := args[0]

	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := describer.Describe(file)
	if err != nil {
		return err
	}

	desc, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return err
	}

	os.Stdout.Write(desc)
	os.Stdout.WriteString("\n")

	return nil
}
