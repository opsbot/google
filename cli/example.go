package cli

import (
	"github.com/opsbot/google/api/admin/directory/user"
	"github.com/spf13/cobra"
)

// ExampleCommand returns a cobra command
func ExampleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "example",
		Short: "example",
		Run: func(cmd *cobra.Command, args []string) {
			user.List()
		},
	}
	return cmd
}
