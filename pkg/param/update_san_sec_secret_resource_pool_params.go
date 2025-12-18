// Copyright (c) ZStack.io, Inc.

package param

// UpdateSanSecSecretResourcePoolDetailParam UpdateSanSecSecretResourcePool detail param
type UpdateSanSecSecretResourcePoolDetailParam struct {
	KeyIndex int `json:"keyIndex,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	Port int `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Sm3Key string `json:"sm3Key,omitempty"`
	Sm4Key string `json:"sm4Key,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Model string `json:"model,omitempty"`
	HeartbeatInterval int `json:"heartbeatInterval,omitempty"`
}

// UpdateSanSecSecretResourcePoolParam UpdateSanSecSecretResourcePool request param
type UpdateSanSecSecretResourcePoolParam struct {
	BaseParam
	Params UpdateSanSecSecretResourcePoolDetailParam `json:"params"`
}
