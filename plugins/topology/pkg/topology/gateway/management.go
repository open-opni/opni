package gateway

import (
	"github.com/open-panoptes/opni/pkg/logger"
	"github.com/open-panoptes/opni/plugins/topology/pkg/topology/gateway/drivers"
)

func (p *Plugin) configureTopologyManagement() {
	drivers.ResetClusterDrivers()

	if kcd, err := drivers.NewTopologyManagerClusterDriver(); err == nil {
		drivers.RegisterClusterDriver(kcd)
	} else {
		drivers.LogClusterDriverFailure(kcd.Name(), err) // Name() is safe to call on a nil pointer
	}
	name := "topology-manager"
	driver, err := drivers.GetClusterDriver(name)
	if err != nil {
		p.logger.With(
			"driver", name,
			logger.Err(err),
		).Error("failed to load cluster driver, using fallback no-op driver")
		driver = &drivers.NoopClusterDriver{}
	}
	p.clusterDriver.Set(driver)
}
