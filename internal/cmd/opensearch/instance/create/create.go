package create

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/stackitcloud/stackit-cli/internal/pkg/args"
	cliErr "github.com/stackitcloud/stackit-cli/internal/pkg/errors"
	"github.com/stackitcloud/stackit-cli/internal/pkg/examples"
	"github.com/stackitcloud/stackit-cli/internal/pkg/flags"
	"github.com/stackitcloud/stackit-cli/internal/pkg/globalflags"
	"github.com/stackitcloud/stackit-cli/internal/pkg/print"
	"github.com/stackitcloud/stackit-cli/internal/pkg/projectname"
	"github.com/stackitcloud/stackit-cli/internal/pkg/services/opensearch/client"
	opensearchUtils "github.com/stackitcloud/stackit-cli/internal/pkg/services/opensearch/utils"
	"github.com/stackitcloud/stackit-cli/internal/pkg/spinner"
	"github.com/stackitcloud/stackit-cli/internal/pkg/utils"

	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-sdk-go/services/opensearch"
	"github.com/stackitcloud/stackit-sdk-go/services/opensearch/wait"
)

const (
	instanceNameFlag         = "name"
	enableMonitoringFlag     = "enable-monitoring"
	graphiteFlag             = "graphite"
	metricsFrequencyFlag     = "metrics-frequency"
	metricsPrefixFlag        = "metrics-prefix"
	monitoringInstanceIdFlag = "monitoring-instance-id"
	pluginFlag               = "plugin"
	sgwAclFlag               = "acl"
	syslogFlag               = "syslog"
	planIdFlag               = "plan-id"
	planNameFlag             = "plan-name"
	versionFlag              = "version"
)

type inputModel struct {
	*globalflags.GlobalFlagModel
	PlanName string
	Version  string

	InstanceName         *string
	EnableMonitoring     *bool
	Graphite             *string
	MetricsFrequency     *int64
	MetricsPrefix        *string
	MonitoringInstanceId *string
	Plugin               *[]string
	SgwAcl               *[]string
	Syslog               *[]string
	PlanId               *string
}

func NewCmd(p *print.Printer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates an OpenSearch instance",
		Long:  "Creates an OpenSearch instance.",
		Args:  args.NoArgs,
		Example: examples.Build(
			examples.NewExample(
				`Create an OpenSearch instance with name "my-instance" and specify plan by name and version`,
				"$ stackit opensearch instance create --name my-instance --plan-name stackit-opensearch-1.2.10-replica --version 2"),
			examples.NewExample(
				`Create an OpenSearch instance with name "my-instance" and specify plan by ID`,
				"$ stackit opensearch instance create --name my-instance --plan-id xxx"),
			examples.NewExample(
				`Create an OpenSearch instance with name "my-instance" and specify IP range which is allowed to access it`,
				"$ stackit opensearch instance create --name my-instance --plan-id xxx --acl 1.2.3.0/24"),
		),
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := context.Background()
			model, err := parseInput(p, cmd)
			if err != nil {
				return err
			}

			// Configure API client
			apiClient, err := client.ConfigureClient(p)
			if err != nil {
				return err
			}

			projectLabel, err := projectname.GetProjectName(ctx, p, cmd)
			if err != nil {
				p.Debug(print.ErrorLevel, "get project name: %v", err)
				projectLabel = model.ProjectId
			}

			if !model.AssumeYes {
				prompt := fmt.Sprintf("Are you sure you want to create an OpenSearch instance for project %q?", projectLabel)
				err = p.PromptForConfirmation(prompt)
				if err != nil {
					return err
				}
			}

			// Call API
			req, err := buildRequest(ctx, model, apiClient)
			if err != nil {
				var dsaInvalidPlanError *cliErr.DSAInvalidPlanError
				if !errors.As(err, &dsaInvalidPlanError) {
					return fmt.Errorf("build OpenSearch instance creation request: %w", err)
				}
				return err
			}
			resp, err := req.Execute()
			if err != nil {
				return fmt.Errorf("create OpenSearch instance: %w", err)
			}
			instanceId := *resp.InstanceId

			// Wait for async operation, if async mode not enabled
			if !model.Async {
				s := spinner.New(p)
				s.Start("Creating instance")
				_, err = wait.CreateInstanceWaitHandler(ctx, apiClient, model.ProjectId, instanceId).WaitWithContext(ctx)
				if err != nil {
					return fmt.Errorf("wait for OpenSearch instance creation: %w", err)
				}
				s.Stop()
			}

			return outputResult(p, model.OutputFormat, model.Async, projectLabel, instanceId, resp)
		},
	}
	configureFlags(cmd)
	return cmd
}

func configureFlags(cmd *cobra.Command) {
	cmd.Flags().StringP(instanceNameFlag, "n", "", "Instance name")
	cmd.Flags().Bool(enableMonitoringFlag, false, "Enable monitoring")
	cmd.Flags().String(graphiteFlag, "", "Graphite host")
	cmd.Flags().Int64(metricsFrequencyFlag, 0, "Metrics frequency")
	cmd.Flags().String(metricsPrefixFlag, "", "Metrics prefix")
	cmd.Flags().Var(flags.UUIDFlag(), monitoringInstanceIdFlag, "Monitoring instance ID")
	cmd.Flags().StringSlice(pluginFlag, []string{}, "Plugin")
	cmd.Flags().Var(flags.CIDRSliceFlag(), sgwAclFlag, "List of IP networks in CIDR notation which are allowed to access this instance")
	cmd.Flags().StringSlice(syslogFlag, []string{}, "Syslog")
	cmd.Flags().Var(flags.UUIDFlag(), planIdFlag, "Plan ID")
	cmd.Flags().String(planNameFlag, "", "Plan name")
	cmd.Flags().String(versionFlag, "", "Instance OpenSearch version")

	err := flags.MarkFlagsRequired(cmd, instanceNameFlag)
	cobra.CheckErr(err)
}

func parseInput(p *print.Printer, cmd *cobra.Command) (*inputModel, error) {
	globalFlags := globalflags.Parse(p, cmd)
	if globalFlags.ProjectId == "" {
		return nil, &cliErr.ProjectIdError{}
	}

	planId := flags.FlagToStringPointer(p, cmd, planIdFlag)
	planName := flags.FlagToStringValue(p, cmd, planNameFlag)
	version := flags.FlagToStringValue(p, cmd, versionFlag)

	if planId == nil && (planName == "" || version == "") {
		return nil, &cliErr.DSAInputPlanError{
			Cmd: cmd,
		}
	}
	if planId != nil && (planName != "" || version != "") {
		return nil, &cliErr.DSAInputPlanError{
			Cmd: cmd,
		}
	}

	model := inputModel{
		GlobalFlagModel:      globalFlags,
		InstanceName:         flags.FlagToStringPointer(p, cmd, instanceNameFlag),
		EnableMonitoring:     flags.FlagToBoolPointer(p, cmd, enableMonitoringFlag),
		MonitoringInstanceId: flags.FlagToStringPointer(p, cmd, monitoringInstanceIdFlag),
		Graphite:             flags.FlagToStringPointer(p, cmd, graphiteFlag),
		MetricsFrequency:     flags.FlagToInt64Pointer(p, cmd, metricsFrequencyFlag),
		MetricsPrefix:        flags.FlagToStringPointer(p, cmd, metricsPrefixFlag),
		Plugin:               flags.FlagToStringSlicePointer(p, cmd, pluginFlag),
		SgwAcl:               flags.FlagToStringSlicePointer(p, cmd, sgwAclFlag),
		Syslog:               flags.FlagToStringSlicePointer(p, cmd, syslogFlag),
		PlanId:               planId,
		PlanName:             planName,
		Version:              version,
	}

	if p.IsVerbosityDebug() {
		modelStr, err := print.BuildDebugStrFromInputModel(model)
		if err != nil {
			p.Debug(print.ErrorLevel, "convert model to string for debugging: %v", err)
		} else {
			p.Debug(print.DebugLevel, "parsed input values: %s", modelStr)
		}
	}

	return &model, nil
}

type openSearchClient interface {
	CreateInstance(ctx context.Context, projectId string) opensearch.ApiCreateInstanceRequest
	ListOfferingsExecute(ctx context.Context, projectId string) (*opensearch.ListOfferingsResponse, error)
}

func buildRequest(ctx context.Context, model *inputModel, apiClient openSearchClient) (opensearch.ApiCreateInstanceRequest, error) {
	req := apiClient.CreateInstance(ctx, model.ProjectId)

	var planId *string
	var err error

	offerings, err := apiClient.ListOfferingsExecute(ctx, model.ProjectId)
	if err != nil {
		return req, fmt.Errorf("get OpenSearch offerings: %w", err)
	}

	if model.PlanId == nil {
		planId, err = opensearchUtils.LoadPlanId(model.PlanName, model.Version, offerings)
		if err != nil {
			var dsaInvalidPlanError *cliErr.DSAInvalidPlanError
			if !errors.As(err, &dsaInvalidPlanError) {
				return req, fmt.Errorf("load plan ID: %w", err)
			}
			return req, err
		}
	} else {
		err := opensearchUtils.ValidatePlanId(*model.PlanId, offerings)
		if err != nil {
			return req, err
		}
		planId = model.PlanId
	}

	var sgwAcl *string
	if model.SgwAcl != nil {
		sgwAcl = utils.Ptr(strings.Join(*model.SgwAcl, ","))
	}

	req = req.CreateInstancePayload(opensearch.CreateInstancePayload{
		InstanceName: model.InstanceName,
		Parameters: &opensearch.InstanceParameters{
			EnableMonitoring:     model.EnableMonitoring,
			Graphite:             model.Graphite,
			MonitoringInstanceId: model.MonitoringInstanceId,
			MetricsFrequency:     model.MetricsFrequency,
			MetricsPrefix:        model.MetricsPrefix,
			Plugins:              model.Plugin,
			SgwAcl:               sgwAcl,
			Syslog:               model.Syslog,
		},
		PlanId: planId,
	})
	return req, nil
}

func outputResult(p *print.Printer, outputFormat string, async bool, projectLabel, instanceId string, resp *opensearch.CreateInstanceResponse) error {
	switch outputFormat {
	case print.JSONOutputFormat:
		details, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			return fmt.Errorf("marshal OpenSearch instance: %w", err)
		}
		p.Outputln(string(details))

		return nil
	case print.YAMLOutputFormat:
		details, err := yaml.MarshalWithOptions(resp, yaml.IndentSequence(true), yaml.UseJSONMarshaler())
		if err != nil {
			return fmt.Errorf("marshal OpenSearch instance: %w", err)
		}
		p.Outputln(string(details))

		return nil
	default:
		operationState := "Created"
		if async {
			operationState = "Triggered creation of"
		}
		p.Outputf("%s instance for project %q. Instance ID: %s\n", operationState, projectLabel, instanceId)
		return nil
	}
}
