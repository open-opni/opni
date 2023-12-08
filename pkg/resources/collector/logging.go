package collector

import (
	"bytes"
	"path/filepath"

	opniloggingv1beta1 "github.com/open-panoptes/opni/apis/logging/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

func (r *Reconciler) generateDistributionReceiver(config *opniloggingv1beta1.CollectorConfig) (receiver []string, retBytes []byte, retErr error) {
	var providerReceiver bytes.Buffer

	switch config.Spec.Provider {
	case opniloggingv1beta1.LogProviderRKE:
		return []string{logReceiverRKE}, []byte(templateLogAgentRKE), nil
	case opniloggingv1beta1.LogProviderK3S:
		journaldDir := "/var/log/journal"
		if config.Spec.K3S != nil && config.Spec.K3S.LogPath != "" {
			journaldDir = config.Spec.K3S.LogPath
		}
		retErr = templateLogAgentK3s.Execute(&providerReceiver, journaldDir)
		if retErr != nil {
			return
		}
		return []string{logReceiverK3s}, providerReceiver.Bytes(), nil
	case opniloggingv1beta1.LogProviderRKE2:
		journaldDir := "/var/log/journal"
		if config.Spec.RKE2 != nil && config.Spec.RKE2.LogPath != "" {
			journaldDir = config.Spec.RKE2.LogPath
		}
		retErr = templateLogAgentRKE2.Execute(&providerReceiver, journaldDir)
		if retErr != nil {
			return
		}
		return []string{
			logReceiverRKE2,
			fileLogReceiverRKE2,
		}, providerReceiver.Bytes(), nil
	default:
		return
	}
}

func (r *Reconciler) generateKubeAuditLogsReceiver(config *opniloggingv1beta1.CollectorConfig) (string, []byte, error) {
	var receiver bytes.Buffer

	if config.Spec.KubeAuditLogs != nil && config.Spec.KubeAuditLogs.Enabled {
		filelogDir := "/var/log/kube-audit"
		if config.Spec.KubeAuditLogs.PathPrefix != "" {
			filelogDir = config.Spec.KubeAuditLogs.PathPrefix
		}

		auditLogFilename := "audit.log"
		if config.Spec.KubeAuditLogs.AuditFilename != "" {
			auditLogFilename = config.Spec.KubeAuditLogs.AuditFilename
		}

		err := templateKubeAuditLogs.Execute(&receiver, filepath.Join(filelogDir, auditLogFilename))
		if err != nil {
			return "", nil, err
		}

		return logReceiverKubeAuditLogs, receiver.Bytes(), nil
	}

	return "", nil, nil
}

func (r *Reconciler) hostLoggingVolumes() (
	retVolumeMounts []corev1.VolumeMount,
	retVolumes []corev1.Volume,
	retErr error,
) {
	config := &opniloggingv1beta1.CollectorConfig{}
	retErr = r.client.Get(r.ctx, types.NamespacedName{
		Name:      r.collector.Spec.LoggingConfig.Name,
		Namespace: r.collector.Spec.SystemNamespace,
	}, config)
	if retErr != nil {
		return
	}

	retVolumeMounts = append(retVolumeMounts, corev1.VolumeMount{
		Name:      "varlogpods",
		MountPath: "/var/log/pods",
		ReadOnly:  true,
	})
	retVolumes = append(retVolumes, corev1.Volume{
		Name: "varlogpods",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/log/pods",
			},
		},
	})

	retVolumeMounts = append(retVolumeMounts, corev1.VolumeMount{
		Name:      "varlibdockercontainers",
		MountPath: "/var/lib/docker/containers",
		ReadOnly:  true,
	})
	retVolumes = append(retVolumes, corev1.Volume{
		Name: "varlibdockercontainers",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/docker/containers",
			},
		},
	})

	kubeAuditLogsDir := "/var/log/kube-audit"
	if config.Spec.KubeAuditLogs != nil && config.Spec.KubeAuditLogs.PathPrefix != "" {
		kubeAuditLogsDir = config.Spec.KubeAuditLogs.PathPrefix
	}

	retVolumeMounts = append(retVolumeMounts, corev1.VolumeMount{
		Name:      "kubeauditlogs",
		MountPath: kubeAuditLogsDir,
		ReadOnly:  true,
	})
	retVolumes = append(retVolumes, corev1.Volume{
		Name: "kubeauditlogs",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: kubeAuditLogsDir,
			},
		},
	})

	switch config.Spec.Provider {
	case opniloggingv1beta1.LogProviderRKE:
		retVolumeMounts = append(retVolumeMounts, corev1.VolumeMount{
			Name:      "rancher",
			MountPath: "/var/lib/rancher/rke/log",
			ReadOnly:  true,
		})
		retVolumes = append(retVolumes, corev1.Volume{
			Name: "rancher",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/lib/rancher/rke/log",
				},
			},
		})
	case opniloggingv1beta1.LogProviderK3S:
		journaldDir := "/var/log/journal"
		if config.Spec.K3S != nil && config.Spec.K3S.LogPath != "" {
			journaldDir = config.Spec.K3S.LogPath
		}
		retVolumeMounts = append(retVolumeMounts, corev1.VolumeMount{
			Name:      "journald",
			MountPath: journaldDir,
			ReadOnly:  true,
		})
		retVolumeMounts = append(retVolumeMounts, corev1.VolumeMount{
			Name:      "machineid",
			MountPath: machineID,
			ReadOnly:  true,
		})
		retVolumes = append(retVolumes, corev1.Volume{
			Name: "journald",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: journaldDir,
				},
			},
		})
		retVolumes = append(retVolumes, corev1.Volume{
			Name: "machineid",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: machineID,
				},
			},
		})

	case opniloggingv1beta1.LogProviderRKE2:
		journaldDir := "/var/log/journal"
		if config.Spec.RKE2 != nil && config.Spec.RKE2.LogPath != "" {
			journaldDir = config.Spec.RKE2.LogPath
		}
		retVolumeMounts = append(retVolumeMounts, corev1.VolumeMount{
			Name:      "journald",
			MountPath: journaldDir,
			ReadOnly:  true,
		})
		retVolumeMounts = append(retVolumeMounts, corev1.VolumeMount{
			Name:      "indir",
			MountPath: rke2AgentLogDir,
		})
		retVolumeMounts = append(retVolumeMounts, corev1.VolumeMount{
			Name:      "machineid",
			MountPath: machineID,
			ReadOnly:  true,
		})
		retVolumes = append(retVolumes, corev1.Volume{
			Name: "journald",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: journaldDir,
				},
			},
		})
		retVolumes = append(retVolumes, corev1.Volume{
			Name: "machineid",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: machineID,
				},
			},
		})
		retVolumes = append(retVolumes, corev1.Volume{
			Name: "indir",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: rke2AgentLogDir,
					Type: &directoryOrCreate,
				},
			},
		})
	}
	return
}
