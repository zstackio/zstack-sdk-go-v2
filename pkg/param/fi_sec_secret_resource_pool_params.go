// Copyright (c) ZStack.io, Inc.

package param

// UpdateFiSecSecretResourcePoolDetailParam UpdateFiSecSecretResourcePool详细参数
type UpdateFiSecSecretResourcePoolDetailParam struct {
	rest string `json:"keyNum,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"model,omitempty"`
	rest int `json:"heartbeatInterval,omitempty"`
}

// UpdateFiSecSecretResourcePoolParam UpdateFiSecSecretResourcePool请求参数
type UpdateFiSecSecretResourcePoolParam struct {
	BaseParam
	Params UpdateFiSecSecretResourcePoolDetailParam `json:"params"` // 详细参数
}

