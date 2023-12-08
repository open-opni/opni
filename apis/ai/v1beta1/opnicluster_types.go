/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// +kubebuilder:validation:Optional
package v1beta1

import (
	opnimeta "github.com/open-panoptes/opni/pkg/util/meta"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type OpniClusterState string

const (
	OpniClusterStateError   OpniClusterState = "Error"
	OpniClusterStateWorking OpniClusterState = "Working"
	OpniClusterStateReady   OpniClusterState = "Ready"
)

// OpniClusterSpec defines the desired state of OpniCluster
type OpniClusterSpec struct {
	// +kubebuilder:default:=latest
	Version string `json:"version"`
	// +optional
	DefaultRepo          *string                        `json:"defaultRepo,omitempty"`
	NatsRef              corev1.LocalObjectReference    `json:"natsCluster"`
	Services             ServicesSpec                   `json:"services,omitempty"`
	Opensearch           *opnimeta.OpensearchClusterRef `json:"opensearch,omitempty"`
	S3                   S3Spec                         `json:"s3,omitempty"`
	NulogHyperparameters map[string]intstr.IntOrString  `json:"nulogHyperparameters,omitempty"`
	DeployLogCollector   *bool                          `json:"deployLogCollector"`
	GlobalNodeSelector   map[string]string              `json:"globalNodeSelector,omitempty"`
	GlobalTolerations    []corev1.Toleration            `json:"globalTolerations,omitempty"`
}

// OpniClusterStatus defines the observed state of OpniCluster
type OpniClusterStatus struct {
	Conditions              []string         `json:"conditions,omitempty"`
	State                   OpniClusterState `json:"state,omitempty"`
	LogCollectorState       OpniClusterState `json:"logState,omitempty"`
	Auth                    AuthStatus       `json:"auth,omitempty"`
	PrometheusRuleNamespace string           `json:"prometheusRuleNamespace,omitempty"`
	IndexState              OpniClusterState `json:"indexState,omitempty"`
}

type AuthStatus struct {
	OpensearchAuthSecretKeyRef *corev1.SecretKeySelector `json:"opensearchAuthSecretKeyRef,omitempty"`
	S3Endpoint                 string                    `json:"s3Endpoint,omitempty"`
	S3AccessKey                *corev1.SecretKeySelector `json:"s3AccessKey,omitempty"`
	S3SecretKey                *corev1.SecretKeySelector `json:"s3SecretKey,omitempty"`
}

type OpensearchStatus struct {
	IndexState  OpniClusterState `json:"indexState,omitempty"`
	Version     *string          `json:"version,omitempty"`
	Initialized bool             `json:"initialized,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// +kubebuilder:printcolumn:name="State",type=boolean,JSONPath=`.status.state`
// +kubebuilder:storageversion

// OpniCluster is the Schema for the opniclusters API
type OpniCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpniClusterSpec   `json:"spec,omitempty"`
	Status OpniClusterStatus `json:"status,omitempty"`
}

type ServicesSpec struct {
	Drain              DrainServiceSpec              `json:"drain,omitempty"`
	Inference          InferenceServiceSpec          `json:"inference,omitempty"`
	Preprocessing      PreprocessingServiceSpec      `json:"preprocessing,omitempty"`
	PayloadReceiver    PayloadReceiverServiceSpec    `json:"payloadReceiver,omitempty"`
	GPUController      GPUControllerServiceSpec      `json:"gpuController,omitempty"`
	Metrics            MetricsServiceSpec            `json:"metrics,omitempty"`
	OpensearchUpdate   OpensearchUpdateServiceSpec   `json:"opensearchUpdate,omitempty"`
	TrainingController TrainingControllerServiceSpec `json:"trainingController,omitempty"`
}

type DrainServiceSpec struct {
	opnimeta.ImageSpec `json:",inline,omitempty"`
	Enabled            *bool                    `json:"enabled,omitempty"`
	NodeSelector       map[string]string        `json:"nodeSelector,omitempty"`
	Tolerations        []corev1.Toleration      `json:"tolerations,omitempty"`
	Replicas           *int32                   `json:"replicas,omitempty"`
	Workload           WorkloadDrainServiceSpec `json:"workload,omitempty"`
}

type WorkloadDrainServiceSpec struct {
	Enabled  *bool  `json:"enabled,omitempty"`
	Replicas *int32 `json:"replicas,omitempty"`
}

type InferenceServiceSpec struct {
	opnimeta.ImageSpec `json:",inline,omitempty"`
	Enabled            *bool                         `json:"enabled,omitempty"`
	PretrainedModels   []corev1.LocalObjectReference `json:"pretrainedModels,omitempty"`
	NodeSelector       map[string]string             `json:"nodeSelector,omitempty"`
	Tolerations        []corev1.Toleration           `json:"tolerations,omitempty"`
}

type PreprocessingServiceSpec struct {
	opnimeta.ImageSpec `json:",inline,omitempty"`
	Enabled            *bool               `json:"enabled,omitempty"`
	NodeSelector       map[string]string   `json:"nodeSelector,omitempty"`
	Tolerations        []corev1.Toleration `json:"tolerations,omitempty"`
	Replicas           *int32              `json:"replicas,omitempty"`
}

type PayloadReceiverServiceSpec struct {
	opnimeta.ImageSpec `json:",inline,omitempty"`
	Enabled            *bool               `json:"enabled,omitempty"`
	NodeSelector       map[string]string   `json:"nodeSelector,omitempty"`
	Tolerations        []corev1.Toleration `json:"tolerations,omitempty"`
}

type GPUControllerServiceSpec struct {
	opnimeta.ImageSpec `json:",inline,omitempty"`
	Enabled            *bool               `json:"enabled,omitempty"`
	RuntimeClass       *string             `json:"runtimeClass,omitempty"`
	NodeSelector       map[string]string   `json:"nodeSelector,omitempty"`
	Tolerations        []corev1.Toleration `json:"tolerations,omitempty"`
}

type TrainingControllerServiceSpec struct {
	opnimeta.ImageSpec `json:",inline,omitempty"`
	Enabled            *bool               `json:"enabled,omitempty"`
	NodeSelector       map[string]string   `json:"nodeSelector,omitempty"`
	Tolerations        []corev1.Toleration `json:"tolerations,omitempty"`
}

type MetricsServiceSpec struct {
	opnimeta.ImageSpec  `json:",inline,omitempty"`
	Enabled             *bool                         `json:"enabled,omitempty"`
	NodeSelector        map[string]string             `json:"nodeSelector,omitempty"`
	Tolerations         []corev1.Toleration           `json:"tolerations,omitempty"`
	ExtraVolumeMounts   []opnimeta.ExtraVolumeMount   `json:"extraVolumeMounts,omitempty"`
	PrometheusEndpoint  string                        `json:"prometheusEndpoint,omitempty"`
	PrometheusReference *opnimeta.PrometheusReference `json:"prometheus,omitempty"`
}

type InsightsServiceSpec struct {
	opnimeta.ImageSpec `json:",inline,omitempty"`
	Enabled            *bool               `json:"enabled,omitempty"`
	NodeSelector       map[string]string   `json:"nodeSelector,omitempty"`
	Tolerations        []corev1.Toleration `json:"tolerations,omitempty"`
}

type UIServiceSpec struct {
	opnimeta.ImageSpec `json:",inline,omitempty"`
	Enabled            *bool               `json:"enabled,omitempty"`
	NodeSelector       map[string]string   `json:"nodeSelector,omitempty"`
	Tolerations        []corev1.Toleration `json:"tolerations,omitempty"`
}

type OpensearchUpdateServiceSpec struct {
	opnimeta.ImageSpec `json:",inline,omitempty"`
	Enabled            *bool               `json:"enabled,omitempty"`
	NodeSelector       map[string]string   `json:"nodeSelector,omitempty"`
	Tolerations        []corev1.Toleration `json:"tolerations,omitempty"`
}

type S3Spec struct {
	// If set, Opni will deploy an S3 pod to use internally.
	// Cannot be set at the same time as `external`.
	Internal *InternalSpec `json:"internal,omitempty"`
	// If set, Opni will connect to an external S3 endpoint.
	// Cannot be set at the same time as `internal`.
	External *ExternalSpec `json:"external,omitempty"`
	// Bucket used to persist nulog models.  If not set will use
	// opni-nulog-models.
	NulogS3Bucket string `json:"nulogS3Bucket,omitempty"`
	// Bucket used to persiste drain models.  It not set will use
	// opni-drain-models
	DrainS3Bucket string `json:"drainS3Bucket,omitempty"`
}

type InternalSpec struct {
	// Persistence configuration for internal S3 deployment. If unset, internal
	// S3 storage is not persistent.
	Persistence *opnimeta.PersistenceSpec `json:"persistence,omitempty"`
}

type ExternalSpec struct {
	// +kubebuilder:validation:Required
	// External S3 endpoint URL.
	Endpoint string `json:"endpoint,omitempty"`
	// +kubebuilder:validation:Required
	// Reference to a secret containing "accessKey" and "secretKey" items. This
	// secret must already exist if specified.
	Credentials *corev1.SecretReference `json:"credentials,omitempty"`
}

// +kubebuilder:object:root=true

// OpniClusterList contains a list of OpniCluster
type OpniClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpniCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OpniCluster{}, &OpniClusterList{})
}
