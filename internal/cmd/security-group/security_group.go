package security_group

import (
	"github.com/stackitcloud/stackit-cli/internal/cmd/security-group/create"
	"github.com/stackitcloud/stackit-cli/internal/cmd/security-group/delete"
	"github.com/stackitcloud/stackit-cli/internal/cmd/security-group/describe"
	"github.com/stackitcloud/stackit-cli/internal/cmd/security-group/list"
	"github.com/stackitcloud/stackit-cli/internal/cmd/security-group/rule"
	"github.com/stackitcloud/stackit-cli/internal/cmd/security-group/update"
	"github.com/stackitcloud/stackit-cli/internal/pkg/args"
	"github.com/stackitcloud/stackit-cli/internal/pkg/print"

	"github.com/stackitcloud/stackit-cli/internal/pkg/utils"

	"github.com/spf13/cobra"
)

func NewCmd(p *print.Printer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-group",
		Short: "Manage security groups",
		Long:  "Manage the lifecycle of security groups and rules.",
		Args:  args.NoArgs,
		Run:   utils.CmdHelp,
	}
	addSubcommands(cmd, p)
	return cmd
}

func addSubcommands(cmd *cobra.Command, p *print.Printer) {
	cmd.AddCommand(
		rule.NewCmd(p),
		create.NewCmd(p),
		delete.NewCmd(p),
		describe.NewCmd(p),
		list.NewCmd(p),
		update.NewCmd(p),
	)
}
