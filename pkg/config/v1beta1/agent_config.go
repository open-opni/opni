package v1beta1

import (
	"github.com/open-panoptes/opni/pkg/config/meta"
	"github.com/open-panoptes/opni/pkg/tokens"
)

type AgentConfig struct {
	meta.TypeMeta `json:",inline"`

	Spec AgentConfigSpec `json:"spec,omitempty"`
}

type TrustStrategyKind string

const (
	TrustStrategyPKP      TrustStrategyKind = "pkp"
	TrustStrategyCACerts  TrustStrategyKind = "cacerts"
	TrustStrategyInsecure TrustStrategyKind = "insecure"
)

type AgentConfigSpec struct {
	// The address which the agent will listen on for incoming connections.
	// This should be in the format "host:port" or ":port", and must not
	// include a scheme.
	ListenAddress string `json:"listenAddress,omitempty"`
	// The address of the gateway's public GRPC API. This should be of the format
	// "host:port" with no scheme.
	GatewayAddress string `json:"gatewayAddress,omitempty"`
	// The name of the identity provider to use. Defaults to "kubernetes".
	IdentityProvider string `json:"identityProvider,omitempty"`
	// The type of trust strategy to use for verifying the authenticity of the
	// gateway server. Defaults to "pkp".
	TrustStrategy TrustStrategyKind `json:"trustStrategy,omitempty"`
	// Configuration for agent keyring storage.
	Storage       StorageSpec       `json:"storage,omitempty"`
	Rules         *RulesSpec        `json:"rules,omitempty"`
	Bootstrap     *BootstrapSpec    `json:"bootstrap,omitempty"`
	LogLevel      string            `json:"logLevel,omitempty"`
	PluginDir     string            `json:"pluginDir,omitempty"`
	Keyring       KeyringSpec       `json:"keyring,omitempty"`
	Upgrade       AgentUpgradeSpec  `json:"upgrade,omitempty"`
	PluginUpgrade PluginUpgradeSpec `json:"pluginUpgrade,omitempty"`
}

type BootstrapSpec struct {
	// Address of the internal management GRPC API. Used for auto-bootstrapping
	// when direct management api access is available, such as when running in
	// the main cluster.
	InClusterManagementAddress *string `json:"inClusterManagementAddress,omitempty"`

	// An optional display name to assign to the cluster when creating it.
	// This value corresponds to the label `opni.io/name`, and can be modified
	// at any time after the cluster is created.
	FriendlyName *string `json:"friendlyName,omitempty"`

	// Bootstrap token
	Token string `json:"token,omitempty"`
	// List of public key pins. Used when the trust strategy is "pkp".
	Pins []string `json:"pins,omitempty"`
	// List of paths to CA Certs. Used when the trust strategy is "cacerts".
	// If empty, the system certs will be used.
	CACerts []string `json:"caCerts,omitempty"`
}

type AgentUpgradeType string

const (
	AgentUpgradeNoop       AgentUpgradeType = "noop"
	AgentUpgradeKubernetes AgentUpgradeType = "kubernetes"
)

type PluginUpgradeType string

const (
	PluginUpgradeNoop   PluginUpgradeType = "noop"
	PluginUpgradeBinary PluginUpgradeType = "binary"
)

type AgentUpgradeSpec struct {
	Type       AgentUpgradeType       `json:"type,omitempty"`
	Kubernetes *KubernetesUpgradeSpec `json:"kubernetes,omitempty"`
}

type PluginUpgradeSpec struct {
	Type   PluginUpgradeType  `json:"type,omitempty"`
	Binary *BinaryUpgradeSpec `json:"binary,omitempty"`
}

type KubernetesUpgradeSpec struct {
	Namespace    string  `json:"namespace,omitempty"`
	RepoOverride *string `json:"repoOverride,omitempty"`
}

type BinaryUpgradeSpec struct{}

func (s *AgentConfigSpec) ContainsBootstrapCredentials() bool {
	if s.Bootstrap == nil {
		return false
	}
	if s.Bootstrap.InClusterManagementAddress != nil {
		return s.Bootstrap.Token == "" &&
			len(s.Bootstrap.Pins) == 0 &&
			len(s.Bootstrap.CACerts) == 0
	}

	_, err := tokens.ParseHex(s.Bootstrap.Token)
	if err != nil {
		return false
	}
	switch s.TrustStrategy {
	case TrustStrategyPKP:
		return len(s.Bootstrap.Pins) > 0
	case TrustStrategyCACerts:
		return len(s.Bootstrap.CACerts) > 0
	}
	return false
}

func (s *AgentConfigSpec) SetDefaults() {
	if s == nil {
		return
	}
	if s.IdentityProvider == "" {
		s.IdentityProvider = "kubernetes"
	}
	if s.ListenAddress == "" {
		s.ListenAddress = ":8080"
	}
	if s.TrustStrategy == "" {
		s.TrustStrategy = "pkp"
	}
}
