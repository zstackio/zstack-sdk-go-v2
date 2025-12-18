// Copyright (c) ZStack.io, Inc.

package param

// UpdateSanSecSecretResourcePoolDetailParam UpdateSanSecSecretResourcePool详细参数
type UpdateSanSecSecretResourcePoolDetailParam struct {
	rest int `json:"keyIndex,omitempty"`
	rest string `json:"managementIp,omitempty"`
	rest int `json:"port,omitempty"`
	rest string `json:"username,omitempty"`
	rest string `json:"password,omitempty"`
	rest string `json:"sm3Key,omitempty"`
	rest string `json:"sm4Key,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"model,omitempty"`
	rest int `json:"heartbeatInterval,omitempty"`
}

// UpdateSanSecSecretResourcePoolParam UpdateSanSecSecretResourcePool请求参数
type UpdateSanSecSecretResourcePoolParam struct {
	BaseParam
	Params UpdateSanSecSecretResourcePoolDetailParam `json:"params"` // 详细参数
}

