// Copyright (c) ZStack.io, Inc.

package param

// UpdateFiSecSecretResourcePoolDetailParam UpdateFiSecSecretResourcePool detail param
type UpdateFiSecSecretResourcePoolDetailParam struct {
	KeyNum string `json:"keyNum,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Model string `json:"model,omitempty"`
	HeartbeatInterval int `json:"heartbeatInterval,omitempty"`
}

// UpdateFiSecSecretResourcePoolParam UpdateFiSecSecretResourcePool request param
type UpdateFiSecSecretResourcePoolParam struct {
	BaseParam
	Params UpdateFiSecSecretResourcePoolDetailParam `json:"params"`
}
