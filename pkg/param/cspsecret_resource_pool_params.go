// Copyright (c) ZStack.io, Inc.

package param

// UpdateCSPSecretResourcePoolDetailParam UpdateCSPSecretResourcePool详细参数
type UpdateCSPSecretResourcePoolDetailParam struct {
	rest string `json:"appId,omitempty"`
	rest string `json:"appKey,omitempty"`
	rest string `json:"managementIp,omitempty"`
	rest int `json:"port,omitempty"`
	rest string `json:"keyId,omitempty"`
	rest string `json:"userId,omitempty"`
	rest string `json:"protocol,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"model,omitempty"`
	rest int `json:"heartbeatInterval,omitempty"`
}

// UpdateCSPSecretResourcePoolParam UpdateCSPSecretResourcePool请求参数
type UpdateCSPSecretResourcePoolParam struct {
	BaseParam
	Params UpdateCSPSecretResourcePoolDetailParam `json:"params"` // 详细参数
}

