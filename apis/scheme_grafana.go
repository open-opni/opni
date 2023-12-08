//go:build !minimal

package apis

import opnigrafanav1beta1 "github.com/open-panoptes/opni/apis/grafana/v1beta1"

func init() {
	addSchemeBuilders(opnigrafanav1beta1.AddToScheme)
}
