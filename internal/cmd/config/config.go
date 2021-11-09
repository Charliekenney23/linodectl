package config

import (
	cmdutil "github.com/Charliekenney23/linodectl/internal/cmd/util"
	"github.com/spf13/cobra"
)

func NewCmdConfig(f cmdutil.Factory, ioStreams cmdutil.IOStreams) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "config",
		Short: "Manage configurations",
	}

	cmds.AddCommand(NewCmdConfigSetProfile(f, ioStreams))
	cmds.AddCommand(NewCmdConfigListProfiles(f, ioStreams))
	cmds.AddCommand(NewCmdConfigGetProfile(f, ioStreams))
	return cmds
}
