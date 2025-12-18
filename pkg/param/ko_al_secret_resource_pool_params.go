// Copyright (c) ZStack.io, Inc.

package param

// UpdateKoAlSecretResourcePoolDetailParam UpdateKoAlSecretResourcePool详细参数
type UpdateKoAlSecretResourcePoolDetailParam struct {
	rest string `json:"managementIp,omitempty"`
	rest int `json:"port,omitempty"`
	rest string `json:"secretKey,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"model,omitempty"`
	rest int `json:"heartbeatInterval,omitempty"`
}

// UpdateKoAlSecretResourcePoolParam UpdateKoAlSecretResourcePool请求参数
type UpdateKoAlSecretResourcePoolParam struct {
	BaseParam
	Params UpdateKoAlSecretResourcePoolDetailParam `json:"params"` // 详细参数
}

