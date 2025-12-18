// Copyright (c) ZStack.io, Inc.

package param

// UpdateInfoSecSecretResourcePoolDetailParam UpdateInfoSecSecretResourcePool detail param
type UpdateInfoSecSecretResourcePoolDetailParam struct {
	ConnectionMode int `json:"connectionMode,omitempty"`
	ActivatedToken string `json:"activatedToken,omitempty"`
	ProtectToken string `json:"protectToken,omitempty"`
	HmacToken string `json:"hmacToken,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Model string `json:"model,omitempty"`
	HeartbeatInterval int `json:"heartbeatInterval,omitempty"`
}

// UpdateInfoSecSecretResourcePoolParam UpdateInfoSecSecretResourcePool request param
type UpdateInfoSecSecretResourcePoolParam struct {
	BaseParam
	Params UpdateInfoSecSecretResourcePoolDetailParam `json:"params"`
}
