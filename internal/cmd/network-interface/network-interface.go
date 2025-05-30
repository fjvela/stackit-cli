package networkinterface

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-cli/internal/cmd/network-interface/create"
	"github.com/stackitcloud/stackit-cli/internal/cmd/network-interface/delete"
	"github.com/stackitcloud/stackit-cli/internal/cmd/network-interface/describe"
	"github.com/stackitcloud/stackit-cli/internal/cmd/network-interface/list"
	"github.com/stackitcloud/stackit-cli/internal/cmd/network-interface/update"
	"github.com/stackitcloud/stackit-cli/internal/pkg/args"
	"github.com/stackitcloud/stackit-cli/internal/pkg/print"
	"github.com/stackitcloud/stackit-cli/internal/pkg/utils"
)

func NewCmd(p *print.Printer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "network-interface",
		Short: "Provides functionality for network interfaces",
		Long:  "Provides functionality for network interfaces.",
		Args:  args.NoArgs,
		Run:   utils.CmdHelp,
	}
	addSubcommands(cmd, p)
	return cmd
}

func addSubcommands(cmd *cobra.Command, p *print.Printer) {
	cmd.AddCommand(create.NewCmd(p))
	cmd.AddCommand(delete.NewCmd(p))
	cmd.AddCommand(update.NewCmd(p))
	cmd.AddCommand(describe.NewCmd(p))
	cmd.AddCommand(list.NewCmd(p))
}
