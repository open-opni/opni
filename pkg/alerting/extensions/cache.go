package extensions

import (
	"github.com/open-panoptes/opni/pkg/alerting/cache"
	"github.com/open-panoptes/opni/pkg/alerting/drivers/config"
	alertingv1 "github.com/open-panoptes/opni/pkg/apis/alerting/v1"
)

func (e *EmbeddedServer) cacheAlarm(msgMeta cache.MessageMetadata, alert config.Alert) error {
	e.alarmCache.Set(alertingv1.OpniSeverity(msgMeta.Severity), e.alarmCache.Key(msgMeta), alert)
	return nil
}

func (e *EmbeddedServer) cacheNotification(msgMeta cache.MessageMetadata, alert config.Alert) error {
	e.notificationCache.Set(alertingv1.OpniSeverity(msgMeta.Severity), e.notificationCache.Key(msgMeta), alert)
	return nil
}
