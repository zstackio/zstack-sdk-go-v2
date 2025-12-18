// Copyright (c) ZStack.io, Inc.

package param

// UpdateFlkSecSecretResourcePoolDetailParam UpdateFlkSecSecretResourcePool详细参数
type UpdateFlkSecSecretResourcePoolDetailParam struct {
	rest string `json:"encryptResult,omitempty"`
	rest string `json:"activatedToken,omitempty"`
	rest string `json:"protectToken,omitempty"`
	rest string `json:"hmacToken,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"model,omitempty"`
	rest int `json:"heartbeatInterval,omitempty"`
}

// UpdateFlkSecSecretResourcePoolParam UpdateFlkSecSecretResourcePool请求参数
type UpdateFlkSecSecretResourcePoolParam struct {
	BaseParam
	Params UpdateFlkSecSecretResourcePoolDetailParam `json:"params"` // 详细参数
}

