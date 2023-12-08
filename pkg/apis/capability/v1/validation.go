package v1

import "github.com/open-panoptes/opni/pkg/validation"

func (req *UninstallRequest) Validate() error {
	if err := validation.Validate(req.Cluster); err != nil {
		return err
	}
	return nil
}

func (req *InstallRequest) Validate() error {
	if err := validation.Validate(req.Cluster); err != nil {
		return err
	}
	return nil
}

func (req *SyncRequest) Validate() error {
	if req.Cluster != nil && req.Cluster.Id != "" {
		// empty string indicates "all" here
		if err := validation.Validate(req.Cluster); err != nil {
			return err
		}
	}
	return nil
}
