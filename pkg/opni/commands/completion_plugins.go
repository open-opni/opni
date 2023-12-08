//go:build !minimal

package commands

import (
	"strings"

	"slices"

	corev1 "github.com/open-panoptes/opni/pkg/apis/core/v1"
	"github.com/open-panoptes/opni/plugins/metrics/apis/remoteread"
	"github.com/spf13/cobra"
)

func completeImportTargets(cmd *cobra.Command, args []string, toComplete string, _ ...func(token *corev1.BootstrapToken) bool) ([]string, cobra.ShellCompDirective) {
	if err := importPreRunE(cmd, nil); err != nil {
		return nil, cobra.ShellCompDirectiveError | cobra.ShellCompDirectiveNoFileComp
	}

	var cluster string
	if len(args) >= 1 {
		cluster = args[1]
	}

	targetList, err := remoteReadClient.ListTargets(cmd.Context(), &remoteread.TargetListRequest{
		ClusterId: cluster,
	})
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	var targets []string
	for _, target := range targetList.Targets {
		name := target.Meta.Name

		if slices.Contains(args, name) {
			continue
		}

		if strings.HasPrefix(name, toComplete) {
			targets = append(targets, name)
		}
	}

	return targets, cobra.ShellCompDirectiveNoFileComp
}
