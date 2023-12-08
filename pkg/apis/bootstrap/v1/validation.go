package v1

import "github.com/open-panoptes/opni/pkg/validation"

func (h *BootstrapAuthRequest) Validate() error {
	if h.ClientID == "" {
		return validation.Errorf("%w: %s", validation.ErrMissingRequiredField, "ClientID")
	}
	if err := validation.ValidateID(h.ClientID); err != nil {
		return validation.ErrInvalidID
	}
	if len(h.ClientPubKey) == 0 {
		return validation.Errorf("%w: %s", validation.ErrMissingRequiredField, "ClientPubKey")
	}
	if h.Capability == "" {
		return validation.Errorf("%w: %s", validation.ErrMissingRequiredField, "Capability")
	}
	return nil
}
