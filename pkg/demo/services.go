package demo

import (
	"fmt"

	demov1alpha1 "github.com/rancher/opni/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	extv1beta1 "k8s.io/api/extensions/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func BuildDrainService(spec *demov1alpha1.OpniDemo) *appsv1.Deployment {
	labels := map[string]string{"app": "drain-service"}
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "drain-service",
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:            "drain-service",
							Image:           "rancher/opni-drain-service:v0.1.1",
							ImagePullPolicy: v1.PullAlways,
							Env: []v1.EnvVar{
								{
									Name:  "NATS_SERVER_URL",
									Value: fmt.Sprintf("nats://nats_client:%s@nats-client.%s.svc:4222", spec.Spec.NatsPassword, spec.Namespace),
								},
								{
									Name:  "MINIO_SERVER_URL",
									Value: fmt.Sprintf("http://minio.%s.svc.cluster.local:9000", spec.Namespace),
								},
								{
									Name:  "MINIO_ACCESS_KEY",
									Value: spec.Spec.MinioAccessKey,
								},
								{
									Name:  "MINIO_SECRET_KEY",
									Value: spec.Spec.MinioSecretKey,
								},
								{
									Name:  "ES_ENDPOINT",
									Value: fmt.Sprintf("https://opendistro-es-client-service.%s.svc.cluster.local:9200", spec.Namespace),
								},
								{
									Name:  "FAIL_KEYWORDS",
									Value: "fail,error,missing,unable",
								},
								{
									Name:  "ES_USERNAME",
									Value: spec.Spec.ElasticsearchUser,
								},
								{
									Name:  "ES_PASSWORD",
									Value: spec.Spec.ElasticsearchPassword,
								},
							},
						},
					},
				},
			},
		},
	}
}

func BuildNulogInferenceServiceControlPlane(spec *demov1alpha1.OpniDemo) *appsv1.Deployment {
	labels := map[string]string{"app": "nulog-inference-service-control-plane"}
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "nulog-inference-service-control-plane",
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:            "nulog-inference-service-control-plane",
							Image:           "rancher/opni-inference-service:v0.1.1",
							ImagePullPolicy: v1.PullAlways,
							Env: []v1.EnvVar{
								{
									Name:  "NATS_SERVER_URL",
									Value: fmt.Sprintf("nats://nats_client:%s@nats-client.%s.svc:4222", spec.Spec.NatsPassword, spec.Namespace),
								},
								{
									Name:  "MINIO_ENDPOINT",
									Value: fmt.Sprintf("http://minio.%s.svc.cluster.local:9000", spec.Namespace),
								},
								{
									Name:  "MINIO_ACCESS_KEY",
									Value: spec.Spec.MinioAccessKey,
								},
								{
									Name:  "MINIO_SECRET_KEY",
									Value: spec.Spec.MinioSecretKey,
								},
								{
									Name:  "ES_ENDPOINT",
									Value: fmt.Sprintf("https://opendistro-es-client-service.%s.svc.cluster.local:9200", spec.Namespace),
								},
								{
									Name:  "MODEL_THRESHOLD",
									Value: "0.6",
								},
								{
									Name:  "MIN_LOG_TOKENS",
									Value: "4",
								},
								{
									Name:  "IS_CONTROL_PLANE_SERVICE",
									Value: "True",
								},
							},
							Resources: v1.ResourceRequirements{
								Requests: v1.ResourceList{
									v1.ResourceMemory: resource.MustParse("1Gi"),
									v1.ResourceCPU:    resource.MustParse(spec.Spec.NulogServiceCpuRequest),
								},
							},
						},
					},
				},
			},
		},
	}
}

func BuildNulogInferenceService(spec *demov1alpha1.OpniDemo) *appsv1.Deployment {
	labels := map[string]string{"app": "nulog-inference-service"}
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "nulog-inference-service",
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:            "nulog-inference-service",
							Image:           "rancher/opni-inference-service:v0.1.1",
							ImagePullPolicy: v1.PullAlways,
							Env: []v1.EnvVar{
								{
									Name:  "NATS_SERVER_URL",
									Value: fmt.Sprintf("nats://nats_client:%s@nats-client.%s.svc:4222", spec.Spec.NatsPassword, spec.Namespace),
								},
								{
									Name:  "MINIO_ENDPOINT",
									Value: fmt.Sprintf("http://minio.%s.svc.cluster.local:9000", spec.Namespace),
								},
								{
									Name:  "MINIO_ACCESS_KEY",
									Value: spec.Spec.MinioAccessKey,
								},
								{
									Name:  "MINIO_SECRET_KEY",
									Value: spec.Spec.MinioSecretKey,
								},
								{
									Name:  "ES_ENDPOINT",
									Value: fmt.Sprintf("https://opendistro-es-client-service.%s.svc.cluster.local:9200", spec.Namespace),
								},
								{
									Name:  "MODEL_THRESHOLD",
									Value: "0.5",
								},
								{
									Name:  "MIN_LOG_TOKENS",
									Value: "5",
								},
							},
							Resources: v1.ResourceRequirements{
								Limits: v1.ResourceList{
									"nvidia.com/gpu": resource.MustParse("1"),
								},
							},
						},
					},
				},
			},
		},
	}
}

var falseVar = false

func BuildNvidiaPlugin(spec *demov1alpha1.OpniDemo) *appsv1.DaemonSet {
	labels := map[string]string{
		"name": "nvidia-device-plugin-ds",
	}
	return &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "nvidia-device-plugin-daemonset",
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			UpdateStrategy: appsv1.DaemonSetUpdateStrategy{
				Type: appsv1.RollingUpdateDaemonSetStrategyType,
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"scheduler.alpha.kubernetes.io/critical-pod": "",
					},
					Labels: labels,
				},
				Spec: v1.PodSpec{
					Tolerations: []v1.Toleration{
						{
							Key:      "CriticalAddonsOnly",
							Operator: v1.TolerationOpExists,
						},
						{
							Key:      "nvidia.com/gpu",
							Operator: v1.TolerationOpExists,
							Effect:   v1.TaintEffectNoSchedule,
						},
					},
					PriorityClassName: "system-node-critical",
					Containers: []v1.Container{
						{
							Name:  "nvidia-device-plugin-ctr",
							Image: fmt.Sprintf("nvidia/k8s-device-plugin:%s", spec.Spec.NvidiaVersion),
							SecurityContext: &v1.SecurityContext{
								AllowPrivilegeEscalation: &falseVar,
								Capabilities: &v1.Capabilities{
									Drop: []v1.Capability{
										v1.Capability("ALL"),
									},
								},
							},
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "device-plugin",
									MountPath: "/var/lib/kubelet/device-plugins",
								},
							},
						},
					},
					Volumes: []v1.Volume{
						{
							Name: "device-plugin",
							VolumeSource: v1.VolumeSource{
								HostPath: &v1.HostPathVolumeSource{
									Path: "/var/lib/kubelet/device-plugins",
								},
							},
						},
					},
				},
			},
		},
	}
}

func BuildPayloadReceiverService(spec *demov1alpha1.OpniDemo) (*v1.Service, *appsv1.Deployment, *extv1beta1.Ingress) {
	labels := map[string]string{
		"app": "payload-receiver-service",
	}
	return &v1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name: "payload-receiver-service",
				Labels: map[string]string{
					"service": "payload-receiver-service",
				},
			},
			Spec: v1.ServiceSpec{
				Selector: labels,
				Ports: []v1.ServicePort{
					{
						Port: 80,
					},
				},
			},
		},
		&appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name: "payload-receiver-service",
			},
			Spec: appsv1.DeploymentSpec{
				Selector: &metav1.LabelSelector{
					MatchLabels: labels,
				},
				Template: v1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: labels,
					},
					Spec: v1.PodSpec{
						Containers: []v1.Container{
							{
								Name:            "payload-receiver-service",
								Image:           "rancher/opni-payload-receiver-service:v0.1.1",
								ImagePullPolicy: v1.PullAlways,
								Env: []v1.EnvVar{
									{
										Name:  "NATS_SERVER_URL",
										Value: fmt.Sprintf("nats://nats_client:%s@nats-client.%s.svc:4222", spec.Spec.NatsPassword, spec.Namespace),
									},
								},
								Ports: []v1.ContainerPort{
									{
										ContainerPort: 80,
									},
								},
							},
						},
					},
				},
			},
		},
		&extv1beta1.Ingress{
			ObjectMeta: metav1.ObjectMeta{
				Name: "payload-receiver-service-ingress",
				Annotations: map[string]string{
					"kubernetes.io/ingress.class": "traefik",
				},
			},
			Spec: extv1beta1.IngressSpec{
				Rules: []extv1beta1.IngressRule{
					{
						IngressRuleValue: extv1beta1.IngressRuleValue{
							HTTP: &extv1beta1.HTTPIngressRuleValue{
								Paths: []extv1beta1.HTTPIngressPath{
									{
										Path: "/",
										Backend: extv1beta1.IngressBackend{
											ServiceName: "payload-receiver-service",
											ServicePort: intstr.FromInt(80),
										},
									},
								},
							},
						},
					},
				},
			},
		}
}

func BuildPreprocessingService(spec *demov1alpha1.OpniDemo) *appsv1.Deployment {
	labels := map[string]string{
		"app": "preprocessing-service",
	}
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "preprocessing-service",
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:            "preprocessing-service",
							Image:           "rancher/opni-preprocessing-service:v0.1.1",
							ImagePullPolicy: v1.PullAlways,
							Env: []v1.EnvVar{
								{
									Name:  "NATS_SERVER_URL",
									Value: fmt.Sprintf("nats://nats_client:%s@nats-client.%s.svc:4222", spec.Spec.NatsPassword, spec.Namespace),
								},
								{
									Name:  "ES_ENDPOINT",
									Value: fmt.Sprintf("https://opendistro-es-client-service.%s.svc.cluster.local:9200", spec.Namespace),
								},
								{
									Name:  "ES_USERNAME",
									Value: spec.Spec.ElasticsearchUser,
								},
								{
									Name:  "ES_PASSWORD",
									Value: spec.Spec.ElasticsearchPassword,
								},
							},
						},
					},
				},
			},
		},
	}
}

func BuildTrainingControllerInfra(spec *demov1alpha1.OpniDemo) []client.Object {
	return []client.Object{
		&corev1.ServiceAccount{
			ObjectMeta: metav1.ObjectMeta{
				Name: "training-controller-rb",
			},
		},
		&rbacv1.ClusterRole{
			ObjectMeta: metav1.ObjectMeta{
				Name: "training-controller-rb",
			},
			Rules: []rbacv1.PolicyRule{
				{
					APIGroups: []string{"", "apps", "batch"},
					Resources: []string{"endpoints", "deployments", "pods", "jobs"},
					Verbs:     []string{"get", "list", "watch", "create", "delete"},
				},
			},
		},
		&rbacv1.ClusterRoleBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name: "training-controller-rb",
			},
			Subjects: []rbacv1.Subject{
				{
					Kind:      "ServiceAccount",
					Name:      "training-controller-rb",
					Namespace: spec.Namespace,
				},
			},
			RoleRef: rbacv1.RoleRef{
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     "ClusterRole",
				Name:     "training-controller-rb",
			},
		},
	}
}

func BuildTrainingController(spec *demov1alpha1.OpniDemo) *appsv1.Deployment {
	labels := map[string]string{"app": "training-controller"}
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "training-controller",
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:            "training-controller",
							Image:           "rancher/opni-training-controller:v0.1.1",
							ImagePullPolicy: v1.PullAlways,
							Env: []v1.EnvVar{
								{
									Name:  "NATS_SERVER_URL",
									Value: fmt.Sprintf("nats://nats_client:%s@nats-client.%s.svc:4222", spec.Spec.NatsPassword, spec.Namespace),
								},
								{
									Name:  "MINIO_SERVER_URL",
									Value: fmt.Sprintf("http://minio.%s.svc.cluster.local:9000", spec.Namespace),
								},
								{
									Name:  "MINIO_ACCESS_KEY",
									Value: spec.Spec.MinioAccessKey,
								},
								{
									Name:  "MINIO_SECRET_KEY",
									Value: spec.Spec.MinioSecretKey,
								},
								{
									Name:  "JOB_NAMESPACE",
									Value: "default",
								},
								{
									Name:  "ES_ENDPOINT",
									Value: fmt.Sprintf("https://opendistro-es-client-service.%s.svc.cluster.local:9200", spec.Namespace),
								},
								{
									Name:  "ES_USERNAME",
									Value: spec.Spec.ElasticsearchUser,
								},
								{
									Name:  "ES_PASSWORD",
									Value: spec.Spec.ElasticsearchPassword,
								},
								{
									Name:  "NODE_TLS_REJECT_UNAUTHORIZED",
									Value: "0",
								},
							},
						},
					},
				},
			},
		},
	}
}