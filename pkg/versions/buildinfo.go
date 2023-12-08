package versions

import (
	"runtime/debug"

	corev1 "github.com/open-panoptes/opni/pkg/apis/core/v1"
	"github.com/samber/lo"
)

const (
	OpniVersionBuildInfoKey = "opni.version"
)

func ReadBuildInfo() (*corev1.BuildInfo, bool) {
	debugBuildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return nil, false
	}
	protoBuildInfo := &corev1.BuildInfo{
		GoVersion: debugBuildInfo.GoVersion,
		Path:      debugBuildInfo.Path,
		Main:      toModule(&debugBuildInfo.Main),
		// Ignore deps for now, they are not important here
		Settings: lo.Map(debugBuildInfo.Settings, func(s debug.BuildSetting, _ int) *corev1.BuildSetting {
			return toBuildSetting(s)
		}),
	}

	protoBuildInfo.Settings = append(protoBuildInfo.Settings,
		&corev1.BuildSetting{
			Key:   OpniVersionBuildInfoKey,
			Value: Version,
		},
	)
	return protoBuildInfo, true
}

func toModule(m *debug.Module) *corev1.Module {
	if m == nil {
		return nil
	}
	return &corev1.Module{
		Path:    m.Path,
		Version: m.Version,
		Replace: toModule(m.Replace),
	}
}

func toBuildSetting(s debug.BuildSetting) *corev1.BuildSetting {
	return &corev1.BuildSetting{
		Key:   s.Key,
		Value: s.Value,
	}
}
