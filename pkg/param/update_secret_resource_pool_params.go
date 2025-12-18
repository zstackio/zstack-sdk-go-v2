// Copyright (c) ZStack.io, Inc.

package param

// UpdateSecretResourcePoolDetailParam UpdateSecretResourcePool detail param
type UpdateSecretResourcePoolDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Model string `json:"model,omitempty"`
	HeartbeatInterval int `json:"heartbeatInterval,omitempty"`
}

// UpdateSecretResourcePoolParam UpdateSecretResourcePool request param
type UpdateSecretResourcePoolParam struct {
	BaseParam
	Params UpdateSecretResourcePoolDetailParam `json:"params"`
}
