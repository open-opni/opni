package metrics

import (
	"bytes"
	"fmt"
	"text/template"

	alertingv1 "github.com/open-panoptes/opni/pkg/apis/alerting/v1"
	"github.com/prometheus/common/model"
	"google.golang.org/protobuf/types/known/durationpb"
)

const FilesystemMatcherName = "node_filesystem_free_bytes"
const NodeExporterMountpointLabel = "mountpoint"
const NodeExportFilesystemDeviceLabel = "device"

var FilesystemRuleAnnotations = map[string]string{}

type filesystemSaturation struct {
	Filters       string
	Operation     string
	ExpectedValue string
}

type filesystemSaturationSpike struct {
	Filters       string
	Operation     string
	ExpectedValue string
	SpikeWindow   string
	NumSpike      string
}

func NewFsRule(
	nodeFilters map[string]*alertingv1.FilesystemInfo,
	operation string,
	expectedValue float64,
	duration *durationpb.Duration,
	annotations map[string]string,
) (*AlertingRule, error) {
	dur := model.Duration(duration.AsDuration())
	filters := NewPrometheusFilters()
	filters.AddFilter(NodeFilter)
	filters.AddFilter(NodeExporterMountpointLabel)
	filters.AddFilter(NodeExportFilesystemDeviceLabel)
	for node, info := range nodeFilters {
		filters.Or(NodeFilter, node)
		for _, device := range info.Devices {
			filters.Or(NodeExportFilesystemDeviceLabel, device)
		}
		for _, mountpoint := range info.Mountpoints {
			filters.Or(NodeFilter, mountpoint)
		}
	}
	filters.Match(NodeFilter)
	filters.Match(NodeExporterMountpointLabel)
	filters.Match(NodeExportFilesystemDeviceLabel)
	tmpl := template.Must(template.New("").Parse(`
		sum(
			1- (
				node_filesystem_free_bytes{{ .Filters }}
				) 
				/ 
				node_filesystem_size_bytes
			)  {{  .Operation }} bool {{ .ExpectedValue }} 
	`))

	var b bytes.Buffer
	err := tmpl.Execute(&b, filesystemSaturation{
		Filters:       filters.Build(),
		Operation:     operation,
		ExpectedValue: fmt.Sprintf("%.7f", expectedValue),
	})
	if err != nil {
		return nil, err
	}
	return &AlertingRule{
		Alert:       "",
		Expr:        PostProcessRuleString(b.String()),
		For:         dur,
		Labels:      annotations,
		Annotations: annotations,
	}, nil
}

func NewFsSpikeRule(
	nodeFilters map[string]*alertingv1.FilesystemInfo,
	operation string,
	expectedValue float64,
	numSpikes int64,
	duration *durationpb.Duration,
	spikeWindow *durationpb.Duration,
	annotations map[string]string,
) (*AlertingRule, error) {
	dur := model.Duration(duration.AsDuration())
	spikeDur := model.Duration(spikeWindow.AsDuration())
	filters := NewPrometheusFilters()
	filters.AddFilter(NodeFilter)
	filters.AddFilter(NodeExporterMountpointLabel)
	filters.AddFilter(NodeExportFilesystemDeviceLabel)
	for node, info := range nodeFilters {
		filters.Or(NodeFilter, node)
		for _, device := range info.Devices {
			filters.Or(NodeExportFilesystemDeviceLabel, device)
		}
		for _, mountpoint := range info.Mountpoints {
			filters.Or(NodeFilter, mountpoint)
		}
	}
	filters.Match(NodeFilter)
	filters.Match(NodeExporterMountpointLabel)
	filters.Match(NodeExportFilesystemDeviceLabel)
	tmpl := template.Must(template.New("").Parse(`
		count_over_time((
			sum(
			1- (
				node_filesystem_free_bytes{{ .Filters }}
				) 
				/ 
				node_filesystem_size_bytes
			)  {{  .Operation }} bool {{ .ExpectedValue }}
		)[{{ .SpikeWindow }}:5s]) > {{ .NumSpike }}
	`))

	var b bytes.Buffer
	err := tmpl.Execute(&b, filesystemSaturationSpike{
		Filters:       filters.Build(),
		Operation:     operation,
		ExpectedValue: fmt.Sprintf("%.7f", expectedValue),
		SpikeWindow:   spikeDur.String(),
		NumSpike:      fmt.Sprintf("%d", numSpikes),
	})
	if err != nil {
		return nil, err
	}
	return &AlertingRule{
		Alert:       "",
		Expr:        PostProcessRuleString(b.String()),
		For:         dur,
		Labels:      annotations,
		Annotations: annotations,
	}, nil
}
