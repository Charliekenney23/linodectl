package edit

import (
	cmdutil "github.com/Charliekenney23/linodectl/internal/cmd/util"
	"github.com/spf13/cobra"
)

func NewCmdEdit(f cmdutil.Factory, ioStreams cmdutil.IOStreams) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "edit [resource]",
		Short: "Edit a resource",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	cmds.AddCommand(NewCmdEditInstance(f, ioStreams))

	return cmds
}
