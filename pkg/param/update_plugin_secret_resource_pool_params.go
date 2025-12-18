// Copyright (c) ZStack.io, Inc.

package param

// UpdatePluginSecretResourcePoolDetailParam UpdatePluginSecretResourcePool detail param
type UpdatePluginSecretResourcePoolDetailParam struct {
	Properties map[string]string `json:"properties,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Model string `json:"model,omitempty"`
	HeartbeatInterval int `json:"heartbeatInterval,omitempty"`
}

// UpdatePluginSecretResourcePoolParam UpdatePluginSecretResourcePool request param
type UpdatePluginSecretResourcePoolParam struct {
	BaseParam
	Params UpdatePluginSecretResourcePoolDetailParam `json:"params"`
}
