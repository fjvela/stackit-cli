package networkarea

import (
	"github.com/stackitcloud/stackit-cli/internal/cmd/network-area/create"
	"github.com/stackitcloud/stackit-cli/internal/cmd/network-area/delete"
	"github.com/stackitcloud/stackit-cli/internal/cmd/network-area/describe"
	"github.com/stackitcloud/stackit-cli/internal/cmd/network-area/list"
	networkrange "github.com/stackitcloud/stackit-cli/internal/cmd/network-area/network-range"
	"github.com/stackitcloud/stackit-cli/internal/cmd/network-area/route"
	"github.com/stackitcloud/stackit-cli/internal/cmd/network-area/update"
	"github.com/stackitcloud/stackit-cli/internal/pkg/args"
	"github.com/stackitcloud/stackit-cli/internal/pkg/print"
	"github.com/stackitcloud/stackit-cli/internal/pkg/utils"

	"github.com/spf13/cobra"
)

func NewCmd(p *print.Printer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "network-area",
		Short: "Provides functionality for STACKIT Network Area (SNA)",
		Long:  "Provides functionality for STACKIT Network Area (SNA).",
		Args:  args.NoArgs,
		Run:   utils.CmdHelp,
	}
	addSubcommands(cmd, p)
	return cmd
}

func addSubcommands(cmd *cobra.Command, p *print.Printer) {
	cmd.AddCommand(create.NewCmd(p))
	cmd.AddCommand(delete.NewCmd(p))
	cmd.AddCommand(describe.NewCmd(p))
	cmd.AddCommand(list.NewCmd(p))
	cmd.AddCommand(networkrange.NewCmd(p))
	cmd.AddCommand(route.NewCmd(p))
	cmd.AddCommand(update.NewCmd(p))
}
