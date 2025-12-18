// Copyright (c) ZStack.io, Inc.

package param

// UpdateKoAlSecretResourcePoolDetailParam UpdateKoAlSecretResourcePool detail param
type UpdateKoAlSecretResourcePoolDetailParam struct {
	ManagementIp string `json:"managementIp,omitempty"`
	Port int `json:"port,omitempty"`
	SecretKey string `json:"secretKey,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Model string `json:"model,omitempty"`
	HeartbeatInterval int `json:"heartbeatInterval,omitempty"`
}

// UpdateKoAlSecretResourcePoolParam UpdateKoAlSecretResourcePool request param
type UpdateKoAlSecretResourcePoolParam struct {
	BaseParam
	Params UpdateKoAlSecretResourcePoolDetailParam `json:"params"`
}
