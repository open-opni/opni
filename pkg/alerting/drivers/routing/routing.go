package routing

/*
Contains the specification of routing operations
on an AlertManager config.
*/

import (
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"

	"slices"

	"github.com/open-panoptes/opni/pkg/alerting/drivers/config"
	"github.com/open-panoptes/opni/pkg/alerting/interfaces"
	"github.com/open-panoptes/opni/pkg/alerting/shared"
	alertingv1 "github.com/open-panoptes/opni/pkg/apis/alerting/v1"
	"github.com/open-panoptes/opni/pkg/capabilities/wellknown"
	"github.com/open-panoptes/opni/pkg/util"
	"github.com/open-panoptes/opni/pkg/validation"
	amCfg "github.com/prometheus/alertmanager/config"
	"github.com/prometheus/alertmanager/pkg/labels"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/durationpb"
	"gopkg.in/yaml.v2"
)

type ProductionConfigSyncer interface {
	// Walks the tree of routes in the config, calling the given function
	Walk(map[string]string, func(depth int, r *config.Route) error) error
	// Returns the routes that match the given labels
	Search(labels map[string]string) []*config.Route
	// Merges two OpniRouting objects (also includes merging plain AlertManager configs for users)
	Merge(other OpniRouting) (OpniRouting, error)
	// Converts a valid AlertManager config to OpniRouting
	// Returns an FailedPrecondition error if the config cannot be unmarshalled,
	// Returns an InternalServerError if the config is invalid
	SyncExternalConfig(content []byte) error
}

type RoutingIdentifer interface {
	HasLabels(routingId string) []*labels.Matcher
	HasReceivers(routingId string) []string
	// SetDefaultReceiver(endpoint url.URL)
	SetDefaultReceiver(config.WebhookConfig)
}

// OpniRouting Responsible for handling the mapping of ids
// to configured endpoints, including indexing external configs
type OpniRouting interface {
	ProductionConfigSyncer
	RoutingIdentifer

	SetDefaultNamespaceConfig(endps []*alertingv1.AlertEndpoint) error
	SetNamespaceSpec(namespace string, routeId string, specs *alertingv1.FullAttachedEndpoints) error
	// When an already attached endpoint is updated, propagate updates to the routing tree
	UpdateEndpoint(id string, spec *alertingv1.AlertEndpoint) error
	// When an already attached endpoint is delete, propagate all deletions to the routing tree
	DeleteEndpoint(endpointId string) error

	// Builders

	// Converts OpniRouting to a valid AlertManager config
	// Returns a NotFound error if the a route to update or delete is not found
	// Returns a Conflict error if we try to insert a duplicate config, unique up to its keys
	BuildConfig() (*config.Config, error)
	Clone() OpniRouting
}

func NewDefaultOpniRouting() OpniRouting {
	url := util.Must(url.Parse(fmt.Sprintf("http://localhost:3000%s", shared.AlertingDefaultHookName)))
	return NewOpniRouterV1(config.WebhookConfig{
		NotifierConfig: config.NotifierConfig{
			VSendResolved: false,
		},
		URL: &amCfg.URL{
			URL: url,
		},
	})
}

var _ interfaces.Cloneable[OpniRouting] = (OpniRouting)(nil)

var _ OpniRouting = (*OpniRouterV1)(nil)

type namespacedSpecs map[string]map[string]map[string]config.OpniReceiver

func (n *namespacedSpecs) UnmarshalYAML(unmarshal func(interface{}) error) error {
	n = &namespacedSpecs{}
	out := map[string]map[string]map[string]interface{}{}
	if err := unmarshal(&out); err != nil {
		return err
	}
	for namespace, routes := range out {
		for routeId, endpoints := range routes {
			for endpointId, spec := range endpoints {
				opniRecv, err := config.ExtractReceiver(unmarshal, spec)
				if err != nil {
					return err
				}
				if _, ok := (*n)[namespace]; !ok {
					(*n)[namespace] = map[string]map[string]config.OpniReceiver{}
				}
				if _, ok := (*n)[namespace][routeId]; !ok {
					(*n)[namespace][routeId] = map[string]config.OpniReceiver{}
				}
				(*n)[namespace][routeId][endpointId] = opniRecv
			}
		}
	}
	return nil
}

type defaultNamespaceConfigs map[string]map[string]config.OpniReceiver

func (n *defaultNamespaceConfigs) UnmarshalYAML(unmarshal func(interface{}) error) error {
	n = &defaultNamespaceConfigs{}
	out := map[string]map[string]interface{}{}
	if err := unmarshal(&out); err != nil {
		return err
	}
	for defaultNamespaceValue, endpoints := range out {
		for endpointId, spec := range endpoints {
			opniRecv, err := config.ExtractReceiver(unmarshal, spec)
			if err != nil {
				return err
			}
			if _, ok := (*n)[defaultNamespaceValue]; !ok {
				(*n)[defaultNamespaceValue] = map[string]config.OpniReceiver{}
			}
			(*n)[defaultNamespaceValue][endpointId] = opniRecv
		}
	}
	return nil
}

type namespaceRateLimiting map[string]map[string]rateLimitingConfig

type rateLimitingConfig struct {
	InitialDelay       time.Duration `yaml:"initialDelay,omitempty" json:"initialDelay,omitempty"`
	RepeatInterval     time.Duration `yaml:"repeatInterval,omitempty" json:"repeatInterval,omitempty"`
	ThrottlingDuration time.Duration `yaml:"throttlingDuration,omitempty" json:"throttlingDuration,omitempty"`
}

// indexes using endpointId for scalability
type OpniRouterV1 struct {
	mu              sync.Mutex
	DefaultReceiver config.WebhookConfig `yaml:"defaultReceiver,omitempty" json:"hookEndpoint,omitempty"`
	// Contains an AlertManager config not created and managed by Opni
	SyncedConfig *config.Config `yaml:"embeddedConfig,omitempty" json:"embeddedConfig,omitempty"`

	// defaultNamespaceValue -> endpointId -> OpniConfig
	DefaultNamespaceConfigs defaultNamespaceConfigs `yaml:"defaultNamespaceConfigs,omitempty" json:"defaultNamespaceConfigs,omitempty"`
	// namespace -> routeId -> endpointId -> OpniConfig
	NamespacedSpecs namespacedSpecs `yaml:"namespacedSpecs,omitempty" json:"namespacedSpecs,omitempty"`
	// namespace -> routeId -> 	rateLimitingConfig
	NamespacedRateLimiting namespaceRateLimiting `yaml:"namespacedRateLimiting,omitempty" json:"namespacedRateLimiting,omitempty"`
}

func NewOpniRouterV1(defaultRevc config.WebhookConfig) *OpniRouterV1 {
	return &OpniRouterV1{
		// am empty config.Config is invalid in many ways, so it is easier to mark no config as nil
		SyncedConfig:            nil,
		DefaultNamespaceConfigs: make(map[string]map[string]config.OpniReceiver),
		NamespacedSpecs:         make(map[string]map[string]map[string]config.OpniReceiver),
		NamespacedRateLimiting:  make(map[string]map[string]rateLimitingConfig),
		DefaultReceiver:         defaultRevc,
	}
}

func newReceiverImplementationFromEndpoint(endp *alertingv1.AlertEndpoint, details *alertingv1.EndpointImplementation) config.OpniReceiver {
	var newConfig config.OpniReceiver
	switch endp.GetEndpoint().(type) {
	case *alertingv1.AlertEndpoint_Email:
		newConfig = (&config.EmailConfig{}).Configure(endp)
		newConfig.StoreInfo(details)
	case *alertingv1.AlertEndpoint_Slack:
		newConfig = (&config.SlackConfig{}).Configure(endp)
		newConfig.StoreInfo(details)
	case *alertingv1.AlertEndpoint_PagerDuty:
		newConfig = (&config.PagerdutyConfig{}).Configure(endp)
		newConfig.StoreInfo(details)
	case *alertingv1.AlertEndpoint_Webhook:
		newConfig = (&config.WebhookConfig{}).Configure(endp)
		newConfig.StoreInfo(details)
	default:
		strRepr, _ := protojson.Marshal(endp)
		panic(fmt.Sprintf("no such endpoint type implemented %s", strRepr))
	}
	if newConfig == nil {
		panic("new config should always be non-nil")
	}
	return newConfig
}

func (o *OpniRouterV1) HasLabels(routingId string) []*labels.Matcher {
	for namespaceName, routes := range o.NamespacedSpecs {
		if _, ok := routes[routingId]; ok {
			return []*labels.Matcher{
				{
					Type:  labels.MatchEqual,
					Name:  namespaceName,
					Value: routingId,
				},
			}
		}
	}
	return nil
}

func (o *OpniRouterV1) HasReceivers(routingId string) []string {
	for namespaceName, routes := range o.NamespacedSpecs {
		if _, ok := routes[routingId]; ok {
			return []string{
				shared.NewOpniReceiverName(shared.OpniReceiverId{
					Namespace:  namespaceName,
					ReceiverId: routingId,
				}),
			}
		}
	}
	return []string{}
}

func (o *OpniRouterV1) SetDefaultReceiver(cfg config.WebhookConfig) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.DefaultReceiver = cfg
}

func (o *OpniRouterV1) SyncExternalConfig(content []byte) error {
	// the default alertmanager validation is embedded into the implementation of yaml.Unmarshallable
	var cfg *config.Config
	if err := yaml.Unmarshal(content, &cfg); err != nil {
		return status.Error(codes.FailedPrecondition, fmt.Sprintf("Alertmanager configuration not supported %s", err))
	}
	o.SyncedConfig = cfg
	return nil
}

func (o *OpniRouterV1) SetDefaultNamespaceConfig(endpoints []*alertingv1.AlertEndpoint) error {
	for _, val := range NotificationSubTreeValues() {
		if len(endpoints) == 0 { // delete
			delete(o.DefaultNamespaceConfigs, val.A)
			return nil
		}
		validKeys := map[string]struct{}{}
		for _, endpoint := range endpoints {
			if err := endpoint.Validate(); err != nil {
				return err
			}
			validKeys[endpoint.GetId()] = struct{}{}
		}

		details := &alertingv1.EndpointImplementation{
			Title: fmt.Sprintf("{{ .%s }}", shared.OpniTitleLabel),
			Body:  fmt.Sprintf("{{ .%s }}", shared.OpniBodyLabel),
		}
		o.DefaultNamespaceConfigs[val.A] = map[string]config.OpniReceiver{}
		for _, spec := range endpoints {
			o.DefaultNamespaceConfigs[val.A][spec.Id] = newReceiverImplementationFromEndpoint(spec, details)
		}
	}
	return nil
}

func (o *OpniRouterV1) SetNamespaceSpec(namespace, routeId string, specs *alertingv1.FullAttachedEndpoints) error {
	o.mu.Lock()
	defer o.mu.Unlock()
	if namespace == "" {
		return validation.Error("namespace cannot be empty when setting specs")
	}
	if namespace == NotificationSubTreeLabel() {
		return validation.Error("namespace cannot be the default namespace label")
	}
	// set receiver specs
	for _, spec := range specs.GetItems() {
		if err := spec.Validate(); err != nil {
			return status.Error(codes.InvalidArgument, fmt.Sprintf("failed to update endpoint with route id %s: %s", routeId, err))
		}
	}
	if _, ok := o.NamespacedSpecs[namespace]; !ok {
		o.NamespacedSpecs[namespace] = make(map[string]map[string]config.OpniReceiver)
	}
	if _, ok := o.NamespacedSpecs[namespace][routeId]; !ok {
		o.NamespacedSpecs[namespace][routeId] = make(map[string]config.OpniReceiver)
	}
	o.NamespacedSpecs[namespace][routeId] = make(map[string]config.OpniReceiver)
	for _, spec := range specs.GetItems() {
		o.NamespacedSpecs[namespace][routeId][spec.EndpointId] = newReceiverImplementationFromEndpoint(spec.GetAlertEndpoint(), specs.GetDetails())
	}

	// set rate limiting specs
	if _, ok := o.NamespacedRateLimiting[namespace]; !ok {
		o.NamespacedRateLimiting[namespace] = make(map[string]rateLimitingConfig)
	}
	o.NamespacedRateLimiting[namespace][routeId] = rateLimitingConfig{
		InitialDelay:       lo.ToPtr(lo.FromPtrOr(specs.GetInitialDelay(), *durationpb.New(time.Minute))).AsDuration(),
		RepeatInterval:     lo.ToPtr(lo.FromPtrOr(specs.GetRepeatInterval(), *durationpb.New(10 * time.Minute))).AsDuration(),
		ThrottlingDuration: lo.ToPtr(lo.FromPtrOr(specs.GetThrottlingDuration(), *durationpb.New(time.Second * 30))).AsDuration(),
	}
	return nil
}

func (o *OpniRouterV1) UpdateEndpoint(id string, spec *alertingv1.AlertEndpoint) error {
	if err := spec.Validate(); err != nil {
		return validation.Errorf("invalid endpoint : %s", err)
	}
	for _ /* namespace */, route := range o.NamespacedSpecs {
		for _ /* routeId */, endpoint := range route {
			if _, ok := endpoint[id]; ok {
				details := endpoint[id].ExtractInfo()
				endpoint[id] = newReceiverImplementationFromEndpoint(spec, details)
			}
		}
	}

	for _ /*defaultValue*/, endpoints := range o.DefaultNamespaceConfigs {
		if _, ok := endpoints[id]; ok {
			details := endpoints[id].ExtractInfo()
			endpoints[id] = newReceiverImplementationFromEndpoint(spec, details)
		}
	}
	return nil
}

func (o *OpniRouterV1) DeleteEndpoint(id string) error {
	for _, route := range o.NamespacedSpecs {
		for _, endpoint := range route {
			delete(endpoint, id)
		}
	}

	for value, endpoints := range o.DefaultNamespaceConfigs {
		delete(endpoints, id)
		if len(endpoints) == 0 {
			delete(o.DefaultNamespaceConfigs, value)
		}
		if len(value) == 0 { // clean up empty default namespace values
			delete(o.DefaultNamespaceConfigs, value)
		}
	}
	return nil
}

func (o *OpniRouterV1) BuildConfig() (*config.Config, error) {
	o.mu.Lock()
	defer o.mu.Unlock()
	root := NewRoutingTree(&o.DefaultReceiver)

	// update the default namespace with the configs
	for i, recv := range root.Receivers {
		if recv.Name == shared.AlertingHookReceiverName { // ingore the default hook with doesn't abide by a (namespace, receiverId) naming convention
			continue
		}

		opniReceiverId, err := shared.ExtractReceiverId(recv.Name)
		if err != nil {
			panic(err)
		}
		recvName := opniReceiverId.ReceiverId

		if _, ok := o.DefaultNamespaceConfigs[recvName]; ok {
			endpIds := lo.Keys(o.DefaultNamespaceConfigs[recvName])
			slices.SortFunc(endpIds, strings.Compare)
			opniReceivers := make([]config.OpniReceiver, len(endpIds))
			for i, endpId := range endpIds {
				opniReceivers[i] = o.DefaultNamespaceConfigs[recvName][endpId]
			}
			recv, err := config.BuildReceiver(shared.NewOpniReceiverName(shared.OpniReceiverId{
				Namespace:  opniReceiverId.Namespace,
				ReceiverId: recvName,
			}), opniReceivers)
			if err != nil {
				panic(fmt.Sprintf("name : %s : %s", recvName, err))
			}
			root.Receivers[i] = recv
		}
	}

	// build each namespaced tree that isn't the default namespace
	opniRoutes := []*config.Route{}
	opniReceivers := []config.Receiver{}
	namespaces := lo.Keys(o.NamespacedSpecs) // needs to be deterministically ordered
	slices.SortFunc(namespaces, strings.Compare)
	for _, namespace := range namespaces {
		routeIds := lo.Keys(o.NamespacedSpecs[namespace]) // needs to be deterministically ordered
		slices.SortFunc(routeIds, strings.Compare)
		namespacedSubTree, _ := NewNamespaceTree(namespace)
		for _, routeId := range routeIds {
			if len(o.NamespacedSpecs[namespace][routeId]) == 0 {
				// no opni receivers attached, do not build & skip...
				continue
			}
			endpointIds := lo.Keys(o.NamespacedSpecs[namespace][routeId]) // needs to be deterministically ordered
			slices.SortFunc(endpointIds, strings.Compare)
			endpoints := make([]config.OpniReceiver, len(endpointIds))
			for i, endpointId := range endpointIds {
				endpoints[i] = o.NamespacedSpecs[namespace][routeId][endpointId]
			}
			namespacedValueSubTree, namespacedReceivers := NewNamespaceLeaf(
				o.NamespacedRateLimiting[namespace][routeId],
				endpoints,
				o.HasLabels(routeId),
				o.HasReceivers(routeId)[0],
			)
			// prepend
			namespacedSubTree.Routes = append([]*config.Route{namespacedValueSubTree}, namespacedSubTree.Routes...)
			opniReceivers = append(opniReceivers, namespacedReceivers)
		}
		opniRoutes = append(opniRoutes, namespacedSubTree)
	}

	// add opni subtree dependencies (opni namespaced & metrics)
	for _, subRoute := range root.Route.Routes {
		for _, m := range subRoute.Matchers {
			if m.Name == alertingv1.RoutingPropertyDatasource && m.Type == labels.MatchEqual && m.Value == "" { // if isDefaultSubTree() {}
				// prepend
				subRoute.Routes = append(opniRoutes, subRoute.Routes...)
			}

			// production configs get added here, to the metrics subtree
			if m.Name == alertingv1.RoutingPropertyDatasource && m.Type == labels.MatchEqual && m.Value == wellknown.CapabilityMetrics {
				if o.SyncedConfig != nil {
					// add the entire tree to the subroute
					subRoute.Routes = []*config.Route{o.SyncedConfig.Route}
					root.Global = o.SyncedConfig.Global
					root.InhibitRules = o.SyncedConfig.InhibitRules
					root.TimeIntervals = o.SyncedConfig.TimeIntervals
					//FIXME: we *may* eventually need to allow some way to import template files
					root.Templates = o.SyncedConfig.Templates
					root.InhibitRules = o.SyncedConfig.InhibitRules
					root.MuteTimeIntervals = append(root.MuteTimeIntervals, o.SyncedConfig.MuteTimeIntervals...)
					// prepend
					root.Receivers = append(o.SyncedConfig.Receivers, root.Receivers...)
				}
			}
		}
	}
	slices.SortFunc(opniReceivers, func(a, b config.Receiver) int {
		return strings.Compare(a.Name, b.Name)
	})
	//prepend
	root.Receivers = append(opniReceivers, root.Receivers...)
	if root.Receivers[len(root.Receivers)-1].Name != shared.AlertingHookReceiverName {
		panic("default receiver should always be last")
	}
	return root, nil
}

func (o *OpniRouterV1) Clone() OpniRouting {
	oCopy := NewOpniRouterV1(o.DefaultReceiver)
	if o.SyncedConfig != nil {
		oCopy.SyncedConfig = util.DeepCopy(o.SyncedConfig)
	}

	// the internal maps are not compatible with deepcopy, since interfaces don't support new() builtin
	oCopy.DefaultNamespaceConfigs = map[string]map[string]config.OpniReceiver{}
	oCopy.NamespacedSpecs = map[string]map[string]map[string]config.OpniReceiver{}
	oCopy.NamespacedRateLimiting = map[string]map[string]rateLimitingConfig{}

	for namespace, namespaceSpecs := range o.NamespacedSpecs {
		oCopy.NamespacedSpecs[namespace] = map[string]map[string]config.OpniReceiver{}
		oCopy.NamespacedRateLimiting[namespace] = map[string]rateLimitingConfig{}
		for routeId, routeSpecs := range namespaceSpecs {
			oCopy.NamespacedSpecs[namespace][routeId] = map[string]config.OpniReceiver{}
			for receiverName, receiver := range routeSpecs {
				oCopy.NamespacedSpecs[namespace][routeId][receiverName] = receiver.Clone()
			}
			oCopy.NamespacedRateLimiting[namespace][routeId] = o.NamespacedRateLimiting[namespace][routeId]
		}
	}

	for namespace, namespaceSpecs := range o.DefaultNamespaceConfigs {
		oCopy.DefaultNamespaceConfigs[namespace] = map[string]config.OpniReceiver{}
		for receiverName, receiver := range namespaceSpecs {
			oCopy.DefaultNamespaceConfigs[namespace][receiverName] = receiver.Clone()
		}
	}

	return oCopy
}

func (o *OpniRouterV1) Walk(map[string]string, func(int, *config.Route) error) error {
	return status.Error(codes.Unimplemented, "OpniRouterV1 does not implement Walk")
}

func (o *OpniRouterV1) Search(map[string]string) []*config.Route {
	return []*config.Route{}
}

func (o *OpniRouterV1) Merge(_ OpniRouting) (OpniRouting, error) {
	return nil, status.Error(codes.Unimplemented, "OpniRouterV1 does not implement Merge")
}
