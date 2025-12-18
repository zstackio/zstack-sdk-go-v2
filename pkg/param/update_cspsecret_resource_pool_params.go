// Copyright (c) ZStack.io, Inc.

package param

// UpdateCSPSecretResourcePoolDetailParam UpdateCSPSecretResourcePool detail param
type UpdateCSPSecretResourcePoolDetailParam struct {
	AppId string `json:"appId,omitempty"`
	AppKey string `json:"appKey,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	Port int `json:"port,omitempty"`
	KeyId string `json:"keyId,omitempty"`
	UserId string `json:"userId,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Model string `json:"model,omitempty"`
	HeartbeatInterval int `json:"heartbeatInterval,omitempty"`
}

// UpdateCSPSecretResourcePoolParam UpdateCSPSecretResourcePool request param
type UpdateCSPSecretResourcePoolParam struct {
	BaseParam
	Params UpdateCSPSecretResourcePoolDetailParam `json:"params"`
}
