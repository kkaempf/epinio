package cli

import (
	"fmt"

	"github.com/epinio/epinio/deployments"
	"github.com/epinio/epinio/helpers/kubernetes"
	"github.com/epinio/epinio/helpers/termui"
	"github.com/epinio/epinio/internal/cli/admincmd"
	"github.com/epinio/epinio/internal/duration"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var CmdDisable = &cobra.Command{
	Use:           "disable",
	Short:         "disable Epinio features",
	Long:          `disable Epinio features which where enabled with "epinio enable"`,
	SilenceErrors: true,
	SilenceUsage:  true,
	Args:          cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmd.Usage(); err != nil {
			return err
		}
		return fmt.Errorf(`Unknown method "%s"`, args[0])
	},
}

// TODO: Implement a flag to also delete provisioned services [TBD]
var CmdDisableInCluster = &cobra.Command{
	Use:   "services-incluster",
	Short: "disable in-cluster services in Epinio",
	Long:  `disable in-cluster services in Epinio which will disable provisioning services on the same cluster as Epinio. Doesn't delete already provisioned services by default.`,
	Args:  cobra.ExactArgs(0),
	RunE:  DisableInCluster,
}

var CmdDisableGoogle = &cobra.Command{
	Use:   "services-google",
	Short: "disable Google Cloud service in Epinio",
	Long:  `disable Google Cloud services in Epinio which will disable the provisioning of those services. Doesn't delete already provisioned services by default.`,
	Args:  cobra.ExactArgs(0),
	RunE:  DisableGoogle,
}

func init() {
	CmdDisable.AddCommand(CmdDisableInCluster)
	CmdDisable.AddCommand(CmdDisableGoogle)
}

// DisableInCluster implements: epinio disable services-incluster
func DisableInCluster(cmd *cobra.Command, args []string) error {
	return UninstallDeployment(
		cmd, &deployments.Minibroker{Timeout: duration.ToDeployment()},
		"in-cluster services functionality has been disabled")
}

// DisableGoogle implements: epinio disable services-google
func DisableGoogle(cmd *cobra.Command, args []string) error {
	return UninstallDeployment(
		cmd, &deployments.GoogleServices{Timeout: duration.ToDeployment()},
		"Google Cloud services functionality has been disabled")
}

// UninstallDeployment is a helper which removes the single referenced deployment
func UninstallDeployment(cmd *cobra.Command, deployment kubernetes.Deployment, successMessage string) error {
	cmd.SilenceUsage = true

	uiUI := termui.NewUI()
	installClient, installCleanup, err := admincmd.NewInstallClient(cmd.Context(), &kubernetes.InstallationOptions{})
	defer func() {
		if installCleanup != nil {
			installCleanup()
		}
	}()

	if err != nil {
		return errors.Wrap(err, "error initializing cli")
	}

	uiUI.Note().Msg(deployment.ID() + " uninstalling...")

	if err := installClient.UninstallDeployment(cmd.Context(), deployment, installClient.Log); err != nil {
		return errors.Wrap(err, "failed to remove")
	}

	uiUI.Note().Msg(successMessage)

	return nil
}
