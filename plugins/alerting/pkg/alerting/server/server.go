package server

import (
	"context"

	"github.com/rancher/opni/pkg/alerting/client"
)

type InitializerF interface {
	InitOnce(f func())
	Initialized() bool
	WaitForInit()
	WaitForInitContext(ctx context.Context) error
}

type ServerComponent interface {
	InitializerF
	ComponentStatus
	Name() string

	SetConfig(config Config)
	// Server components that manage independent dependencies
	// should implement this method to sync them
	// Sync implementations should be cancellable on ctx.Err()
	Sync(ctx context.Context, shouldSync bool) error
}

type Config struct {
	Client client.AlertingClient
}

type ComponentStatus interface {
	Ready() bool
	Healthy() bool
	Status() Status
}

type Status struct {
	Running bool
}