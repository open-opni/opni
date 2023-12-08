package complete

import (
	"context"
	"fmt"
	"time"

	"github.com/open-panoptes/opni/pkg/plugins/driverutil"
	"github.com/spf13/cobra"
)

func Revisions[
	C driverutil.HistoryClientInterface[T, HR],
	T driverutil.ConfigType[T],
	HR driverutil.HistoryResponseType[T],
](ctx context.Context, target driverutil.Target, client C) ([]string, cobra.ShellCompDirective) {
	history, err := client.ConfigurationHistory(ctx, &driverutil.ConfigurationHistoryRequest{
		Target:        target,
		IncludeValues: false,
	})
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}
	revisions := make([]string, len(history.GetEntries()))
	for i, entry := range history.GetEntries() {
		comp := fmt.Sprint(entry.GetRevision().GetRevision())
		ts := entry.GetRevision().GetTimestamp().AsTime()
		if !ts.IsZero() {
			comp = fmt.Sprintf("%s\t%s", comp, ts.Format(time.Stamp))
		}
		revisions[i] = comp
	}
	return revisions, cobra.ShellCompDirectiveNoFileComp
}
