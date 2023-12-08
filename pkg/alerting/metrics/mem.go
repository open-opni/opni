package metrics

import (
	"bytes"
	"fmt"
	"regexp"
	"text/template"

	"slices"

	alertingv1 "github.com/open-panoptes/opni/pkg/apis/alerting/v1"
	"github.com/prometheus/common/model"
	"google.golang.org/protobuf/types/known/durationpb"
)

const MemoryDeviceFilter = "device"
const MemoryMatcherName = "node_memory_.*_bytes"

var MemoryUsageTypeExtractor = regexp.MustCompile("node_memory_(.*)_bytes")
var MemoryMatcherRegexFilter = []string{"Buffers", "Cached", "MemFree", "Slab"}
var MemRuleAnnotations = map[string]string{}

type memorySaturation struct {
	Filters       string
	Operation     string
	ExpectedValue string
}

type memorySaturationSpike struct {
	Filters       string
	Operation     string
	ExpectedValue string
	SpikeWindow   string
	NumSpike      string
}

func NewMemRule(
	deviceFilters map[string]*alertingv1.MemoryInfo,
	usageTypes []string,
	operation string,
	expectedRatio float64,
	duration *durationpb.Duration,
	annotations map[string]string,
) (*AlertingRule, error) {

	dur := model.Duration(duration.AsDuration())

	filters := NewPrometheusFilters()
	filters.AddFilter(NodeFilter)
	for node, state := range deviceFilters {
		filters.Or(NodeFilter, node)
		for _, device := range state.Devices {
			filters.Or(MemoryDeviceFilter, device)
		}
	}
	filters.Match(NodeFilter)
	outputFilters := filters.Build()
	aggrMetrics := ""
	for _, utype := range usageTypes {
		if !slices.Contains(MemoryMatcherRegexFilter, utype) {
			continue // FIXME: warn
		}
		if aggrMetrics != "" {
			aggrMetrics += " + "
		}
		aggrMetrics += fmt.Sprintf("node_memory_%s_bytes%s", utype, outputFilters)
	}
	tmpl := template.Must(template.New("").Parse(`
	(1 -  (
		  {{ .AggrMetrics }}
		)
	/
	  node_memory_MemTotal_bytes{{ .Filters }}
	) {{ .Operation }} bool {{ .ExpectedValue }}
`))
	var b bytes.Buffer
	err := tmpl.Execute(&b, map[string]string{
		"Filters":       filters.Build(),
		"AggrMetrics":   aggrMetrics,
		"Operation":     operation,
		"ExpectedValue": fmt.Sprintf("%.7f", expectedRatio),
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

func NewMemSpikeRule(
	deviceFilters map[string]*alertingv1.MemoryInfo,
	usageTypes []string,
	operation string,
	expectedRatio float64,
	numSpikes int64,
	duration *durationpb.Duration,
	spikeWindow *durationpb.Duration,
	annotations map[string]string,
) (*AlertingRule, error) {
	dur := model.Duration(duration.AsDuration())
	spikeDur := model.Duration(spikeWindow.AsDuration())

	filters := NewPrometheusFilters()
	filters.AddFilter(NodeFilter)
	for node, state := range deviceFilters {
		filters.Or(NodeFilter, node)
		for _, device := range state.Devices {
			filters.Or(MemoryDeviceFilter, device)
		}
	}
	filters.Match(NodeFilter)
	outputFilters := filters.Build()
	aggrMetrics := ""
	for _, utype := range usageTypes {
		if !slices.Contains(MemoryMatcherRegexFilter, utype) {
			continue // FIXME: warn
		}
		if aggrMetrics != "" {
			aggrMetrics += " + "
		}
		aggrMetrics += fmt.Sprintf("node_memory_%s_bytes%s", utype, outputFilters)
	}
	tmpl := template.Must(template.New("").Parse(`
	count_over_time(
		((
			1 -
			(
				{{ .AggrMetrics }}
			)
			/
			node_memory_MemTotal_bytes{{ .Filters }}
	) {{ .Operation }} bool {{ .ExpectedValue }})[{{ .SpikeWindow }}:5s]) > {{ .NumSpikes }}
`))
	var b bytes.Buffer
	err := tmpl.Execute(&b, map[string]string{
		"Filters":       filters.Build(),
		"AggrMetrics":   aggrMetrics,
		"Operation":     operation,
		"ExpectedValue": fmt.Sprintf("%.7f", expectedRatio),
		"SpikeWindow":   spikeDur.String(),
		"NumSpikes":     fmt.Sprintf("%d", numSpikes),
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
