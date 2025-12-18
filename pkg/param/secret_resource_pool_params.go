// Copyright (c) ZStack.io, Inc.

package param

// UpdateSecretResourcePoolDetailParam UpdateSecretResourcePool详细参数
type UpdateSecretResourcePoolDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"model,omitempty"`
	rest int `json:"heartbeatInterval,omitempty"`
}

// UpdateSecretResourcePoolParam UpdateSecretResourcePool请求参数
type UpdateSecretResourcePoolParam struct {
	BaseParam
	Params UpdateSecretResourcePoolDetailParam `json:"params"` // 详细参数
}

