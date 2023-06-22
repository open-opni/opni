package notifications

import (
	"context"
	"sync"

	"github.com/rancher/opni/pkg/alerting/storage"
	alertingv1 "github.com/rancher/opni/pkg/apis/alerting/v1"
	"github.com/rancher/opni/pkg/util"
	"github.com/rancher/opni/pkg/util/future"
	"github.com/rancher/opni/plugins/alerting/pkg/alerting/server"
	"go.uber.org/zap"
)

type NotificationServerComponent struct {
	alertingv1.UnsafeAlertNotificationsServer

	util.Initializer

	mu sync.Mutex
	server.Config

	logger *zap.SugaredLogger

	conditionStorage future.Future[storage.ConditionStorage]
}

var _ server.ServerComponent = (*NotificationServerComponent)(nil)

func NewNotificationServerComponent(
	logger *zap.SugaredLogger,
) *NotificationServerComponent {
	return &NotificationServerComponent{
		logger:           logger,
		conditionStorage: future.New[storage.ConditionStorage](),
	}
}

type NotificationServerConfiguration struct {
	storage.ConditionStorage
}

func (n *NotificationServerComponent) Name() string {
	return "notification"
}

func (n *NotificationServerComponent) Status() server.Status {
	return server.Status{
		Running: n.Initialized(),
	}
}

func (n *NotificationServerComponent) Ready() bool {
	return n.Initialized()
}

func (n *NotificationServerComponent) Healthy() bool {
	return n.Initialized()
}

func (n *NotificationServerComponent) SetConfig(conf server.Config) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.Config = conf
}

func (n *NotificationServerComponent) Sync(_ context.Context, _ bool) error {
	return nil
}

func (n *NotificationServerComponent) Initialize(conf NotificationServerConfiguration) {
	n.InitOnce(func() {
		n.conditionStorage.Set(conf.ConditionStorage)
	})
}
