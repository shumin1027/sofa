package version

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// Command creates version command
func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version.",
		Long:  `Print the version and build information.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			info := Get()
			json, err := json.Marshal(info)
			if err != nil {
				return err
			}
			fmt.Println(string(json))

			return nil
		},
	}
}
