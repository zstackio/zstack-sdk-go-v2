// Copyright (c) ZStack.io, Inc.

package param

// CreateJitSecretResourcePoolDetailParam CreateJitSecretResourcePool详细参数
type CreateJitSecretResourcePoolDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"model,omitempty"`
	rest string `json:"ability,omitempty"`
	rest string `json:"type" validate:"required"` // 必填
	rest int `json:"heartbeatInterval" validate:"required"` // 必填
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateJitSecretResourcePoolParam CreateJitSecretResourcePool请求参数
type CreateJitSecretResourcePoolParam struct {
	BaseParam
	Params CreateJitSecretResourcePoolDetailParam `json:"params"` // 详细参数
}

