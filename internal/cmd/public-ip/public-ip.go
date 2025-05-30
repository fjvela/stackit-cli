package publicip

import (
	"github.com/stackitcloud/stackit-cli/internal/cmd/public-ip/associate"
	"github.com/stackitcloud/stackit-cli/internal/cmd/public-ip/create"
	"github.com/stackitcloud/stackit-cli/internal/cmd/public-ip/delete"
	"github.com/stackitcloud/stackit-cli/internal/cmd/public-ip/describe"
	"github.com/stackitcloud/stackit-cli/internal/cmd/public-ip/disassociate"
	"github.com/stackitcloud/stackit-cli/internal/cmd/public-ip/list"
	"github.com/stackitcloud/stackit-cli/internal/cmd/public-ip/update"
	"github.com/stackitcloud/stackit-cli/internal/pkg/args"
	"github.com/stackitcloud/stackit-cli/internal/pkg/print"
	"github.com/stackitcloud/stackit-cli/internal/pkg/utils"

	"github.com/spf13/cobra"
)

func NewCmd(p *print.Printer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "public-ip",
		Short: "Provides functionality for public IPs",
		Long:  "Provides functionality for public IPs.",
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
	cmd.AddCommand(update.NewCmd(p))
	cmd.AddCommand(associate.NewCmd(p))
	cmd.AddCommand(disassociate.NewCmd(p))
}
